##################################################
# Stage 1: Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY main.go .
ENV GO111MODULE auto

RUN  go mod init main && \
     go mod tidy && \
     go build -o app

##################################################
# Stage 2: Final stage
FROM alpine:3.19 AS final
WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/app ./

ENTRYPOINT ["/bin/sh", "-c"]
CMD [ "./app"]
##################################################
