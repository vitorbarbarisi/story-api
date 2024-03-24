FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/vitorbarbarisi/story-api
COPY . $GOPATH/src/github.com/vitorbarbarisi/story-api
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./story-api"]
