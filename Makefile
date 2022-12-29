buf: clean protoc-gen-go-equal
	~/go/bin/buf generate

protoc-gen-go-equal:
	go build

clean:
	rm -f protoc-gen-go-equal
