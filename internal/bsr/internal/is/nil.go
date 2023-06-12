// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package is

import "reflect"

// Nil checks if the interface is nil
func Nil(i any) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
