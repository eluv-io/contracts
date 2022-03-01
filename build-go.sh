#!/bin/bash
set -Eeuo pipefail

## WARNING: these are 'manual steps'
##          specifically the correct version of solc needs to be in the path
##          as specified by comments

# check solc
if [[ ! -f /usr/local/bin/solc ]]
then
    echo "Error : solc is not found"
    exit 1
fi

# build abigen
pushd cmds > /dev/null 2>&1
go build -o ../abigen ./abigen
ret=$?
popd > /dev/null 2>&1
if [[ $ret -ne 0 ]]; then
    echo "Error : building abigen failed"
    exit 1
fi

mkdir -p tmp > /dev/null 2>&1
cd tmp
git clone https://github.com/eluv-io/contracts.git
cd contracts/

# recent contracts require solc 0.5.4
../../abigen --sol base_content_space.sol --pkg contracts --out ../../contracts-go/contracts/base_content_space.go
../../abigen --sol tradable/elv_tradable_full.sol  --pkg elv_tradable --out ../../contracts-go/tradable/elv_tradable_full.go

# older contracts require solc 0.4.24
git checkout 53ee85a858257d32c9650ff76f50e5ee2a1379d9
../../abigen --sol base_content_space.sol --pkg contracts_20190331 --out ../../contracts-go/contracts_20190331/base_content_space.go
git checkout 5ed6b64537175315be6c79fdbfb9a3d35086344d
../../abigen --sol base_content_space.sol --pkg contracts_20200206 --out ../../contracts-go/contracts_20200206/base_content_space.go
git checkout 2536dfb9b29e394feedcc1c5c7d16f67d21a35bb
../../abigen --sol base_content_space.sol --pkg contracts_20200803 --out ../../contracts-go/contracts_20200803/base_content_space.go
