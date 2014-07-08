include $(GOROOT)/src/Make.inc

TARG=github.com/miolini/goaero
GOFILES=
CGOFILES=client.go config.go
CGO_OFILES=aerospike.o

include $(GOROOT)/src/Make.pkg

format:
	gofmt -w *.go

docs:
	gomake clean
	godoc ${TARG} > README.txt
