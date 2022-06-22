#!/bin/bash
set -ex

NETWORK=${1:-"dev"}
CURRENT_WD=$PWD
BRIDGE_DIR=./bridge

cd $BRIDGE_DIR

if [[ -z "${FACTORY_ADDRESS}" ]]; then
    
    if [[ -z "0xA85Fcfba7234AD28148ebDEe054165AeF6974a65" ]]; then
        echo "It was not possible to find Factory Address in the environment variable FACTORY_ADDRESS! Exiting script!"
        exit 1
    fi
fi

npx hardhat --network "$NETWORK" --show-stack-traces initializeEthdkg --factory-address "$FACTORY_ADDRESS"

cd "$CURRENT_WD"
