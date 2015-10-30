FROM gliderlabs/alpine
MAINTAINER Karthik Gaekwad

ENV GOPATH /go
RUN apk-install go
COPY . /go/src/github.com/stackhub/service-golang
WORKDIR /go/src/github.com/stackhub/service-golang
RUN go get && go build -o /bin/newyear

ENTRYPOINT ["/bin/newyear"]
