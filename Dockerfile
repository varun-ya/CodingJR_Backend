FROM golang:1.21-alpine
WORKDIR /app
RUN apk add --no-cache bash
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main .
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh
EXPOSE 3000
ENTRYPOINT ["/wait-for-it.sh", "db", "3306", "--", "./main"]

