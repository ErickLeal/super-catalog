FROM golang:1.24-alpine
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download && go install github.com/air-verse/air@latest
COPY . .
CMD ["air"]
