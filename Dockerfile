FROM golang:alpine as rate-finder
ENV GO111MODULE=on
WORKDIR /server
COPY go.mod /server/
COPY go.sum /server/
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -mod=readonly -o /go/bin/rate-finder

FROM scratch
COPY --from=rate-finder /go/bin/rate-finder /go/bin/rate-finder
EXPOSE 8080
ENTRYPOINT ["/go/bin/rate-finder"]