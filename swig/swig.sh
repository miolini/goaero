#!/bin/bash

ASPATH=~/goaero-preinstall/aerospike-client-c
MODULE=aerospike
OUTDIR=gen

mkdir $OUTDIR

find \
	$ASPATH/target/*/include/aerospike/aerospike*.h \
	$ASPATH/target/*/include/aerospike/as_*.h \
	| while read -r line; do \
	swig2.0 -go -outdir gen -intgosize 32 -module aerospike $line \
	; done