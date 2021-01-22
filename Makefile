startup:
	sh scripts/startup.sh

test:
	go test -timeout 9000s -cover -a -v ./...
