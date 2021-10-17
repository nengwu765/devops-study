export IMAGE_NAME = nengwu765/httpserver
export TAG_NAME = v0.1

# 默认编译目标系统是 Linux 的二进制文件
all: clean build docker push

clean:
	echo "clean old binary"
	rm -rf bin
build:
	echo "building httpserver binary"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/httpserver main.go
docker: build
	echo "building httpserver container"
	docker build -t $(IMAGE_NAME):$(TAG_NAME) -f Dockerfile .
push:
	echo "pushing container to dockerhub"
	docker push $(IMAGE_NAME):$(TAG_NAME)