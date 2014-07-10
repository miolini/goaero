/*******************************************************************************
 * Copyright 2014 by Artem Andreenko, Openstat.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to
 * deal in the Software without restriction, including without limitation the
 * rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
 * sell copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
 * IN THE SOFTWARE.
 ******************************************************************************/

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
)

type Record struct {
	p_as_record *C.as_record
}

// TODO: add list methods
type List struct {
	as_list C.as_list
}

// TODO: add map methods
type Map struct {
	as_map C.as_map
}

func NewRecord(num_bins uint16) (r *Record) {
	r = &Record{}
	if num_bins > 0 {
		r.p_as_record = C.as_record_new(C.uint16_t(num_bins))
	}
	return
}

func DestroyRecord(rec *Record) {
	if rec.p_as_record != nil {
		C.as_record_destroy(rec.p_as_record)
	}
}

func (self *Record) SetInt64(name string, value int64) {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	C.as_record_set_int64(self.p_as_record, n, C.int64_t(value))
}

func (self *Record) SetString(name string, value string) {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	v := C.as_string_new(C.CString(value), true)
	// value	The value of the bin. Must last for the lifetime of the record.
	C.as_record_set_string(self.p_as_record, n, v)
}

func (self *Record) SetList(name string, value *List) {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	C.as_record_set_list(self.p_as_record, n, &value.as_list)
}

func (self *Record) SetMap(name string, value *Map) {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	C.as_record_set_map(self.p_as_record, n, &value.as_map)
}

func (self *Record) GetInt64(name string) int64 {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return int64(C.as_record_get_int64(self.p_as_record, n, 0x7FFFFFFFFFFFFFFF))
}

func (self *Record) GetString(name string) string {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return C.GoString(C.as_string_get(C.as_record_get_string(self.p_as_record, n)))
}
