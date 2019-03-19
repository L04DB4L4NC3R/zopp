FROM golang

RUN mkdir -p /go/src/github.com/angadsharma1016/technica

ADD . /go/src/github.com/angadsharma1016/technica
WORKDIR /go/src/github.com/angadsharma1016/technica
#RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher
RUN go get github.com/ethereum/go-ethereum
EXPOSE 3000
ENTRYPOINT watcher -run github.com/angadsharma1016/technica/ -watch github.com/angadsharma1016/technica