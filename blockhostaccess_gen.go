// DO NOT EDIT.
// GENERATED by go:generate at 2019-06-13 09:18:08.158640415 +0000 UTC.
package gounity

// BlockHostAccess defines `blockHostAccess` type.
type BlockHostAccess struct {
	Resource

	Host       *Host             `json:"host"`
	AccessMask HostLUNAccessEnum `json:"accessMask"`
}
