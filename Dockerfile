# syntax=docker/dockerfile:1@sha256:ac85f380a63b13dfcefa89046420e1781752bab202122f8f50032edf31be0021

# Build go binary
FROM golang:1.22-alpine@sha256:0466223b8544fb7d4ff04748acc4d75a608234bf4e79563bff208d2060c0dd79 as build

WORKDIR /go/src/app

# COPY --from=buf /root/.netrc /root/.netrc
ENV GOPRIVATE=buf.build/gen/go

COPY go.mod go.sum /
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" ./cmd/api-server

# Run
FROM gcr.io/distroless/static:nonroot@sha256:f41b84cda410b05cc690c2e33d1973a31c6165a2721e2b5343aab50fecb63441

COPY --from=build /go/bin/app .

EXPOSE 50051
ENTRYPOINT ["./app"]