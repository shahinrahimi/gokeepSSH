build: 
	@go build -o ./bin/gokeepSSH
run: build
	@./bin/gokeepSSH

test:
	@go test -v