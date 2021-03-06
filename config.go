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
)

type Config struct {
	as_config C.as_config
	lastHost  int
}

func NewConfig() (self *Config) {
	self = new(Config)
	C.as_config_init(&self.as_config)
	//self.as_config.policies.retry = C.AS_POLICY_RETRY_ONCE
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
