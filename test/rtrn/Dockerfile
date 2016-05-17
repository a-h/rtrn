FROM golang:1.6
ENV GOPATH /go
ENV HTTP_PROXY http://10.0.2.2:3128
ENV HTTPS_PROXY http://10.0.2.2:3128
ENV http_proxy http://10.0.2.2:3128
ENV https_proxy http://10.0.2.2:3128
ENV NO_PROXY .hcuk.pri,10.0.2.2,192.168.99.,127.0.0.1
ENV no_proxy .hcuk.pri,10.0.2.2,192.168.99.,127.0.0.1
COPY . /go/src/github.com/a-h/rtrn

RUN go get -d -v github.com/a-h/rtrn
RUN go install github.com/a-h/rtrn

WORKDIR /go/src/github.com/a-h/rtrn
RUN go get -d -v

