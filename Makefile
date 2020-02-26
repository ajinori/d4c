SOURCES := cmd/d4c/main.go
TARGET := bin/d4c

.PHONY: build clean distclean

build: $(TARGET)

$(TARGET): $(SOURCES)
	go build -o $@ $^

clean:
	-rm -rf *.swp *~

distclean:
	-rm -rf bin
