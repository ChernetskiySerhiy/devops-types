# Етап 1: компілюємо Go-програму
FROM golang:1.27 AS builder
WORKDIR /build
COPY go.mod ./
COPY src/ ./src/
RUN CGO_ENABLED=0 go build -o /app ./src

# Етап 2: порожній образ лише з програмою і сторінкою
FROM scratch
COPY --from=builder /app /app
COPY html /html
EXPOSE 8080
ENTRYPOINT ["/app", "/html"]
