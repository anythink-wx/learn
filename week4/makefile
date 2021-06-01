export GO111MODULE=on
export GOPROXY=https://goproxy.cn/
export GIN_MODE=release

BINNAME := "loginServer"
LDFLAGSPro := "-s -w -X 'main.BuildTime=$(shell date -u '+%F-%Z/%T')' -X 'main.GoVersion=$(shell go version)' -X 'main.BuildMode=release'"

GO ?= go

TIME := "log"$(shell date +"%Y%m%d")
APP_RELATIVE_PATH=$(shell echo $${PWD\#/*/*/*/*/*/})
##API_PROTO_FILES=$(shell cd $(APP_RELATIVE_PATH) && find -L . -name *.proto)

API_PROTO_FILES=$(shell cd $(APP_RELATIVE_PATH) && find -L . -name "*.proto")


proto3:
	protoc --go_out=$(APP_RELATIVE_PATH) $(API_PROTO_FILES)

clean:
	@rm -fr ./build/$(BINNAME)
	@rm -fr ./crash
.PHONY: release

## find -L .  -name "*.proto" -L：跟踪连接


##cd proto && protoc --go_out=../. *.proto