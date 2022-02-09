FROM golang:alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN apk add build-base
RUN go mod download
COPY . .
RUN go build -o engine .

FROM alpine
WORKDIR /app
COPY --from=builder /app/engine .
CMD [ "./engine" ]