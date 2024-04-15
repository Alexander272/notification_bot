FROM golang:1.21.9-alpine3.19 as builder
WORKDIR /build
COPY ./go.mod . 
RUN go mod download
COPY . .
RUN go build -o main cmd/app/main.go

FROM alpine:3.18
RUN apk add --no-cache tzdata
COPY ./configs/ /configs/
COPY --from=builder /build/main /bin/main
ENTRYPOINT ["/bin/main"]