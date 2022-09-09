# Build the Go Binary.
FROM golang:1.19 as build_engine
ENV CGO_ENABLED 0
ARG BUILD_REF

# Copy the source code into the container.
COPY . /engine

# Build the service binary.
WORKDIR /service/app/services/engine
RUN go build -ldflags "-X main.build=${BUILD_REF}"


# Run the Go Binary in Alpine.
FROM alpine:3.16
ARG BUILD_DATE
ARG BUILD_REF
RUN addgroup -g 1000 -S engine && \
    adduser -u 1000 -h /engine -G engine -S engine
COPY --from=build_engine --chown=engine:engine /service/zarf/keys/. /service/zarf/keys/.
COPY --from=build_engine --chown=engine:engine /service/app/services/engine/engine /engine/engine
WORKDIR /engine
USER engine
CMD ["./engine"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="engine" \
      org.opencontainers.image.authors="William Kennedy <bill@ardanlabs.com>" \
      org.opencontainers.image.source="https://github.com/ardanlabs/liarsdice/app/services/engine" \
      org.opencontainers.image.revision="${BUILD_REF}" \
      org.opencontainers.image.vendor="Ardan Labs"