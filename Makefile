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
	genbibledata

progress:
	progress

checkbibletxt:
	@echo "【检查是否修改到已校验过的行，如果有内容，需要二次校对确认】"
	@git diff -- ./logic/bible/resources/bible.txt |grep '-'|grep -v '@@'|grep -v 'bible.txt' |grep -v '*' || true
	@echo "【查找半角的冒号，如果有，修改成全角的冒号】"
	@awk '{rest = $$0; sub(/^[^ ]* /, "", rest); if (rest ~ /:/) print $$0}' ./logic/bible/resources/bible.txt
	@echo "【查找半角的特殊字符，如果有，修改成全角的特殊字符】"
	@grep --color=always "[,!;.\"?'\(\)]" ./logic/bible/resources/bible.txt || true
	@echo "【查找末尾的开引号，如果有，需要二次校验确认】"
	@cat ./logic/bible/resources/bible.txt | \
	grep -v '徒7:48' |\
	grep -v '徒26:1' |\
	grep -v '代下18:25' |\
	grep -v '代下34:20' |\
	grep -v '王上2:1' |\
	grep -v '出16:11' |\
	grep -v '出19:3' |\
	grep -v '民6:23' |\
	grep --color=always "[“‘]$$"  || true
		