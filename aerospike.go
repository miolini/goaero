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
