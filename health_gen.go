// DO NOT EDIT.
// GENERATED by go:generate at 2018-12-31 08:45:04.778825746 +0000 UTC.
package gounity

// Health defines `health` type.
type Health struct {
	Resource

	Value          int      `json:"value"`
	DescriptionIds []string `json:"descriptionIds"`
	Descriptions   []string `json:"descriptions"`
}
