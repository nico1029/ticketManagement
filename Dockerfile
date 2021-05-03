FROM golang:alpine as builder 
COPY go.mod go.sum /go/src/github.com/nico1029/ticketManagement/
WORKDIR /go/src/github.com/nico1029/ticketManagement
RUN go mod download
COPY . /go/src/github.com/nico1029/ticketManagement
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/ticketManagement github.com/nico1029/ticketManagement

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/nico1029/ticketManagement/build/ticketManagement /usr/bin/ticketManagement
EXPOSE 8082 8082
ENTRYPOINT ["/usr/bin/ticketManagement"]