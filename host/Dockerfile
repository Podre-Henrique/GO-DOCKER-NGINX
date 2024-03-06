FROM golang:1.21.4-alpine3.18 AS builder

WORKDIR /
COPY . .

# RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend

FROM gcr.io/distroless/base-debian12
WORKDIR /
COPY --from=builder /backend /backend  
USER nonroot:nonroot

# EXPOSE 8080
ENTRYPOINT ["/backend"]
