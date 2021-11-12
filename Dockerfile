FROM golang:alpine AS build

ARG RELEASE_VERSION="dev"
ARG RELEASE_GIT_COMMIT="build"

WORKDIR /build
COPY . /build

RUN go build \
  -ldflags "-X github.com/taskmedia/nuScrape/pkg/http/rest.version=${RELEASE_VERSION}-${RELEASE_GIT_COMMIT}" \
  cmd/nuScrape/nuScrape.go

FROM alpine

RUN addgroup -S tm && \
  adduser -S tm -G tm

COPY --from=build /build/nuScrape /bin/nuScrape

USER tm
EXPOSE 8080

ENTRYPOINT [ "/bin/nuScrape" ]
