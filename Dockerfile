FROM golang:1.25-bookworm

WORKDIR /app

COPY . ./

RUN go mod tidy
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/exe .
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY migrations ./migrations
RUN go build -o /app/exe main.go

CMD ["/app/exe"]