package gounity

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const (
	reqTraceHead = `
--------------------
GOUNITY HTTP REQUEST
--------------------
`
	respTraceHead = `
---------------------
GOUNITY HTTP RESPONSE
---------------------
`
)

func dumpBody(header *http.Header) bool {
	return header.Get(HeaderKeyContentType) != HeaderValueContentTypeBinaryOctetStream
}

func traceRequest(ctx context.Context, req *http.Request) {

	reqBuffer, err := httputil.DumpRequest(req, dumpBody(&req.Header))
	if err != nil {
		return
	}
	log.Debug(strings.Join([]string{reqTraceHead, string(reqBuffer)}, "\n"))
}

func traceResponse(ctx context.Context, resp *http.Response) {

	respBuffer, err := httputil.DumpResponse(resp, dumpBody(&resp.Header))
	if err != nil {
		return
	}
	log.Debug(strings.Join([]string{respTraceHead, string(respBuffer)}, "\n"))
}

const (
	typeStorageResource    = "storageResource"
	actionCreateLun        = "createLun"
	actionModifyLun        = "modifyLun"
	actionCreateFilesystem = "createFilesystem"
	actionModifyFilesystem = "modifyFilesystem"
)

const (
	pathAPITypes     = "api/types"
	pathAPIInstances = "api/instances"
)

func buildUrl(baseURL, fields string, filter *filter) string {
	queryParams := map[string]string{"compact": "true", "fields": fields}
	if filter != nil {
		queryParams["filter"] = filter.String()
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		log.WithError(err).WithField("baseURL", baseURL).Error("failed to parse url")
		return ""
	}
	query := u.Query()
	for k, v := range queryParams {
		query.Add(k, v)
	}
	u.RawQuery = query.Encode()
	return u.String()
}

func queryCollectionUrl(res, fields string, filter *filter) string {
	return buildUrl(strings.Join([]string{pathAPITypes, res, "instances"}, "/"), fields,
		filter)
}

func queryInstanceUrl(res, id, fields string) string {
	return buildUrl(strings.Join([]string{pathAPIInstances, res, id}, "/"), fields, nil)
}

func postTypeUrl(typeName, action string) string {
	return strings.Join([]string{pathAPITypes, typeName, "action", action}, "/")
}

func postInstanceUrl(typeName, resId, action string) string {
	return strings.Join([]string{pathAPIInstances, typeName, resId, "action", action}, "/")
}

// UnityErrorMessage defines the error message struct returned by Unity.
type unityErrorMessage struct {
	Message string `json:"en-US"`
}

// UnityError defines the error struct returned by Unity.
type unityError struct {
	ErrorCode      int                 `json:"errorCode"`
	HttpStatusCode int                 `json:"httpStatusCode"`
	Messages       []unityErrorMessage `json:"messages"`
	Message        string
}

type unityErrorResp struct {
	Error *unityError `json:"error,omitempty"`
}

func (e *unityError) Error() string {
	return fmt.Sprintf(
		"error from unity, status code: %v, error code: %v, message: %v",
		e.HttpStatusCode, e.ErrorCode, e.Message,
	)
}

var errUnableParseRespToError = errors.New("unable parse response body to unity error")

func parseUnityError(reader io.Reader) (error, error) {
	resp := &unityErrorResp{}
	if err := json.NewDecoder(reader).Decode(resp); err != nil {
		return nil, err
	}
	if resp.Error == nil {
		// not a `unity error` json in reader
		return nil, errUnableParseRespToError
	}

	respError := resp.Error
	respError.Message = respError.Messages[0].Message
	return respError, nil
}

func parseResourceName(url string) (string, error) {
	// url is like `/api/xxx` without host name/IP and port
	parts := strings.Split(url, "/")
	// i.e. `/api/instances/lun/sv_1` and `/api/types/lun/instances`
	if len(parts) < 4 {
		msg := "cannot find resource name in url"
		log.WithField("url", url).Error(msg)
		return "", errors.New(msg)
	}
	return parts[3], nil
}

type mockIndex struct {
	Indices []struct {
		URL      string      `json:"url"`
		Body     interface{} `json:"body"`
		Response string      `json:"response"`
	} `json:"indices"`
}

