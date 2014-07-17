package goaero

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: /usr/local/lib/libaerospike.a -llua -lcrypto -lz
// #include <stdlib.h>
// #include <aerospike/aerospike.h>
// #include <aerospike/aerospike_key.h>
// #include <aerospike/aerospike_scan.h>
// #include <aerospike/as_log.h>
import "C"

import (
	"unsafe"
	"runtime"
)

type Scan struct {
	as_scan C.as_scan
}

type PCAsVal * C.as_val
type PCVoid * C.void

// destroy required
func NewScan(namespace string, set string) (s * Scan) {
	s = &Scan{}
	s.init(namespace, set)
	return
}

// self destroys
func NewScanEx() (s * Scan) {
	s = &Scan{}
	runtime.SetFinalizer(s, DestroyScan)
	return
}

func (self * Scan) init(namespace string, set string) {
	n := C.CString(namespace)
	defer C.free(unsafe.Pointer(n))
	s := C.CString(set)
	defer C.free(unsafe.Pointer(s))
	C.as_scan_init(&self.as_scan, n, s)
}

func (self * Scan) Reset(namespace string, set string) {
	DestroyScan(self)
	self.init(namespace, set)
}

func DestroyScan(s * Scan) {
	C.as_scan_destroy(&s.as_scan)
}
