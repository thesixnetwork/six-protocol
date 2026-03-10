FROM golang:1.25.7-bookworm AS go-builder

RUN apt-get update && apt-get install -y \
  ca-certificates \
  build-essential \
  libudev-dev \
  git \
  make \
  jq \
  curl \
  binutils-gold \
  --no-install-recommends \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/github.com/thesixnetwork/six-protocol
COPY . /go/src/github.com/thesixnetwork/six-protocol/

RUN LEDGER_ENABLED=false BUILD_TAGS=muslc make build

FROM cgr.dev/chainguard/wolfi-base

WORKDIR /root
COPY --from=go-builder /go/src/github.com/thesixnetwork/six-protocol/build/sixd /usr/bin/sixd

# install necessary libraries for binary to run
RUN apk upgrade --no-cache && apk add bash git libgcc jq curl tzdata

# Set timezone
ENV TZ=Asia/Bangkok

# RUN mkdir -p /root/.six
COPY docker/* /opt/
RUN chmod +x /opt/*.sh

WORKDIR /opt

# Blockchain API
EXPOSE 1317
# Tendermint p2p
EXPOSE 26656
# Tendermint node
EXPOSE 26657


CMD ["sixd"]
