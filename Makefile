build:
	GOOS=linux GOARCH=amd64 go build -o ./bin  ./...
deploy: build
	ssh git-server "/root/deploy/lmrl/run.sh stop"
	scp ./bin/lmrl git-server:/root/deploy/lmrl/
	scp ./scripts/run.sh git-server:/root/deploy/lmrl/
	scp templates/*.* git-server:/root/deploy/lmrl/templates/
	ssh git-server "/root/deploy/lmrl/run.sh start"