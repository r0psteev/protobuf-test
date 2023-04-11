.PHONY: proto

all:
	go build  -o cmd/main cmd/main.go
proto:
	protoc --proto_path=proto proto/*.proto --go_out=proto --go_opt=paths=source_relative

clean:
	rm cmd/main ping.json ping.wire ping.wire.lte
