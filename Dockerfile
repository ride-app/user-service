# syntax=docker/dockerfile:1@sha256:ac85f380a63b13dfcefa89046420e1781752bab202122f8f50032edf31be0021

# Build go binary
FROM golang:1.22-alpine@sha256:fa4add5ca88c1dfddec8fe3fa57c2956f318f41fd224382a7d91388c4b6929c9 as build

WORKDIR /go/src/app

# COPY --from=buf /root/.netrc /root/.netrc
ENV GOPRIVATE=buf.build/gen/go

COPY go.mod go.sum /
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" ./cmd/api-server

# Run
FROM gcr.io/distroless/static:nonroot@sha256:55c636171053dbc8ae07a280023bd787d2921f10e569f3e319f1539076dbba11

COPY --from=build /go/bin/app .

EXPOSE 50051
ENTRYPOINT ["./app"]