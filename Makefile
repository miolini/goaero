all: submodules

submodules:
	git submodule update --init --recursive