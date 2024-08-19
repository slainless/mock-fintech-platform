# setup base for building, tooling setup, or running dev server
FROM golang:latest AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

FROM base AS builder
COPY . .
RUN go generate ./...

# setup delve
FROM builder AS delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# install user
FROM builder AS user
RUN go build -o ./bin/user ./cmd/user
RUN stat /app/bin/user

# install payment
FROM builder AS payment
RUN go build -o ./bin/payment ./cmd/payment
RUN stat /app/bin/payment

# dev
FROM delve AS dev
ENTRYPOINT ["dlv", "debug", "--headless", "--api-version=2", "--accept-multiclient", "--continue"]

# prod
FROM golang:latest AS production
WORKDIR /usr/local/bin
COPY --from=user /app/bin/user .
COPY --from=payment /app/bin/payment .
CMD ["echo", "Try user or payment."]