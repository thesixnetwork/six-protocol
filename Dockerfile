# docker build . -t six-protocol/sixd:latest
# docker run --rm -it cosmoscontracts/sixd:latest /bin/sh
FROM golang:1.24-bookworm AS go-builder

# Install build dependencies using apt-get for a Debian-based system.
# build-essential includes gcc, make, and other core build tools.
# libudev-dev is the Debian equivalent of eudev-dev for Cgo and Ledger.
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

# Install cosmovisor - use local binary if exists, otherwise download
RUN if [ -f "./build/bin/cosmovisor" ]; then \
  echo "Using local cosmovisor binary"; \
  cp ./build/bin/cosmovisor /go/bin/cosmovisor; \
  else \
  echo "Local cosmovisor not found, installing from source"; \
  go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@latest; \
  fi

RUN LEDGER_ENABLED=false BUILD_TAGS=muslc make build

FROM cgr.dev/chainguard/wolfi-base

WORKDIR /root
COPY --from=go-builder /go/src/github.com/thesixnetwork/six-protocol/build/sixd /usr/bin/sixd
COPY --from=go-builder /go/bin/cosmovisor /usr/bin/cosmovisor

RUN chmod +x /usr/bin/cosmovisor
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
