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

// #cgo CFLAGS: -Iaerospike-client-c/target/Linux-x86_64/include
// #cgo LDFLAGS: aerospike-client-c/target/Linux-x86_64/lib/libaerospike.a
// #include <aerospike/as_config.h>
// #include <aerospike/aerospike.h>
import "C"

import (
	"log"
)

type Config struct {
	c_config *C.as_config
	lastHost int
}

func (c *Config) AddHost(addr string) {
	if c.lastHost >= C.AS_CONFIG_HOSTS_SIZE {
		log.Printf("WARNING: reached max hosts size: %d", C.AS_CONFIG_HOSTS_SIZE)
		return
	}
	c_host := C.as_config_host{}
	c_host.addr = C.CString(addr)
	c_host.port = 3000
	c.c_config.hosts[c.lastHost] = c_host
	c.lastHost++
}