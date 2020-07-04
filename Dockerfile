############################
# STEP 1 build executable binary
############################
# Compile stage
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/Cheunn-Panaa/misterv-discord/
COPY . .

# Fetch dependencies.
RUN go get -d -v

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s -extldflags "-static"' -a \
      -o /go/bin/app .

############################
# STEP 2 build a small image
############################
FROM alpine

# Copy our static executable
COPY --from=builder /go/bin/app /go/bin/app

RUN apk add --no-cache ca-certificates
# Run the hello binary.
ENTRYPOINT ["/go/bin/app"]
