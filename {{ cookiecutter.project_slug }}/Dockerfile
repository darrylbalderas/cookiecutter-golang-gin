FROM golang:{{ cookiecutter.golang_version }} AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM gcr.io/distroless/static-debian11
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]