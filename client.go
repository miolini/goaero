package goaero

// #cgo CFLAGS: -Iaerospike-client-c/src/include
// #include <aerospike/aerospike.h>
// #include <aerospike/aerospike.h>
// #include <aerospike/aerospike_key.h>
// #include <aerospike/as_error.h>
// #include <aerospike/as_operations.h>
// #include <aerospike/as_record.h>
// #include <aerospike/as_status.h>

import (
	"C"
)

type Aerospike struct {
}