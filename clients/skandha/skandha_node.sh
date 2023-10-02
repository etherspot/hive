#!/bin/sh

# Immediately abort the script on any error encountered
set -e

if [ -z "$HIVE_ERC4337_ETH1_RPC_ADDRS" ]; then
    echo "No Eth1 RPC addr was provided for the bundler."
    # TODO: Attempt to start a temp geth node
    exit 1
fi

echo Starting Skandha Bundler

node ./packages/cli/bin/skandha \
    node \
    --paramsFile=/hive/input/config.yaml \
