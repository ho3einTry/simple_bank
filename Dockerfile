# Build Stage
FROM golang:1.21.3-bullseye AS builder
LABEL authors="Hossein"
WORKDIR /app
COPY . .
RUN go build -o simpleBank main.go


# Run Stage
FROM debian:bullseye
WORKDIR /app
COPY --from=builder /app/simpleBank .
COPY  app.env .
EXPOSE 8070
CMD ["/app/simpleBank"]