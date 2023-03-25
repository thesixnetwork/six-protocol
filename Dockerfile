# docker build . -t six-protocol/sixd:latest
# docker run --rm -it cosmoscontracts/sixd:latest /bin/sh
FROM golang:1.18-alpine3.15 AS go-builder
ARG GITHUB_TOKEN

# this comes from standard alpine nightly file
#  https://github.com/rust-lang/docker-rust-nightly/blob/master/alpine3.12/Dockerfile
# with some changes to support our toolchain, etc
RUN set -eux; apk add --no-cache ca-certificates build-base;

#libc-dev, gcc, linux-headers, eudev-dev are used for cgo and ledger installation
RUN apk upgrade --no-cache && apk add bash git make libgcc libc-dev gcc linux-headers eudev-dev jq curl

WORKDIR /go/src/github.com/thesixnetwork/six-protocol
COPY . /go/src/github.com/thesixnetwork/six-protocol/

# install comovisor
# RUN go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@latest
COPY ./build/bin/cosmovisor /go/bin/cosmovisor

RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
RUN export GOPRIVATE=github.com/thesixnetwork/sixnft 
RUN go get github.com/thesixnetwork/sixnft@v0.8.1-0.20230309050716-72b3d1a3671b

# force it to use static lib (from above) not standard libgo_cosmwasm.so file
RUN LEDGER_ENABLED=false BUILD_TAGS=muslc make build
# RUN file /go/src/github.com/thesixnetwork/six-protocol/build/sixd

## Final image for running sixd
# --------------------------------------------------------
FROM alpine:3.15

WORKDIR /root
COPY --from=go-builder /go/src/github.com/thesixnetwork/six-protocol/build/sixd /usr/bin/sixd
COPY --from=go-builder /go/bin/cosmovisor /usr/bin/cosmovisor

RUN chmod +x /usr/bin/cosmovisor
# install necessary libraries for binary to run
RUN apk upgrade --no-cache && apk add bash git libgcc jq curl tzdata

# Set timezone
ENV TZ Asia/Bangkok

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