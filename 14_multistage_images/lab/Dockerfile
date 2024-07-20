# syntax=docker/dockerfile:1

FROM golang:1.22
WORKDIR /app
ENV USER="gopher"

# ########## useradd options ##########
# -m, --create-home  Create the user's home directory
# -G, --groups       GROUPS List of supplementary groups
# -r, --system       Create a system account

RUN addgroup app && \
    useradd -m -r gopher -G app

COPY --chown=gopher:app . .

RUN go mod init main \
    && \
    go mod tidy \
    && \
    go build -o hello

ENTRYPOINT ["/bin/sh", "-c"]
CMD [ "./hello | tee -a hello_output.txt"]


