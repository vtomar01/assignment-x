FROM golang:latest

WORKDIR /app

COPY src/main/resources/local.json /app/local.json
COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/src/main
RUN go build -o server .

EXPOSE 4001

CMD ["./server"]