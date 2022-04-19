#!/bin/bash
set -Eeuo pipefail

run_solc() {
    bsc_name=$(basename "${1}")
    if solc "${1}" --abi --hashes --optimize -o "${abi_dir}" --overwrite; then
        echo -e "\n SUCCESS : ${bsc_name} ABI and function hashes present in ${abi_dir}"
    fi
}

# to get path to sol folder
sol_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

separator="====================================================================================================="

hash solc || {
    echo "Error : solc is not found"
    exit 1
}

# we should also check the version here
hash abigen || {
    echo "Error : abigen is not found, install ethereum using Homebrew"
    echo "Steps"
    echo "- brew tap ethereum/ethereum"
    echo "- brew install ethereum"
    exit 1
}

out_dir="$( mkdir -p "$sol_dir/build" && cd "$sol_dir/build" && pwd )"
abi_dir=$sol_dir/dist

# abigen --sol ./base_content_space.sol --pkg=contracts --out build/base_content_space.go
if abigen --sol "${sol_dir}/base_content_space.sol" --pkg=contracts --out "${out_dir}/base_content_space.go"; then
    echo "SUCCESS : The go binding for base_content_space.sol is present at ${out_dir}/base_content_space.go"
else
    echo "FAILED : error occured while creating go binding!"
    exit 1
fi
echo -e "\n${separator}\n"

run_solc "${sol_dir}/base_content_space.sol"
echo -e "\n${separator}\n"
run_solc "${sol_dir}/lv_recording.sol"
echo -e "\n${separator}\n"
run_solc "${sol_dir}/payment_service.sol"
echo -e "\n${separator}\n"
run_solc "${sol_dir}/tradable/elv_tradable_full.sol"
echo -e "\n${separator}\n"

if go generate ./cmds/abi-parser; then
    echo "SUCCESS : parsed ABI events present in ${abi_dir}"
else
    echo "FAILED : error occured while parsing abi"
    exit 1
fi
