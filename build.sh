#!/bin/bash
set -Eeuo pipefail

run_solc()
{
    bsc_name=$(basename "${1}")
    solc "${1}" --abi --hashes --optimize -o "${abi_dir}" --overwrite
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

# build abigen
pushd cmds > /dev/null 2>&1
go build -o ../abigen ./abigen
ret=$?
popd > /dev/null 2>&1
if [[ $ret -ne 0 ]]; then
    echo "Error : building abigen failed"
    exit 1
fi

out_dir="$( mkdir -p "$sol_dir/build" && cd "$sol_dir/build" && pwd )"
abi_dir=$sol_dir/dist

# report current version of contracts
branch=$(git rev-parse --abbrev-ref HEAD)
revision=$(git rev-parse HEAD)
echo "$branch@$revision" > "${out_dir}"/contracts_version.txt

# abigen --sol ./base_content_space.sol --pkg=contracts --out build/base_content_space.go
abigen(){
    ./abigen --sol "${sol_dir}/${1}" --pkg=${2} --out "${out_dir}/${3}"
    ret=$?
    if [[ $ret -ne 0 ]]; then
        echo "FAILED  : error creating go binding for ${1}!"
        exit 1
    else
        echo "SUCCESS : go binding for ${1} generated at ${out_dir}/${3}"
    fi
}

abigen base_content_space.sol contracts base_content_space.go
abigen tradable/elv_tradable_full.sol elv_tradable tradable/elv_tradable_full.go


rm ./abigen
echo -e "\n${separator}\n"

run_solc "${sol_dir}/base_content_space.sol"
echo -e "\n${separator}\n"
run_solc "${sol_dir}/lv_recording.sol"
echo -e "\n${separator}\n"
run_solc "${sol_dir}/payment_service.sol"
echo -e "\n${separator}\n"
run_solc "${sol_dir}/tradable/elv_tradable_full.sol"
echo -e "\n${separator}\n"

pushd cmds > /dev/null 2>&1
go generate ./abi-parser/abi-parser.go
ret=$?
popd > /dev/null 2>&1
if [[ $? -ne 0 ]]; then
    echo "FAILED : error occurred while parsing abi"
    exit 1
else
    echo "SUCCESS : parsed ABI events present in ${abi_dir}"
fi
