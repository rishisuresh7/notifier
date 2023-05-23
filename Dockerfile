FROM golang:1.19-alpine as buildEnv
ARG VERSION=0.0.0
RUN mkdir -p /opt/notifier
COPY . /opt/notifier/
WORKDIR /opt/notifier
RUN go build -ldflags="-X main.Version=${VERSION}" -o ./build/notifier ./apps/main/main.go

FROM alpine:latest
ENV PORT 9000
RUN mkdir -p /opt/notifier
WORKDIR /opt/notifier
COPY --from=buildEnv /opt/notifier/build/notifier ./notifier
# no server is run using this Dockerfile. This is being used as a placeholder for future use
EXPOSE $PORT
CMD ./notifier
