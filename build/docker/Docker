FROM golang:latest

WORKDIR /go/src/app

RUN go get github.com/oxequa/realize
RUN go get github.com/google/wire/cmd/wire
RUN go get -u github.com/go-delve/delve/cmd/dlv && \
  go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

EXPOSE 8080
EXPOSE 5002

CMD ["realize", "start", "--server"]
