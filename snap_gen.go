// DO NOT EDIT.
// GENERATED by go:generate at 2019-05-17 11:31:28.045395 +0000 UTC.
package gounity

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Snap defines `snap` type.
type Snap struct {
	Resource

	Id              string              `json:"id"`
	Name            string              `json:"name"`
	Description     string              `json:"description"`
	StorageResource *StorageResource    `json:"storageResource"`
	Lun             *Lun                `json:"lun"`
	CreationTime    string              `json:"creationTime"`
	ExpirationTime  string              `json:"expirationTime"`
	CreatorType     SnapCreatorTypeEnum `json:"creatorType"`
	IsSystemSnap    bool                `json:"isSystemSnap"`
	IsModifiable    bool                `json:"isModifiable"`
	IsReadOnly      bool                `json:"isReadOnly"`
	IsModified      bool                `json:"isModified"`
	IsAutoDelete    bool                `json:"isAutoDelete"`
	State           SnapStateEnum       `json:"state"`
	Size            uint64              `json:"size"`
	HostAccess      []*SnapHostAccess   `json:"hostAccess"`
}

var (
	typeNameSnap   = "snap"
	typeFieldsSnap = strings.Join([]string{
		"id",
		"name",
		"description",
		"storageResource",
		"lun",
		"creationTime",
		"expirationTime",
		"creatorType",
		"isSystemSnap",
		"isModifiable",
		"isReadOnly",
		"isModified",
		"isAutoDelete",
		"state",
		"size",
		"hostAccess",
	}, ",")
)

type SnapOperatorGen interface {
	NewSnapById(id string) *Snap

	NewSnapByName(name string) *Snap

	GetSnapById(id string) (*Snap, error)

	GetSnapByName(name string) (*Snap, error)

	GetSnaps() ([]*Snap, error)

	FillSnaps(respEntries []*instanceResp) ([]*Snap, error)

	FilterSnaps(filter *filter) ([]*Snap, error)
}

// NewSnapById constructs a `Snap` object with id.
func (u *Unity) NewSnapById(
	id string,
) *Snap {

	return &Snap{
		Resource: Resource{
			typeName: typeNameSnap, typeFields: typeFieldsSnap, Unity: u,
		},
		Id: id,
	}
}

// NewSnapByName constructs a `snap` object with name.
func (u *Unity) NewSnapByName(
	name string,
) *Snap {

	return &Snap{
		Resource: Resource{
			typeName: typeNameSnap, typeFields: typeFieldsSnap, Unity: u,
		},
		Name: name,
	}
}

// Refresh updates the info from Unity.
func (r *Snap) Refresh() error {

	if r.Id == "" && r.Name == "" {
		return fmt.Errorf(
			"cannot refresh on snap without Id nor Name, resource:%v", r,
		)
	}

	var (
		latest *Snap
		err    error
	)

	switch r.Id {

	case "":
		if latest, err = r.Unity.GetSnapByName(r.Name); err != nil {
			return err
		}
		*r = *latest
	default:
		if latest, err = r.Unity.GetSnapById(r.Id); err != nil {
			return err
		}
		*r = *latest
	}
	return nil
}

// GetSnapById retrives the `snap` by given its id.
func (u *Unity) GetSnapById(
	id string,
) (*Snap, error) {

	res := u.NewSnapById(id)
	err := u.GetInstanceById(res.typeName, id, res.typeFields, res)
	if err != nil {
		if IsUnityError(err) {
			return nil, err
		}
		return nil, errors.Wrap(err, "get snap by id failed")
	}
	return res, nil
}

// GetSnapByName retrives the `snap` by given its name.
func (u *Unity) GetSnapByName(
	name string,
) (*Snap, error) {

	res := u.NewSnapByName(name)
	if err := u.GetInstanceByName(res.typeName, name, res.typeFields, res); err != nil {
		return nil, errors.Wrap(err, "get snap by name failed")
	}
	return res, nil
}

// GetSnaps retrives all `snap` objects.
func (u *Unity) GetSnaps() ([]*Snap, error) {

	return u.FilterSnaps(nil)
}

// FilterSnaps filters the `snap` objects by given filters.
func (u *Unity) FilterSnaps(
	filter *filter,
) ([]*Snap, error) {

	respEntries, err := u.GetCollection(typeNameSnap, typeFieldsSnap, filter)
	if err != nil {
		return nil, errors.Wrap(err, "filter snap failed")
	}
	res, err := u.FillSnaps(respEntries)
	if err != nil {
		return nil, errors.Wrap(err, "fill snaps failed")
	}
	return res, nil
}

// FillSnaps generates the `snap` objects from collection query response.
func (u *Unity) FillSnaps(
	respEntries []*instanceResp,
) ([]*Snap, error) {

	resSlice := []*Snap{}
	for _, entry := range respEntries {
		res := u.NewSnapById("") // empty id for fake `Snap` object
		if err := u.unmarshalResource(entry.Content, res); err != nil {
			return nil, errors.Wrap(err, "decode to Snap failed")
		}
		resSlice = append(resSlice, res)
	}
	return resSlice, nil
}

// Repr represents a `snap` object using its id.
func (r *Snap) Repr() *idRepresent {

	log := logrus.WithField("snap", r)
	if r.Id == "" {
		log.Info("refreshing snap from unity")
		err := r.Refresh()
		if err != nil {
			log.WithError(err).Error("refresh snap from unity failed")
			return nil
		}
	}
	return &idRepresent{Id: r.Id}
}

// Delete deletes a snap object.
func (r *Snap) Delete() error {

	log := logrus.WithField("snap", r)
	if r.StorageResource == nil {
		log.Info("refreshing snap from unity")
		err := r.Refresh()
		if err != nil {
			return errors.Wrap(err, "refresh snap from unity failed")
		}
	}

	err := r.Unity.DeleteInstance(typeStorageResource, r.StorageResource.Id)
	if err != nil {
		return errors.Wrap(err, "delete snap from unity failed")
	}
	return nil
}
