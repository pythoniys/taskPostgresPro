FROM golang:1.22

WORKDIR /usr/local/src

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
#COPY go.mod go.sum ./
#RUN go mod download

COPY cmd/app ./
RUN go build -o ./bin/app main.go