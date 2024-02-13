# syntax=docker/dockerfile:1@sha256:ac85f380a63b13dfcefa89046420e1781752bab202122f8f50032edf31be0021

# Build go binary
FROM golang:1.22-alpine@sha256:8e96e6cff6a388c2f70f5f662b64120941fcd7d4b89d62fec87520323a316bd9 as build

WORKDIR /go/src/app

# COPY --from=buf /root/.netrc /root/.netrc
ENV GOPRIVATE=buf.build/gen/go

COPY go.mod go.sum /
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" ./cmd/api-server

# Run
FROM gcr.io/distroless/static:nonroot@sha256:6a3500b086c2856fbc189f5d11351bdbcf7c4dc5673c2b6070aac9d607da90d7

COPY --from=build /go/bin/app .

EXPOSE 50051
ENTRYPOINT ["./app"]