APPLICATION = God
VERSION = 0.0.1

.PHONY: build docker all

# 编译任务
build:
	GOOS=linux GOARCH=amd64 go build -o ./${APPLICATION}

# 镜像地址:registry/project/application
# 镜像构建/上传
docker:
	docker build -t ${REGISTRY}/${PROJECT}/${APPLICATION}:${VERSION} .
	docker push ${REGISTRY}/${PROJECT}/${APPLICATION}:${VERSION}

# 全部任务
all: build docker
