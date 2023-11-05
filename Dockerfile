FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/homepage-bridge ./cmd/homepage-bridge


FROM alpine:3 AS runtime

WORKDIR /app

COPY --from=builder /app/homepage-bridge /app/homepage-bridge

ARG BUILD_DATE
ARG VCS_REF
ARG VCS_URL
ARG VERSION

LABEL org.opencontainers.image.version=$VERSION \
      org.opencontainers.image.title="Homepage Bridge" \
      org.opencontainers.image.description="A custom request proxy for my homepage dashboard" \
      org.opencontainers.image.authors="Julien W <cefadrom1@gmail.com>" \
      org.opencontainers.image.url=$VCS_URL \
      org.opencontainers.image.source=$VCS_URL \
      org.opencontainers.image.revision=$VCS_REF \
      org.opencontainers.image.created=$BUILD_DATE

EXPOSE 8080

CMD ["/app/homepage-bridge"]
