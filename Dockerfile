# Builder

ARG GITHUB_PATH=github.com/denlipov/com-request-facade

FROM golang:1.16-alpine AS builder
RUN apk add --update make git
#protoc protobuf protobuf-dev curl
COPY . /home/${GITHUB_PATH}
WORKDIR /home/${GITHUB_PATH}
#RUN make deps-go && make build-go
RUN make build

# facade Service

FROM alpine:latest as server
LABEL org.opencontainers.image.source https://${GITHUB_PATH}
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /home/${GITHUB_PATH}/bin/facade .
COPY --from=builder /home/${GITHUB_PATH}/config.yml .
COPY --from=builder /home/${GITHUB_PATH}/migrations ./migrations

RUN chown root:root facade

#EXPOSE 50051
#EXPOSE 8080
#EXPOSE 8082
#EXPOSE 9100

CMD ["./facade"]
