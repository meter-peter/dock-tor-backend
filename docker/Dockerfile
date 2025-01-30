FROM golang:1.22.2-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o clinic-app ./main.go

EXPOSE 8080

CMD ["./clinic-app"]