FROM --platform=$BUILDPLATFORM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o main .

FROM alpine:3

WORKDIR /app
COPY --from=builder /app/main main

ENV GIN_MODE=release
ENV PORT=8080

EXPOSE 8080

CMD ["./main"]
