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
	"runtime"
)

type Key struct {
	as_key C.as_key
}

// destroy required
func NewKey(namespace string, set string, key string) (k *Key) {
	k = &Key{}
	k.init(namespace, set, key)
	return
}

// destroy required
func NewKeyInt64(namespace string, set string, key int64) (k *Key) {
	k = &Key{}
	k.initInt64(namespace, set, key)
	return
}

// self destroys
func NewKeyEx() (k *Key) {
	k = &Key{}
	runtime.SetFinalizer(k, DestroyKey)
	return
}

func (self * Key) init(namespace string, set string, key string) {
	n := C.CString(namespace)
	defer C.free(unsafe.Pointer(n))
	s := C.CString(set)
	defer C.free(unsafe.Pointer(s))
	C.as_key_init_strp(&self.as_key, n, s, C.CString(key), true)
}

func (self * Key) initInt64(namespace string, set string, key int64) {
	n := C.CString(namespace)
	defer C.free(unsafe.Pointer(n))
	s := C.CString(set)
	defer C.free(unsafe.Pointer(s))
	C.as_key_init_int64(&self.as_key, n, s, C.int64_t(key))
}

func (self * Key) Reset(namespace string, set string, key string) {
	DestroyKey(self)
	self.init(namespace, set, key)
	return
}

func (self * Key) ResetInt64(namespace string, set string, key int64) {
	DestroyKey(self)
	self.initInt64(namespace, set, key)
	return
}

func DestroyKey(key *Key) {
	C.as_key_destroy(&key.as_key)
}
