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

package main

import (
	"log"
	goaero "../../.."
	"flag"
)

var (
	ns = flag.String("ns", "test", "namespace")
	set = flag.String("s", "test-set", "test set name")
)

func main() {
	flag.Parse()
	var err error
	log.Printf("goaero example - get")
	config := goaero.NewConfig()
	config.AddHost("localhost", 3000)
	as := goaero.NewAerospike(config)
	err = as.Connect()
	checkErr(err)
	defer as.Close()
	key := goaero.NewKey(*ns, *set, "test-key")
	record := goaero.NewRecord(1)
	err = as.Put(key, record, nil)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("errro: %s", err)
	}
}