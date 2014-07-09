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
	"log"
	"unsafe"
	"fmt"
)

type Config struct {
	as_config C.as_config
	lastHost int
}

func NewConfig() (self *Config) {
	self = new(Config)
	C.as_config_init(&self.as_config)
	return
}

func (self *Config) AddHost(addr string, port uint16) {
	if self.lastHost >= C.AS_CONFIG_HOSTS_SIZE {
		log.Printf("WARNING: reached max hosts size: %d", C.AS_CONFIG_HOSTS_SIZE)
		return
	}
	c_host := C.as_config_host{}
	cAddr := C.CString(addr)
	c_host.addr = C.CString(addr)
	c_host.port = C.uint16_t(port)
	self.as_config.hosts[self.lastHost] = c_host
	self.lastHost++
	C.free(unsafe.Pointer(cAddr))
}

func (self *Config) getCStruct() *C.as_config {
	return &self.as_config
}

type Aerospike struct {
	aerospike	C.aerospike
	config		*Config
}

func NewAerospike(config *Config) (self *Aerospike) {
	self = new(Aerospike)
	self.config = config
	C.aerospike_init(&self.aerospike, self.config.getCStruct())
	return
}

func (self *Aerospike) Connect() (err error) {
	var e C.as_error
	if C.aerospike_connect(&self.aerospike, &e) != C.AEROSPIKE_OK {
		return asErr(e)
	}
	return
}

func (self * Aerospike) Close() (err error) {
	var e C.as_error
	if C.aerospike_close(&self.aerospike, &e) != C.AEROSPIKE_OK {
		return asErr(e)
	}
	C.aerospike_destroy(&self.aerospike)
	return
}

func (self *Aerospike) Get() (err error) {
	var e C.as_error
	var as_record * C.as_record
	var as_key C.as_key
	result := C.aerospike_key_get(&self.aerospike, &e, nil, &as_key, &as_record)
	if result != C.AEROSPIKE_OK {
		err = asErr(e)
	} else {
	}
	return
}

func asErr(e C.as_error) error {
	return fmt.Errorf("err(%d) %s at [%s:%d]", e.code, C.GoString(&e.message[0]), C.GoString(e.file), e.line)
}

func (self *Aerospike) Set() {
}

type Key struct {
	as_key C.as_key
	n, s, k * C.char
}

func NewKey(namespace string, set string, key string) (k * Key) {
	k = &Key{}
	k.n = C.CString(namespace)
	k.s = C.CString(set)
	k.k = C.CString(key)
	C.as_key_init_str(&k.as_key, k.n, k.s, k.k)
	return
}

func (self * Key) Destroy() {
	C.free(unsafe.Pointer(self.n))
	C.free(unsafe.Pointer(self.s))
	C.free(unsafe.Pointer(self.k))
}

type Record struct {
	as_record C.as_record
}

func NewRecord(num_bins uint16) (r * Record) {
	r = &Record{}
	C.as_record_init(&r.as_record, C.uint16_t(num_bins))
	return
}

func (self * Record) Destroy() {
	C.as_record_destroy(&self.as_record)
}
