
FROM golang:1.22-alpine AS builder
WORKDIR /build

COPY . .

RUN go mod download
RUN go build -o ./userapi ./userapi/main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/userapi ./userapi

CMD [ "/app/userapi" ]
