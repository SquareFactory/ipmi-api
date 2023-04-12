# ------------------------------------
FROM golang:latest as api-builder
# ------------------------------------

WORKDIR /work 
COPY . ./
RUN make build-all 

# ------------------------------------
FROM docker.io/library/alpine:edge
# ------------------------------------

RUN apk add --no-cache ipmitool
WORKDIR /app
COPY --from=api-builder /work/bin /app/
ENTRYPOINT [ "app/ipmi-api" ]