// DO NOT EDIT.
// GENERATED by go:generate at 2019-05-17 11:31:28.024658 +0000 UTC.
package gounity

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// EthernetPort defines `ethernetPort` type.
type EthernetPort struct {
	Resource

	Health   *Health `json:"health"`
	Id       string  `json:"id"`
	IsLinkUp bool    `json:"isLinkUp"`
	Name     string  `json:"name"`
}

var (
	typeNameEthernetPort   = "ethernetPort"
	typeFieldsEthernetPort = strings.Join([]string{
		"health",
		"id",
		"isLinkUp",
		"name",
	}, ",")
)

type EthernetPortOperatorGen interface {
	NewEthernetPortById(id string) *EthernetPort

	NewEthernetPortByName(name string) *EthernetPort

	GetEthernetPortById(id string) (*EthernetPort, error)

	GetEthernetPortByName(name string) (*EthernetPort, error)

	GetEthernetPorts() ([]*EthernetPort, error)

	FillEthernetPorts(respEntries []*instanceResp) ([]*EthernetPort, error)

	FilterEthernetPorts(filter *filter) ([]*EthernetPort, error)
}

// NewEthernetPortById constructs a `EthernetPort` object with id.
func (u *Unity) NewEthernetPortById(
	id string,
) *EthernetPort {

	return &EthernetPort{
		Resource: Resource{
			typeName: typeNameEthernetPort, typeFields: typeFieldsEthernetPort, Unity: u,
		},
		Id: id,
	}
}

// NewEthernetPortByName constructs a `ethernetPort` object with name.
func (u *Unity) NewEthernetPortByName(
	name string,
) *EthernetPort {

	return &EthernetPort{
		Resource: Resource{
			typeName: typeNameEthernetPort, typeFields: typeFieldsEthernetPort, Unity: u,
		},
		Name: name,
	}
}

// Refresh updates the info from Unity.
func (r *EthernetPort) Refresh() error {

	if r.Id == "" && r.Name == "" {
		return fmt.Errorf(
			"cannot refresh on ethernetPort without Id nor Name, resource:%v", r,
		)
	}

	var (
		latest *EthernetPort
		err    error
	)

	switch r.Id {

	case "":
		if latest, err = r.Unity.GetEthernetPortByName(r.Name); err != nil {
			return err
		}
		*r = *latest
	default:
		if latest, err = r.Unity.GetEthernetPortById(r.Id); err != nil {
			return err
		}
		*r = *latest
	}
	return nil
}

// GetEthernetPortById retrives the `ethernetPort` by given its id.
func (u *Unity) GetEthernetPortById(
	id string,
) (*EthernetPort, error) {

	res := u.NewEthernetPortById(id)
	err := u.GetInstanceById(res.typeName, id, res.typeFields, res)
	if err != nil {
		if IsUnityError(err) {
			return nil, err
		}
		return nil, errors.Wrap(err, "get ethernetPort by id failed")
	}
	return res, nil
}

// GetEthernetPortByName retrives the `ethernetPort` by given its name.
func (u *Unity) GetEthernetPortByName(
	name string,
) (*EthernetPort, error) {

	res := u.NewEthernetPortByName(name)
	if err := u.GetInstanceByName(res.typeName, name, res.typeFields, res); err != nil {
		return nil, errors.Wrap(err, "get ethernetPort by name failed")
	}
	return res, nil
}

// GetEthernetPorts retrives all `ethernetPort` objects.
func (u *Unity) GetEthernetPorts() ([]*EthernetPort, error) {

	return u.FilterEthernetPorts(nil)
}

// FilterEthernetPorts filters the `ethernetPort` objects by given filters.
func (u *Unity) FilterEthernetPorts(
	filter *filter,
) ([]*EthernetPort, error) {

	respEntries, err := u.GetCollection(typeNameEthernetPort, typeFieldsEthernetPort, filter)
	if err != nil {
		return nil, errors.Wrap(err, "filter ethernetPort failed")
	}
	res, err := u.FillEthernetPorts(respEntries)
	if err != nil {
		return nil, errors.Wrap(err, "fill ethernetPorts failed")
	}
	return res, nil
}

// FillEthernetPorts generates the `ethernetPort` objects from collection query response.
func (u *Unity) FillEthernetPorts(
	respEntries []*instanceResp,
) ([]*EthernetPort, error) {

	resSlice := []*EthernetPort{}
	for _, entry := range respEntries {
		res := u.NewEthernetPortById("") // empty id for fake `EthernetPort` object
		if err := u.unmarshalResource(entry.Content, res); err != nil {
			return nil, errors.Wrap(err, "decode to EthernetPort failed")
		}
		resSlice = append(resSlice, res)
	}
	return resSlice, nil
}

// Repr represents a `ethernetPort` object using its id.
func (r *EthernetPort) Repr() *idRepresent {

	log := logrus.WithField("ethernetPort", r)
	if r.Id == "" {
		log.Info("refreshing ethernetPort from unity")
		err := r.Refresh()
		if err != nil {
			log.WithError(err).Error("refresh ethernetPort from unity failed")
			return nil
		}
	}
	return &idRepresent{Id: r.Id}
}
