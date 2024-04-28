FROM golang:1.22-alpine as imagen

RUN apk add --no-progress --no-cache gcc musl-dev
WORKDIR /consumer

COPY . .
RUN go build -tags musl -ldflags '-extldflags "-static"' -o ./main
FROM alpine
WORKDIR /app

COPY --from=imagen /consumer/main ./main

CMD ["/app/main"]