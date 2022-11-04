#!/bin/sh

mkdir -p /root/log
sixd start --rpc.laddr tcp://0.0.0.0:26657 --trace
