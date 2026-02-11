//go:build !windows
// +build !windows

package certstore

import "errors"

// StoreType exists for API compatibility on non-Windows platforms.
// OpenSpecificStore is currently only implemented on Windows in this repo.
type StoreType int

func openSpecificStore(storeType StoreType, name string) (Store, error) {
	return nil, errors.New("OpenSpecificStore is only supported on Windows")
}
