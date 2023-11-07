# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /app

RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor
EXPOSE 8080

CMD ["./officebuddy"]