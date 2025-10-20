fbuild:
	cd frontend && npm install && npm run build
build: fbuild
	GOOS=linux GOARCH=amd64 go build -o ./bin  ./...
deploy: build
	ssh git-server "/root/deploy/lmrl/run.sh stop"
	scp ./bin/lmrl git-server:/root/deploy/lmrl/
	scp ./scripts/run.sh git-server:/root/deploy/lmrl/
	ssh git-server "/root/deploy/lmrl/run.sh start"
install: fbuild
	go install ./...
restart:
	ssh git-server "/root/deploy/lmrl/run.sh restart"
pb: 
	protoc --go_out=. --go_opt=paths=source_relative logic/bible/schema.proto