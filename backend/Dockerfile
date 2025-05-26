# syntax=docker/dockerfile:1

#base image
FROM golang:1.24-alpine

#to create a working dir
WORKDIR /app

#to download dependencies
COPY go.mod go.sum ./

#to download dependencies
RUN go mod download

#copying files into image
COPY . .

#building the go binary
RUN go build -o main .

#port
EXPOSE 5000

#running server
CMD ["./main"]
