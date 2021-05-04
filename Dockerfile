FROM golang
MAINTAINER Miroslaw Sztorc

RUN mkdir /app
ADD . /app/
WORKDIR /app

#RUN export GOFLAGS=-mod=vendor
#RUN export GO111MODULE=on
RUN go mod init appd

ADD golang-sdk-x64-linux-4.5.2.0.tar /usr/local/go/src

RUN go build -o main .
EXPOSE 8080
CMD ["/app/main"]
