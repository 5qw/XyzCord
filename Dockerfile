FROM golang:latest

WORKDIR /go/src/XyzCord

COPY . ./

RUN go build .

VOLUME ["/root/.config/XyzCord"]

ENTRYPOINT ["/go/src/XyzCord/XyzCord"]

