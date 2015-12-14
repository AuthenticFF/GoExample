FROM golang:latest


#build Go App
RUN go get github.com/AuthenticFF/GoExample

#if local changes are to be included
ADD . /go/src/github.com/AuthenticFF/GoExample
WORKDIR /go/src/github.com/AuthenticFF/GoExample

RUN go get
RUN go install

ENV PORT=9091
EXPOSE 9091

#init script
ADD ./init.sh /init.sh
RUN chmod a+x /init.sh
ENTRYPOINT /init.sh go run server.go
