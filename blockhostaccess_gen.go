// DO NOT EDIT.
// GENERATED by go:generate at 2019-06-06 09:02:53.35019 +0000 UTC.
package gounity

// BlockHostAccess defines `blockHostAccess` type.
type BlockHostAccess struct {
	Resource

	Host       *Host             `json:"host"`
	AccessMask HostLUNAccessEnum `json:"accessMask"`
}
