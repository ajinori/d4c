SOURCES := cmd/d4c/main.go
TARGET := bin/d4c

.PHONY: build test clean distclean

build: $(TARGET)

$(TARGET): $(SOURCES)
	go build -o $@ $^

test:
	go test ./plugins/...

clean:
	-rm -rf *.swp *~

distclean:
	-rm -rf bin
