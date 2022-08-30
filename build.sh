#!/bin/bash
set -Eeuo pipefail


source ./build-common.sh

## NOTE: this script does not generate go-bindings anymore.
##       Use build-go.sh to generate go-bindings

run_solc() {
    local version=$1
    solc_bin="${solc_folder}/${version}/solc"

    bsc_name=$(basename "${2}")

    if $($solc_bin "${2}" --abi --hashes --optimize -o "${abi_dir}" --overwrite) ; then
        echo -e "\n SUCCESS : ${bsc_name} ABI and function hashes present in ${abi_dir}"
    fi

    if $($solc_bin "${2}" --bin --optimize -o "${abi_dir}/bin" --overwrite) ; then
        echo -e "\n SUCCESS : ${bsc_name} BIN present in ${abi_dir}/bin"
    fi
}

# to get path to sol folder
base_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
abi_dir=$base_dir/dist

separator="====================================================================================================="



run_solc $solc_0_5_4 "${base_dir}/base_content_space.sol"
echo -e "\n${separator}\n"

run_solc $solc_0_5_4 "${base_dir}/tradable/elv_tradable_full.sol"
echo -e "\n${separator}\n"

pushd cmds > /dev/null 2>&1
go generate ./abi-parser/abi-parser.go
ret=$?
popd > /dev/null 2>&1
if [[ $? -ne 0 ]]; then
    echo "FAILED : error occurred while parsing abi"
    exit 1
fi
