FROM golang:alpine AS build

WORKDIR /go/src/github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/challenge5/producer cmd/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/challenge5/producer /go/bin/challenge5/producer
ENTRYPOINT ["/go/bin/challenge5/producer"]
