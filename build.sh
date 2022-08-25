#!/bin/bash
set -Eeuo pipefail

# Must use:
# - solc 0.5.4
badreq=
solc --version | grep -q 0.5.4+commit.9549d8ff || badreq=true
if test $badreq; then
  echo "Must use abigen 1.9.19 and solc 0.5.4"
  exit
fi

## NOTE: this script does not generate go-bindings anymore.
##       Use build-go.sh to generate go-bindings

run_solc() {
    bsc_name=$(basename "${1}")

    if $(solc "${1}" --abi --hashes --optimize -o "${abi_dir}" --overwrite) ; then
        echo -e "\n SUCCESS : ${bsc_name} ABI and function hashes present in ${abi_dir}"
    fi

    if $(solc "${1}" --bin --optimize -o "${abi_dir}/bin" --overwrite) ; then
        echo -e "\n SUCCESS : ${bsc_name} BIN present in ${abi_dir}/bin"
    fi
}

# to get path to sol folder
sol_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
abi_dir=$sol_dir/dist

separator="====================================================================================================="

hash solc || {
    echo "Error : solc is not found"
    exit 1
}

#run_solc "${sol_dir}/base_content_space.sol"
#echo -e "\n${separator}\n"
#run_solc "${sol_dir}/lv_recording.sol"
#echo -e "\n${separator}\n"
#run_solc "${sol_dir}/payment_service.sol"
#echo -e "\n${separator}\n"
run_solc "${sol_dir}/tradable/elv_tradable_full.sol"
echo -e "\n${separator}\n"

pushd cmds > /dev/null 2>&1
go generate ./abi-parser/abi-parser.go
ret=$?
popd > /dev/null 2>&1
if [[ $? -ne 0 ]]; then
    echo "FAILED : error occurred while parsing abi"
    exit 1
fi

