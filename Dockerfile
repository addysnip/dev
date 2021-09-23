FROM golang:1.16-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN ls -l && \
    CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM alpine:latest
RUN mkdir -p /app/static
COPY --from=builder /app/app /app/sso
ADD static /app/static
WORKDIR /app
ENTRYPOINT [ "./app", "run" ]