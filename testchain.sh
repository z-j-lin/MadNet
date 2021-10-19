#!/bin/sh


gnome-terminal --tab -e ./scripts/bootnode.sh 
./scripts/geth-local-snapshot-restore.sh
gnome-terminal --tab -e ./scripts/geth-local-resume.sh
for i in 0 1 2 3 4
do
    gnome-terminal --tab -e ./scripts/validator$i.sh
done 
./madnet --config=./assets/config/owner.toml utils deposit
./madnet --config=./assets/config/owner.toml utils deposit
./madnet --config=./assets/config/owner.toml utils deposit
./madnet --config=./assets/config/owner.toml utils deposit


