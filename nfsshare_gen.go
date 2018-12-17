package gounity

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type genNfsShareOperator interface {
	NewNfsShareById(id string) *NfsShare

	NewNfsShareByName(name string) *NfsShare

	GetNfsShareById(id string) (*NfsShare, error)

	GetNfsShareByName(name string) (*NfsShare, error)

	GetNfsShares() ([]*NfsShare, error)

	FillNfsShares(respEntries []*instanceResp) ([]*NfsShare, error)

	FilterNfsShares(filter *filter) ([]*NfsShare, error)
}

// NewNfsShareById constructs a `NfsShare` object with id.
func (u *Unity) NewNfsShareById(id string) *NfsShare {
	return &NfsShare{
		Resource: Resource{
			typeName: typeNameNfsShare, typeFields: typeFieldsNfsShare, unity: u,
		},
		Id: id,
	}
}

// NewNfsShareByName constructs a `NfsShare` object with name.
func (u *Unity) NewNfsShareByName(name string) *NfsShare {
	return &NfsShare{
		Resource: Resource{
			typeName: typeNameNfsShare, typeFields: typeFieldsNfsShare, unity: u,
		},
		Name: name,
	}
}

// Refresh updates the info from Unity.
func (r *NfsShare) Refresh() error {
	if r.Id == "" && r.Name == "" {
		return fmt.Errorf(
			"cannot refresh on resource without Id nor Name, resource:%v", r,
		)
	}

	var (
		latest *NfsShare
		err    error
	)

	switch r.Id {
	case "":
		if latest, err = r.unity.GetNfsShareByName(r.Name); err != nil {
			return err
		}
		r = latest
	default:
		if latest, err = r.unity.GetNfsShareById(r.Id); err != nil {
			return err
		}
		r = latest
	}
	return nil
}

// GetNfsShareById retrives the `NfsShare` by given its id.
func (u *Unity) GetNfsShareById(id string) (*NfsShare, error) {
	res := u.NewNfsShareById(id)

	err := u.GetInstanceById(res.typeName, id, res.typeFields, res)
	if err != nil {
		if IsUnityError(err) {
			return nil, err
		}
		return nil, errors.Wrap(err, "get NfsShare by id failed")
	}
	return res, nil
}

// GetNfsShareByName retrives the `NfsShare` by given its name.
func (u *Unity) GetNfsShareByName(name string) (*NfsShare, error) {
	res := u.NewNfsShareByName(name)
	if err := u.GetInstanceByName(res.typeName, name, res.typeFields, res); err != nil {
		return nil, errors.Wrap(err, "get NfsShare by name failed")
	}
	return res, nil
}

// GetNfsShares retrives all `NfsShare` objects.
func (u *Unity) GetNfsShares() ([]*NfsShare, error) {

	return u.FilterNfsShares(nil)
}

// FilterNfsShares filters the `NfsShare` objects by given filters.
func (u *Unity) FilterNfsShares(filter *filter) ([]*NfsShare, error) {
	respEntries, err := u.GetCollection(typeNameNfsShare, typeFieldsNfsShare, filter)
	if err != nil {
		return nil, errors.Wrap(err, "filter NfsShare failed")
	}
	res, err := u.FillNfsShares(respEntries)
	if err != nil {
		return nil, errors.Wrap(err, "fill NfsShares failed")
	}
	return res, nil
}

// FillNfsShares generates the `NfsShare` objects from collection query response.
func (u *Unity) FillNfsShares(respEntries []*instanceResp) ([]*NfsShare, error) {
	resSlice := []*NfsShare{}
	for _, entry := range respEntries {
		res := u.NewNfsShareById("") // empty id for fake `NfsShare` object
		if err := json.Unmarshal(entry.Content, res); err != nil {
			return nil, errors.Wrapf(err, "decode to %v failed", res)
		}
		resSlice = append(resSlice, res)
	}
	return resSlice, nil
}

// Repr represents a `NfsShare` object using its id.
func (r *NfsShare) Repr() *idRepresent {
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