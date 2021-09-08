FROM golang:1.17.0

ENV GO111MODULE=on

COPY api/ /go/api/app

WORKDIR /go/api/app

# cosmtrek/installation of air
RUN go get -u github.com/cosmtrek/air
#installation of delve
RUN go get -u github.com/go-delve/delve/cmd/dlv

#Start air
CMD air -c .air.toml
