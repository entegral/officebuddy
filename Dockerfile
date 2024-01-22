# syntax=docker/dockerfile:1

# build stage
FROM golang:1.21 AS build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor

# test stage
FROM build AS test
WORKDIR /app
RUN go test -v ./...

# final stage
FROM golang:1.21
WORKDIR /root/
COPY --from=build /app/officebuddy .
EXPOSE 8080
CMD ["./officebuddy"]