FROM golang:alpine
RUN apk add --no-cache git
WORKDIR /app
COPY go.* .
COPY *air.toml .
RUN go mod download
COPY . .
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]