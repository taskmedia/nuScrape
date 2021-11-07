FROM golang:alpine AS build

WORKDIR /build
COPY . /build

RUN go build cmd/nuScrape/nuScrape.go

FROM alpine

RUN addgroup -S tm && \
  adduser -S tm -G tm

COPY --from=build /build/nuScrape /bin/nuScrape

USER tm
EXPOSE 8080

ENTRYPOINT [ "/bin/nuScrape" ]
