FROM golang:alpine
COPY . /app
WORKDIR /app
RUN go build ./cmd/server
RUN go build ./cmd/client