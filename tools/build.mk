all: clean build

clean:
	rm -rf $(TARGET)

build:
	go build -o $(TARGET) $(SOURCE)

build_debug:
	go build -race -o $(TARGET) $(SOURCE)

run:
	$(TARGET) $*