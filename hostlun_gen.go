package gounity

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type genHostLunOperator interface {
	NewHostLunById(id string) *HostLun

	NewHostLunByName(name string) *HostLun

	GetHostLunById(id string) (*HostLun, error)

	GetHostLunByName(name string) (*HostLun, error)

	GetHostLuns() ([]*HostLun, error)

	FillHostLuns(respEntries []*instanceResp) ([]*HostLun, error)

	FilterHostLuns(filter *filter) ([]*HostLun, error)
}

// NewHostLunById constructs a `HostLun` object with id.
func (u *Unity) NewHostLunById(id string) *HostLun {
	return &HostLun{
		Resource: Resource{
			typeName: typeNameHostLun, typeFields: typeFieldsHostLun, unity: u,
		},
		Id: id,
	}
}

// NewHostLunByName constructs a `HostLun` object with name.
func (u *Unity) NewHostLunByName(name string) *HostLun {
	return &HostLun{
		Resource: Resource{
			typeName: typeNameHostLun, typeFields: typeFieldsHostLun, unity: u,
		},
		Name: name,
	}
}

// Refresh updates the info from Unity.
func (r *HostLun) Refresh() error {
	if r.Id == "" && r.Name == "" {
		return fmt.Errorf(
			"cannot refresh on resource without Id nor Name, resource:%v", r,
		)
	}

	var (
		latest *HostLun
		err    error
	)

	switch r.Id {
	case "":
		if latest, err = r.unity.GetHostLunByName(r.Name); err != nil {
			return err
		}
		r = latest
	default:
		if latest, err = r.unity.GetHostLunById(r.Id); err != nil {
			return err
		}
		r = latest
	}
	return nil
}

// GetHostLunById retrives the `HostLun` by given its id.
func (u *Unity) GetHostLunById(id string) (*HostLun, error) {
	res := u.NewHostLunById(id)

	err := u.GetInstanceById(res.typeName, id, res.typeFields, res)
	if err != nil {
		if IsUnityError(err) {
			return nil, err
		}
		return nil, errors.Wrap(err, "get HostLun by id failed")
	}
	return res, nil
}

// GetHostLunByName retrives the `HostLun` by given its name.
func (u *Unity) GetHostLunByName(name string) (*HostLun, error) {
	res := u.NewHostLunByName(name)
	if err := u.GetInstanceByName(res.typeName, name, res.typeFields, res); err != nil {
		return nil, errors.Wrap(err, "get HostLun by name failed")
	}
	return res, nil
}

// GetHostLuns retrives all `HostLun` objects.
func (u *Unity) GetHostLuns() ([]*HostLun, error) {

	return u.FilterHostLuns(nil)
}

// FilterHostLuns filters the `HostLun` objects by given filters.
func (u *Unity) FilterHostLuns(filter *filter) ([]*HostLun, error) {
	respEntries, err := u.GetCollection(typeNameHostLun, typeFieldsHostLun, filter)
	if err != nil {
		return nil, errors.Wrap(err, "filter HostLun failed")
	}
	res, err := u.FillHostLuns(respEntries)
	if err != nil {
		return nil, errors.Wrap(err, "fill HostLuns failed")
	}
	return res, nil
}

// FillHostLuns generates the `HostLun` objects from collection query response.
func (u *Unity) FillHostLuns(respEntries []*instanceResp) ([]*HostLun, error) {
	resSlice := []*HostLun{}
	for _, entry := range respEntries {
		res := u.NewHostLunById("") // empty id for fake `HostLun` object
		if err := json.Unmarshal(entry.Content, res); err != nil {
			return nil, errors.Wrapf(err, "decode to %v failed", res)
		}
		resSlice = append(resSlice, res)
	}
	return resSlice, nil
}

// Repr represents a `HostLun` object using its id.
func (r *HostLun) Repr() *idRepresent {
	if r.Id == "" {
		log.Infof("refreshing %v from unity", r)
		err := r.Refresh()
		if err != nil {
			log.Errorf("refresh %v from unity failed, %+v", r, err)
			return nil
		}
	}
	return &idRepresent{Id: r.Id}
}
