# syntax=docker/dockerfile:1@sha256:ac85f380a63b13dfcefa89046420e1781752bab202122f8f50032edf31be0021

# Build go binary
FROM golang:1.22-alpine@sha256:298646364548cc5e1372e2612a6a2aaa53d44bed00284ecbc89b7aa8a83ad602 as build

WORKDIR /go/src/app

# COPY --from=buf /root/.netrc /root/.netrc
ENV GOPRIVATE=buf.build/gen/go

COPY go.mod go.sum /
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" ./cmd/api-server

# Run
FROM gcr.io/distroless/static:nonroot@sha256:112a87f19e83c83711cc81ce8ed0b4d79acd65789682a6a272df57c4a0858534

COPY --from=build /go/bin/app .

EXPOSE 50051
ENTRYPOINT ["./app"]