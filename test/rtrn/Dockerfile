FROM golang:1.6
ENV GOPATH /go

# Dependencies we know we need, to make the build faster
RUN go get gopkg.in/tomb.v2 \
    && go get github.com/gorilla/mux \
    && go get gopkg.in/mgo.v2 \
    && go get github.com/go-sql-driver/mysql

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /
RUN chmod +x /wait-for-it.sh

COPY . /go/src/github.com/a-h/rtrn

WORKDIR /go/src/github.com/a-h/rtrn
RUN go get -d -v ...

CMD /wait-for-it.sh mariadb:3306 -t 60 && /wait-for-it.sh mongo:27017 -t 60 && go test github.com/a-h/rtrn/...
