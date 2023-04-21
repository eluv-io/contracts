#!/bin/bash
set -Eeuo pipefail

# This build script regenerates go bindings for go-ethereum 1.10.19
# NOTE:
# * abigen does not anymore invoke solc (solc output must be explicitly provided).


# build abigen
(
  echo "#### building abigen command"
  pushd cmds
  go build -o ../abigen ./abigen
  popd
)

source ./build-common.sh


# solc ./base_content_space.sol --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize | abigen --pkg=contracts --out build/base_content_space.go --combined-json -
abigen(){
    # 1: solc version
    # 2: contract file (.sol)
    # 3: package name
    # 4: output file - go bindings
    local version=$1
    solc_bin="${solc_folder}/${version}/solc"

    "${solc_bin}" "$2" --evm-version constantinople --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize | "$abigen_dir/abigen" --pkg=${3} --out "${4}" --combined-json -
    ###./abigen --sol "${sol_dir}/${2}" --pkg=${3} --out "${out_dir}/${5}"
    ret=$?
    if [[ $ret -ne 0 ]]; then
        echo "FAILED  : error creating go binding for ${2}!"
        exit 1
    else
        echo "SUCCESS : go binding for ${2} generated at ${4}"
    fi
}

# TEMPORARY / generate tenant factory binding directly

# shellcheck disable=SC2030
echo "#### generating latest contracts"
mkdir -p "contracts-go/tenant_factory"
abigen $solc_0_5_4 base_tenant_factory.sol tenant_factory contracts-go/tenant_factory/base_tenant_factory.go
exit 0
