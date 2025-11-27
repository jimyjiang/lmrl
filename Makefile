pb: 
	protoc --go_out=. --go_opt=paths=source_relative logic/bible/schema.proto

fbuild:
	cd frontend && npm install && npm run build

build: fbuild
	# GOOS=linux,darwin GOARCH=amd64 go build -o ./bin  ./...
	./scripts/build.sh

dev: fbuild
	go build -o ./bin  ./...
	./bin/lmrl

deploy: build
	scp ./bin/linux_amd64/lmrl git-server:/root/deploy/lmrl/lmrl.new
	ssh git-server "mv -f /root/deploy/lmrl/lmrl.new /root/deploy/lmrl/lmrl && systemctl  restart lmrl.service"

install: fbuild
	go install ./...

start:
	ssh git-server "systemctl start lmrl"

stop:
	ssh git-server "systemctl stop lmrl"

status:
	ssh git-server "systemctl status lmrl"

restart:
	ssh git-server "systemctl  restart lmrl"
log:
	ssh git-server "journalctl -u lmrl.service"

genbibledata:
	go run ./cmd/genbibledata/