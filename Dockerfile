###########
# BUILDER #
###########

FROM golang:1.15 as builder

# set work directory
WORKDIR /usr/src/app

COPY . .

# set build environments and build the app
ARG GOOS=linux
ARG GOARCH=$TARGETARCH
ARG CGO_ENABLED=0
RUN go build -a -installsuffix cgo -o bin/trex
RUN chmod +x bin/trex

# create non-privelleged user for running the app
RUN addgroup app && adduser --group app --no-create-home --disabled-login --disabled-password app
RUN chown -R app:app /trex_exporter


#########
# FINAL #
#########

FROM scratch

# set label for Docker image repo
LABEL org.opencontainers.image.source="https://github.com/starcatmeow/trex_exporter"

# create directory for the app and set work dircetory
WORKDIR /trex_exporter
USER app

# copy built files
COPY --from=builder /usr/src/app/bin .
COPY --from=builder /etc/passwd /etc/


# define environments
ENV TREX_EXPORTER_PORT=9788
ENV TREX_EXPORTER_BIND_ADDRESS=0.0.0.0
ENV TREX_MINER_URL=http://localhost:4057
ENV TREX_WORKER_NAME=trex

# run the app
CMD ./trex --api-address=$TREX_MINER_URL --web.listen-address=$TREX_EXPORTER_BIND_ADDRESS:$TREX_EXPORTER_PORT --worker=$TREX_WORKER_NAME
