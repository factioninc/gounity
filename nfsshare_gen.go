// DO NOT EDIT.
// GENERATED by go:generate at 2019-06-13 09:18:08.210763142 +0000 UTC.
package gounity

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// NfsShare defines `nfsShare` type.
type NfsShare struct {
	Resource

	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	ExportPaths []string    `json:"exportPaths"`
	Filesystem  *Filesystem `json:"filesystem"`
}

var (
	typeNameNfsShare   = "nfsShare"
	typeFieldsNfsShare = strings.Join([]string{
		"id",
		"name",
		"description",
		"exportPaths",
		"filesystem",
		"filesystem.storageResource",
	}, ",")
)

type NfsShareOperatorGen interface {
	NewNfsShareById(id string) *NfsShare

	NewNfsShareByName(name string) *NfsShare

	GetNfsShareById(id string) (*NfsShare, error)

	GetNfsShareByName(name string) (*NfsShare, error)

	GetNfsShares() ([]*NfsShare, error)

	FillNfsShares(respEntries []*instanceResp) ([]*NfsShare, error)

	FilterNfsShares(filter *filter) ([]*NfsShare, error)
}

// NewNfsShareById constructs a `NfsShare` object with id.
func (u *Unity) NewNfsShareById(
	id string,
) *NfsShare {

	return &NfsShare{
		Resource: Resource{
			typeName: typeNameNfsShare, typeFields: typeFieldsNfsShare, Unity: u,
		},
		Id: id,
	}
}

// NewNfsShareByName constructs a `nfsShare` object with name.
func (u *Unity) NewNfsShareByName(
	name string,
) *NfsShare {

	return &NfsShare{
		Resource: Resource{
			typeName: typeNameNfsShare, typeFields: typeFieldsNfsShare, Unity: u,
		},
		Name: name,
	}
}

// Refresh updates the info from Unity.
func (r *NfsShare) Refresh() error {

	if r.Id == "" && r.Name == "" {
		return fmt.Errorf(
			"cannot refresh on nfsShare without Id nor Name, resource:%v", r,
		)
	}

	var (
		latest *NfsShare
		err    error
	)

	switch r.Id {

	case "":
		if latest, err = r.Unity.GetNfsShareByName(r.Name); err != nil {
			return err
		}
		*r = *latest
	default:
		if latest, err = r.Unity.GetNfsShareById(r.Id); err != nil {
			return err
		}
		*r = *latest
	}
	return nil
}

// GetNfsShareById retrives the `nfsShare` by given its id.
func (u *Unity) GetNfsShareById(
	id string,
) (*NfsShare, error) {

	res := u.NewNfsShareById(id)
	err := u.GetInstanceById(res.typeName, id, res.typeFields, res)
	if err != nil {
		if IsUnityError(err) {
			return nil, err
		}
		return nil, errors.Wrap(err, "get nfsShare by id failed")
	}
	return res, nil
}

// GetNfsShareByName retrives the `nfsShare` by given its name.
func (u *Unity) GetNfsShareByName(
	name string,
) (*NfsShare, error) {

	res := u.NewNfsShareByName(name)
	if err := u.GetInstanceByName(res.typeName, name, res.typeFields, res); err != nil {
		return nil, errors.Wrap(err, "get nfsShare by name failed")
	}
	return res, nil
}

// GetNfsShares retrives all `nfsShare` objects.
func (u *Unity) GetNfsShares() ([]*NfsShare, error) {

	return u.FilterNfsShares(nil)
}

// FilterNfsShares filters the `nfsShare` objects by given filters.
func (u *Unity) FilterNfsShares(
	filter *filter,
) ([]*NfsShare, error) {

	respEntries, err := u.GetCollection(typeNameNfsShare, typeFieldsNfsShare, filter)
	if err != nil {
		return nil, errors.Wrap(err, "filter nfsShare failed")
	}
	res, err := u.FillNfsShares(respEntries)
	if err != nil {
		return nil, errors.Wrap(err, "fill nfsShares failed")
	}
	return res, nil
}

// FillNfsShares generates the `nfsShare` objects from collection query response.
func (u *Unity) FillNfsShares(
	respEntries []*instanceResp,
) ([]*NfsShare, error) {

	resSlice := []*NfsShare{}
	for _, entry := range respEntries {
		res := u.NewNfsShareById("") // empty id for fake `NfsShare` object
		if err := u.unmarshalResource(entry.Content, res); err != nil {
			return nil, errors.Wrap(err, "decode to NfsShare failed")
		}
		resSlice = append(resSlice, res)
	}
	return resSlice, nil
}

// Repr represents a `nfsShare` object using its id.
func (r *NfsShare) Repr() *idRepresent {

	log := logrus.WithField("nfsShare", r)
	if r.Id == "" {
		log.Info("refreshing nfsShare from unity")
		err := r.Refresh()
		if err != nil {
			log.WithError(err).Error("refresh nfsShare from unity failed")
			return nil
		}
	}
	return &idRepresent{Id: r.Id}
}
