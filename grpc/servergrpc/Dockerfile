FROM golang:1.22-alpine as imagen

RUN apk add --no-progress --no-cache gcc musl-dev
WORKDIR /server

COPY . .
RUN go build -tags musl -ldflags '-extldflags "-static"' -o ./main
FROM alpine
WORKDIR /app

COPY --from=imagen /server/main ./main

CMD ["/app/main"]