func mockServerHandler(resp http.ResponseWriter, req *http.Request) {
	reqURL := req.URL.String() // url is like `/api/xxx` without host IP and port
	unescapeURL, err := url.QueryUnescape(reqURL)
	if err != nil {
		log.WithError(err).Error("failed to get unescape url")
		resp.WriteHeader(http.StatusNotFound)
		return
	}
	resource, err := parseResourceName(unescapeURL)
	if err != nil {
		log.WithError(err).WithField("unescapeURL", unescapeURL).Error(
			"failed to get resource type from unescape url")
		resp.WriteHeader(http.StatusNotFound)
		return
	}
	cwd, err := os.Getwd()
	if err != nil {
		log.WithError(err).Error("failed to get current working directory")
		resp.WriteHeader(http.StatusNotFound)
		return
	}
	resourceDataDir := filepath.Join(cwd, "testdata", resource)
	indexFilePath := filepath.Join(resourceDataDir, "index.json")

	indicesBytes, err := ioutil.ReadFile(indexFilePath)
	if err != nil {
		log.WithError(err).WithField("filepath", indexFilePath).Error(
			"failed to read index.json file")
		resp.WriteHeader(http.StatusNotFound)
		return
	}
	dec := json.NewDecoder(bytes.NewReader(indicesBytes))
	dec.UseNumber()
	var indices mockIndex
	if err = dec.Decode(&indices); err != nil {
		log.WithError(err).WithField("filepath", indexFilePath).Error(
			"failed to parse index.json file")
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	dec = json.NewDecoder(req.Body)
	dec.UseNumber()
	var reqBody interface{}
	if err = dec.Decode(&reqBody); err != nil && err != io.EOF {
		log.WithError(err).Error("failed to decode request body")
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	respFileName := ""
	for _, index := range indices.Indices {
		if index.URL == unescapeURL {
			log.WithField("requestBody", reqBody).WithField(
				"mockBody", index.Body).Debug("check if request and mock body matches")
			if reflect.DeepEqual(reqBody, index.Body) {
				respFileName = index.Response
				break
			}
		}
	}
	if respFileName == "" {
		log.WithFields(map[string]interface{}{
			"filepath":    indexFilePath,
			"urlAfterIP":  unescapeURL,
			"requestBody": reqBody,
		}).Error("failed to find response for request in index.json")
		resp.WriteHeader(http.StatusNotFound)
		return
	}
	respFilePath := filepath.Join(resourceDataDir, respFileName)
	if respBytes, err := ioutil.ReadFile(respFilePath); err != nil {
		log.WithField("respFilePath", respFilePath).Error(
			"failed to read the response file")
		resp.WriteHeader(http.StatusNotFound)
	} else {
		mockError, err := parseUnityError(bytes.NewReader(respBytes))
		if mockError != nil {
			if ue, ok := mockError.(*unityError); ok {
				resp.WriteHeader(ue.HttpStatusCode)
			}
		} else if err != nil && err != io.EOF && err != errUnableParseRespToError {
			resp.WriteHeader(http.StatusNotFound)
		}
		resp.Write(respBytes)
	}
}

func setupMockServer() *httptest.Server {
	return httptest.NewTLSServer(http.HandlerFunc(mockServerHandler))
}

type testContext struct {
	mockServer *httptest.Server
	context    context.Context
	restClient RestClient
	unity      *Unity
}

func newTestContext() (*testContext, error) {
	mockServer := setupMockServer()
	ctx := context.Background()
	restClient, err := NewRestClient(ctx, mockServer.URL,
		"", "", NewRestClientOptions(true, true))
	if err != nil {
		return nil, err
	}
	unity, err := NewUnity(extractIp(mockServer.URL), "", "", true)
	if err != nil {
		return nil, err
	}
	return &testContext{mockServer: mockServer, context: ctx, restClient: restClient,
		unity: unity}, nil
}

func (c *testContext) tearDown() {
	c.mockServer.Close()
}

func extractIp(url string) string {
	return strings.Split(url, "/")[2]
}

func gbToBytes(gb uint64) uint64 {
	return gb * 1024 * 1024 * 1024
}

type idRepresent struct {
	Id string `json:"id"`
}
