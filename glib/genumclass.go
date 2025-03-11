package glib

/*
#include "glib.go.h"
*/
import "C"

import (
	"unsafe"
)

// EnumClass is a binding around the glib GEnumClass. It exposes methods
// to be used during the construction of objects backed by the go runtime.
type EnumClass struct {
	ptr *C.GEnumClass
}

// Unsafe is a convenience wrapper to return the unsafe.Pointer of the underlying C instance.
func (e *EnumClass) Unsafe() unsafe.Pointer { return unsafe.Pointer(e.ptr) }

// Instance returns the underlying C GEnumClass pointer
func (e *EnumClass) Instance() *C.GEnumClass { return e.ptr }

func (e *EnumClass) Minimum() int {
	return int(e.ptr.minimum)
}

func (e *EnumClass) Maximum() int {
	return int(e.ptr.maximum)
}

func (e *EnumClass) NValues() uint {
	return uint(e.ptr.n_values)
}

func (e *EnumClass) Values() []EnumValue {
	gSlice := (*[1 << 30]C.GEnumValue)(unsafe.Pointer(e.ptr.values))[:e.NValues():e.NValues()]
	out := []EnumValue{}

	for i := range e.NValues() {
		out = append(out, EnumValue{
			cvalue: gSlice[i],
		})
	}

	return out
}

type EnumValue struct {
	cvalue C.GEnumValue
}

func (v *EnumValue) Value() int {
	return int(v.cvalue.value)
}

func (v *EnumValue) Name() string {
	return C.GoString((*C.char)(v.cvalue.value_name))
}

func (v *EnumValue) Nick() string {
	return C.GoString((*C.char)(v.cvalue.value_nick))
}
