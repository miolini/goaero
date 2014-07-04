all: submodules clean_native clean build_native build

submodules:
	git submodule update --init --recursive

build_native:
	CPATH=$(CPATH):/usr/include/lua5.1 $(MAKE) -C aerospike-client-c

clean_native:
	$(MAKE) -C aerospike-client-c clean

build:
	go build

clean:
