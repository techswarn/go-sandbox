FROM golang:1.20-alpine as builder
WORKDIR /app

COPY go.mod ./
COPY . ./

RUN go build -v -o bin/worker

FROM alpine
# Install any required dependencies.
RUN apk --no-cache add ca-certificates

WORKDIR /root/
RUN echo $(ls)
COPY --from=builder /app/bin/worker /usr/local/bin/

CMD ["worker"]