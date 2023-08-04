FROM golang:1.18-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/crawler
# COPY ../go.mod ./
# COPY ../go.sum ./
ADD . .

RUN go env -w GO111MODULE=on
RUN go mod download

# Unit test
#RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/crawler cmd/server/main.go

#start runner
FROM alpine:3.9
RUN apk add ca-certificates
WORKDIR /app/crawler
COPY --from=build_base /tmp/crawler/out/crawler .
COPY --from=build_base /tmp/crawler/config.yml .

CMD ["./crawler"]