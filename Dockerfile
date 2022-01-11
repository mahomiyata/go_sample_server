FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
COPY ./db ./db
COPY ./entity ./entity

RUN go build -o ./sample-server

EXPOSE 8080

CMD [ "./sample-server" ]