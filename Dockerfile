FROM golang:1.19-alpine

WORKDIR /usr/src/app/
EXPOSE 7000

COPY go.mod /usr/src/app/
COPY go.sum /usr/src/app/

RUN go mod tidy -v

COPY ./cmd /usr/src/app/
COPY . /usr/src/app/

RUN go build -o /go-server

CMD ["/go-server"]