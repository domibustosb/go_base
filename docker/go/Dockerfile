# Base image golang in https://hub.docker.com/_/golang
FROM golang:1.15.2 AS base

WORKDIR /go/src/github.com/domibustosb/go_base

EXPOSE 8080

ENV LANG en_US.UTF-8
ENV LC_ALL=C
ENV LANGUAGE en_US.UTF-8

FROM base AS dev
# application and watches for changes - ONLY development
ARG FRESHER_VERSION=1.2.1

RUN wget -c https://github.com/roger-russel/fresher/releases/download/v${FRESHER_VERSION}/fresher_${FRESHER_VERSION}_Linux_x86_64.tar.gz \
      && tar xzf fresher_${FRESHER_VERSION}_Linux_x86_64.tar.gz -C /go/bin/ \
      && rm -f fresher_*tar.gz

 CMD ["fresher", "-c", "./config/fresher.yaml"]

 EXPOSE 33333

# Configuration for build the application
FROM base AS build

COPY . .

RUN cd cmd/base/api && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s -extldflags "-static"' -a \
      -o /go/bin/api .

ENV USER=base

RUN addgroup --gid 1001 --system $USER && adduser -u 1001 --system $USER --gid 1001
RUN chown $USER:$USER /go/bin/api -R

# Final image for production
FROM scratch AS final

COPY --from=build  /go/bin/api /go/bin/api

USER $USER:$USER

ENTRYPOINT ["/go/bin/api"]