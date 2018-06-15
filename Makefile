build:
	GOOS="linux" GOARCH="amd64" go build -v -ldflags '-d -s -w'

deploy:
	scp ChainChronicleGo dev:/home/mong/go/src/github.com/mong0520/ChainChronicleGo
