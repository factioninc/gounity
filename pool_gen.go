// DO NOT EDIT.
// GENERATED by go:generate at 2019-01-01 09:01:44.30889602 +0000 UTC.
package gounity

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Pool defines `pool` type.
type Pool struct {
	Resource

	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Health      *Health `json:"health"`
	SizeFree    uint64  `json:"sizeFree"`
	SizeTotal   uint64  `json:"sizeTotal"`
	SizeUsed    uint64  `json:"sizeUsed"`
}

var (
	typeNamePool   = "pool"
	typeFieldsPool = strings.Join([]string{
		"id",
		"name",
		"description",
		"health",
		"sizeFree",
		"sizeTotal",
		"sizeUsed",
	}, ",")
)

type PoolOperatorGen interface {
	NewPoolById(id string) *Pool

	NewPoolByName(name string) *Pool

	GetPoolById(id string) (*Pool, error)

	GetPoolByName(name string) (*Pool, error)

	GetPools() ([]*Pool, error)

	FillPools(respEntries []*instanceResp) ([]*Pool, error)

	FilterPools(filter *filter) ([]*Pool, error)
}

// NewPoolById constructs a `Pool` object with id.
func (u *Unity) NewPoolById(
	id string,
) *Pool {

	return &Pool{
		Resource: Resource{
			typeName: typeNamePool, typeFields: typeFieldsPool, Unity: u,
		},
		Id: id,
	}
}

// NewPoolByName constructs a `pool` object with name.
func (u *Unity) NewPoolByName(
	name string,
) *Pool {

	return &Pool{
		Resource: Resource{
			typeName: typeNamePool, typeFields: typeFieldsPool, Unity: u,
		},
		Name: name,
	}
}

// Refresh updates the info from Unity.
func (r *Pool) Refresh() error {

	if r.Id == "" && r.Name == "" {
		return fmt.Errorf(
			"cannot refresh on pool without Id nor Name, resource:%v", r,
		)
	}

	var (
		latest *Pool
		err    error
	)

	switch r.Id {

	case "":
		if latest, err = r.Unity.GetPoolByName(r.Name); err != nil {
			return err
		}
		*r = *latest
	default:
		if latest, err = r.Unity.GetPoolById(r.Id); err != nil {
			return err
		}
		*r = *latest
	}
	return nil
}

// GetPoolById retrives the `pool` by given its id.
func (u *Unity) GetPoolById(
	id string,
) (*Pool, error) {

	res := u.NewPoolById(id)
	err := u.GetInstanceById(res.typeName, id, res.typeFields, res)
	if err != nil {
		if IsUnityError(err) {
			return nil, err
		}
		return nil, errors.Wrap(err, "get pool by id failed")
	}
	return res, nil
}

// GetPoolByName retrives the `pool` by given its name.
func (u *Unity) GetPoolByName(
	name string,
) (*Pool, error) {

	res := u.NewPoolByName(name)
	if err := u.GetInstanceByName(res.typeName, name, res.typeFields, res); err != nil {
		return nil, errors.Wrap(err, "get pool by name failed")
	}
	return res, nil
}

// GetPools retrives all `pool` objects.
func (u *Unity) GetPools() ([]*Pool, error) {

	return u.FilterPools(nil)
}

// FilterPools filters the `pool` objects by given filters.
func (u *Unity) FilterPools(
	filter *filter,
) ([]*Pool, error) {

	respEntries, err := u.GetCollection(typeNamePool, typeFieldsPool, filter)
	if err != nil {
		return nil, errors.Wrap(err, "filter pool failed")
	}
	res, err := u.FillPools(respEntries)
	if err != nil {
		return nil, errors.Wrap(err, "fill pools failed")
	}
	return res, nil
}

// FillPools generates the `pool` objects from collection query response.
func (u *Unity) FillPools(
	respEntries []*instanceResp,
) ([]*Pool, error) {

	resSlice := []*Pool{}
	for _, entry := range respEntries {
		res := u.NewPoolById("") // empty id for fake `Pool` object
		if err := u.unmarshalResource(entry.Content, res); err != nil {
			return nil, errors.Wrap(err, "decode to Pool failed")
		}
		resSlice = append(resSlice, res)
	}
	return resSlice, nil
}

// Repr represents a `pool` object using its id.
func (r *Pool) Repr() *idRepresent {

	log := logrus.WithField("pool", r)
	if r.Id == "" {
		log.Info("refreshing pool from unity")
		err := r.Refresh()
		if err != nil {
			log.WithError(err).Error("refresh pool from unity failed")
			return nil
		}
	}
	return &idRepresent{Id: r.Id}
}
