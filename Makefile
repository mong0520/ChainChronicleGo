service:
	docker-compose up -d

linebot:
	go run main.go -c conf/mong_iPhone.conf -m d

build:
	# GOOS="linux" GOARCH="amd64" go build -v -ldflags '-d -s -w'
	GOOS="linux" GOARCH="amd64" go build 

deploy:
	scp ChainChronicleGo dev:/home/mong/go/src/github.com/mong0520/ChainChronicleGo
