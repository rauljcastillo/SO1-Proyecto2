FROM golang:1.22-alpine as imagen


WORKDIR /client

COPY . .
RUN go build -o ./main

FROM alpine
WORKDIR /app
COPY --from=imagen /client/main ./main
RUN apk --no-cache add tzdata
EXPOSE 8080
CMD ["/app/main"]