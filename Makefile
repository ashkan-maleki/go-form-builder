#===============================
# Go commands
#===============================
tidy:
	# export GOPROXY="https://mirrors.tencent.com/go,https://proxy.golang.org,direct"
	#export GOPROXY="https://proxy.golang.org,direct"
	# export GOPROXY="https://goproxy.cn,https://proxy.golang.org,direct"
	# export GOPROXY="https://proxy.golang.com.cn,https://proxy.golang.org,direct"
	go mod tidy
	go mod vendor

run: pull
	go run cmd/main.go


pull-tidy: pull tidy

#===============================
# Git commands
#===============================

commit:
	git add .
	git commit -m "$(CMSG)"
	git push

pull:
	git fetch
	git pull


#===============================
# Service: API gateway
#===============================

run-gateway:
	go run server/gateway/cmd/main.go