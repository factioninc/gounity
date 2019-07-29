package gounity

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type createSnapshotResourceResp struct {
	Content struct {
		Id string `json:"id"`
	} `json:"content"`
}

type attachSnapshotResourceResp struct {
	Id string `json:"id"`
}

type AttachSnapshotRequest struct {
	Host          *Host
	AllowedAccess SnapAccessLevelEnum
}

func newCreateSnapshotBody(s *Snap, sr *StorageResource) map[string]interface{} {
	body := map[string]interface{}{
		"name":            s.Name,
		"storageResource": *sr.Repr(),
	}

	return body
}

func (s *Snap) Create(sr *StorageResource) error {
	body := newCreateSnapshotBody(s, sr)

	fields := map[string]interface{}{
		"requestBody": body,
	}
	log := logrus.WithFields(fields)
	msg := newMessage().withFields(fields)

	log.Debug("creating snapshot")
	resp := &createSnapshotResourceResp{}
	if err := s.Unity.CreateOnType(typeNameSnap, body, resp); err != nil {
		return errors.Wrapf(err, "create snapshot failed: %s", err)
	}

	log.WithField("createdSnapshotId", s.Id).Debug("snapshot created")
	s.Id = resp.Content.Id

	err := s.Refresh()
	if err != nil {
		return errors.Wrapf(
			err,
			"could not retrieve snapshot: %s", msg.withField("createdSnapshotId", s.Id),
		)
	}

	return err
}

func (s *Snap) Copy(copyName string) (*Snap, error) {
	body := map[string]interface{}{
		"copyName": copyName,
	}

	fields := map[string]interface{}{
		"requestBody": body,
	}

	log := logrus.WithFields(fields)
	msg := newMessage().withFields(fields)

	log.Debug("copying snapshot")
	err := s.Unity.PostOnInstance(
		typeNameSnap, s.Id, "copy", body, nil,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "copying snapshot failed: %s", msg)
	}

	snap := s.Unity.NewSnapByName(copyName)
	if err = snap.Refresh(); err != nil {
		return nil, errors.Wrapf(err, "get snapshot failed: %s", msg)
	}

	log.WithField("copySnapId", snap.Id).Debug("Snapshot successfully copied")
	return snap, err
}

func (s *Snap) AttachToHost(host *Host, access SnapAccessLevelEnum) (string, error) {
	hostRequests := []*AttachSnapshotRequest{
		{
			Host:          host,
			AllowedAccess: access,
		},
	}

	return s.AttachToHosts(hostRequests)
}

func (s *Snap) AttachToHosts(hostRequests []*AttachSnapshotRequest) (string, error) {
	if hostRequests == nil || len(hostRequests) == 0 {
		return "", nil
	}

	hostAccesses := []interface{}{}

	for _, hostRequest := range hostRequests {
		hostAccesses = append(hostAccesses, map[string]interface{}{
			"host":          *hostRequest.Host.Repr(),
			"allowedAccess": hostRequest.AllowedAccess,
		})
	}

	fields := map[string]interface{}{
		"requestBody": hostAccesses,
	}

	body := map[string]interface{}{"hostAccess": hostAccesses}

	log := logrus.WithFields(fields)
	msg := newMessage().withFields(fields)

	log.Debug("attaching snapshot")
	resp := &attachSnapshotResourceResp{}
	err := s.Unity.PostOnInstance(
		typeNameSnap, s.Id, "attach", body, resp,
	)
	if err != nil {
		return "", errors.Wrapf(err, "attaching snapshot failed: %s", msg)
	}

	log.WithField("copySnapId", s.Id).Debug("Snapshot successfully attached")
	return resp.Id, nil
}

func (s *Snap) DetachFromHost() (string, error) {
	body := map[string]interface{}{}

	logrus.Debug("detaching snapshot")
	resp := &attachSnapshotResourceResp{}
	err := s.Unity.PostOnInstance(
		typeNameSnap, s.Id, "detach", body, resp,
	)
	if err != nil {
		return "", errors.Wrapf(err, "detaching snapshot failed: %s", err)
	}

	logrus.WithField("snapId", s.Id).Debug("Snapshot successfully attached")
	return resp.Id, nil
}
