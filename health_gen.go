// DO NOT EDIT.
// GENERATED by go:generate at 2019-05-17 11:31:28.027652 +0000 UTC.
package gounity

// Health defines `health` type.
type Health struct {
	Resource

	Value          int      `json:"value"`
	DescriptionIds []string `json:"descriptionIds"`
	Descriptions   []string `json:"descriptions"`
}
