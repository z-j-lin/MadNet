#!/bin/sh

CONFIG=${1:-./scripts/base-files/integrationBootnode.toml}

./madnet --config "$CONFIG" bootnode