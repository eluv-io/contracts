#!/bin/sh

# to get path to sol folder
sol_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

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
else
echo "SUCCESS : The go binding for base_content_space.sol is present at $out_dir/base_content_space.go"
fi

out="$(solc ${sol_dir}/base_content_space.sol --abi --hashes --optimize -o ${abi_dir} --overwrite)"
if [[ $? -ne 0 ]]; then
echo "FAILED : error occured while running solc"
echo ${out}
else
echo "SUCCESS : ABI and function hashes present in ${abi_dir}"
fi

$(go generate ./abi-parser)
if [[ $? -ne 0 ]]; then
echo "FAILED : error occured while parsing abi"
else
echo "SUCCESS : parsed ABI events present in ${abi_dir}"
fi
