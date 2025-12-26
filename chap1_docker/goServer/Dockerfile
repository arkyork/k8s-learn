FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go build -o main .

FROM scratch
COPY --from=builder /app/main /main
ENTRYPOINT [ "/main" ]