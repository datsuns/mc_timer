BIN := mct.exe mock.exe

build: $(BIN)

test: $(BIN)
	./mct.exe

%.exe: %.go
	go build -o $@ $^

.PHONY: build test
