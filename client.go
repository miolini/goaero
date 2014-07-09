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
	"fmt"
	"log"
	"runtime"
	"unsafe"
)

func as_error(e C.as_error) error {
	return fmt.Errorf("err(%d) %s at [%s:%d]", e.code, C.GoString(&e.message[0]), C.GoString(e.file), e.line)
}

type Config struct {
	as_config C.as_config
	lastHost  int
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
	aerospike C.aerospike
	config    *Config
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
		return as_error(e)
	}
	return
}

func (self *Aerospike) Close() (err error) {
	var e C.as_error
	if C.aerospike_close(&self.aerospike, &e) != C.AEROSPIKE_OK {
		return as_error(e)
	}
	C.aerospike_destroy(&self.aerospike)
	return
}

func (self *Aerospike) Put(key *Key, rec *Record, policy_write *PolicyWrite) (err error) {
	var e C.as_error
	var policy *C.as_policy_write
	if policy_write != nil {
		policy = &policy_write.as_policy_write
	}
	if C.aerospike_key_put(&self.aerospike, &e, policy, &key.as_key, rec.p_as_record) != C.AEROSPIKE_OK {
		return as_error(e)
	}
	return
}

func (self *Aerospike) Get(key *Key, rec *Record, policy_read *PolicyRead) (err error) {
	var e C.as_error
	var policy *C.as_policy_read
	if policy_read != nil {
		policy = &policy_read.as_policy_read
	}
	if C.aerospike_key_get(&self.aerospike, &e, policy, &key.as_key, &rec.p_as_record) != C.AEROSPIKE_OK {
		return as_error(e)
	}
	return
}

// TODO: policy methods
type PolicyWrite struct {
	as_policy_write C.as_policy_write
}

// TODO: policy methods
type PolicyRead struct {
	as_policy_read C.as_policy_read
}

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
	runtime.SetFinalizer(r, DestroyRecord)
	return
}

func DestroyRecord(rec *Record) {
	C.as_record_destroy(rec.p_as_record)
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
