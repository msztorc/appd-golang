FROM golang

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go env -w GO111MODULE=off
#RUN go mod init appd

ADD golang-sdk-x64-linux-4.5.2.0.tar $GOPATH/src

RUN go build -o main .
EXPOSE 8080
CMD ["/app/main"]
