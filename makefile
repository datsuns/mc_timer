BIN := mct.exe mock.exe

build: $(BIN)

%.exe: %.go
	go build -o $@ $^

.PHONY: build
