FROM golang:1.15.6-alpine as builder
ENV GO111MODULE=on
RUN apk add --no-cache ca-certificates git
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/connect-database .
CMD ["./connect-database"]
