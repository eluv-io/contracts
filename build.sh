#!/bin/bash
set -Eeuo pipefail

run_solc()
{
    bsc_name=$(basename ${1})
    $(solc ${1} --abi --hashes --optimize -o ${abi_dir} --overwrite)
    if [[ $? -eq 0 ]]; then
        echo -e "\n SUCCESS : ${bsc_name} ABI and function hashes present in ${abi_dir}"
    fi
}

# to get path to sol folder
sol_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

separator=
for i in {0..100}
do
  separator+="="
done

# check solc
if [[ ! -f /usr/local/bin/solc ]]
then
    echo "Error : solc is not found"
    exit 1
fi

# check the abigen is installed using HomeBrew
if [[ ! -f /usr/local/bin/abigen ]]
then
    echo "Error : abigen is not found, install ethereum using Homebrew"
    echo "Steps"
    echo "- brew tap ethereum/ethereum"
    echo "- brew install ethereum"
    exit 1
fi

out_dir="$( mkdir -p "$sol_dir/build" && cd "$sol_dir/build" && pwd )"
abi_dir=$sol_dir/dist

# abigen --sol ./base_content_space.sol --pkg=contracts --out build/base_content_space.go
$(abigen --sol ${sol_dir}/base_content_space.sol --pkg=contracts --out ${out_dir}/base_content_space.go)
if [[ $? -ne 0 ]]; then
    echo "FAILED : error occured while creating go binding!"
    exit 1
else
    echo "SUCCESS : The go binding for base_content_space.sol is present at ${out_dir}/base_content_space.go"
fi
echo -e "\n${separator}\n"


run_solc "${sol_dir}/base_content_space.sol"
echo -e "\n${separator}\n"
run_solc "${sol_dir}/lv_recording.sol"
echo -e "\n${separator}\n"
#run_solc "${sol_dir}/payment_service.sol"
#echo -e "\n${separator}\n"


$(go generate ./abi-parser)
if [[ $? -ne 0 ]]; then
    echo "FAILED : error occured while parsing abi"
    exit 1
else
    echo "SUCCESS : parsed ABI events present in ${abi_dir}"
fi
