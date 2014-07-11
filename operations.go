package goaero

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: /usr/local/lib/libaerospike.a -llua -lcrypto -lz
// #include <stdlib.h>
// #include <aerospike/aerospike.h>
// #include <aerospike/aerospike_key.h>
// #include <aerospike/as_log.h>
import "C"

import (
	"unsafe"
	"runtime"
)

type Operations struct {
	as_ops C.as_operations
}

// destroy required
func NewOperations(num uint16) (o * Operations) {
	o = &Operations{}
	o.init(num)
	return
}

// self destroys
func NewOperationsEx() (o * Operations) {
	o = &Operations{}
	runtime.SetFinalizer(o, DestroyOperations)
	return
}

func (self * Operations) init(num uint16) {
	C.as_operations_init(&self.as_ops, C.uint16_t(num))
}

func (self * Operations) Reset(num uint16) {
	DestroyOperations(self)
	self.init(num)
}

func DestroyOperations(o * Operations) {
	C.as_operations_destroy(&o.as_ops)
}

func (self * Operations) AddIncr(bin_name string, value int64) {
	b := C.CString(bin_name)
	defer C.free(unsafe.Pointer(b))
	C.as_operations_add_incr(&self.as_ops, b, C.int64_t(value))
}

func (self * Operations) AddRead(bin_name string) {
	b := C.CString(bin_name)
	defer C.free(unsafe.Pointer(b))
	C.as_operations_add_read(&self.as_ops, b)
}

func (self * Operations) AddTouch() {
	C.as_operations_add_touch(&self.as_ops)
}
