FROM golang:alpine
RUN apk add git
RUN go get github.com/jinzhu/gorm
RUN go get github.com/gorilla/mux
RUN apk add gcc musl-dev
RUN go get github.com/mattn/go-sqlite3
WORKDIR /go/src/url-shortener
ADD . .
RUN go install url-shortener
ENTRYPOINT /go/bin/url-shortener
EXPOSE 8080