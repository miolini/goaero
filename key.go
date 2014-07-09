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
	as_key  C.as_key
	n, s, k *C.char
}

func NewKey(namespace string, set string, key string) (k *Key) {
	k = &Key{}
	k.n = C.CString(namespace)
	k.s = C.CString(set)
	k.k = C.CString(key)
	C.as_key_init(&k.as_key, k.n, k.s, k.k)
	runtime.SetFinalizer(k, DestroyKey)
	return
}

func DestroyKey(key *Key) {
	C.free(unsafe.Pointer(key.n))
	C.free(unsafe.Pointer(key.s))
	C.free(unsafe.Pointer(key.k))
}

