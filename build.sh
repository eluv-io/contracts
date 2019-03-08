#!/bin/sh

# to get path to sol folder
sol_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# check the abigen is installed using HomeBrew
if [ ! -f /usr/local/bin/abigen ]
then
echo "Error : abigen is not found, install ethereum using Homebrew"
echo "Steps"
echo "- brew tap ethereum/ethereum"
echo "- brew install ethereum"
exit 1
fi

out_dir="$( mkdir -p "$sol_dir/build" && cd "$sol_dir/build" && pwd )"

# abigen --sol ./base_content_space.sol --pkg=contracts --out build/base_content_space.go
$(abigen --sol $sol_dir/base_content_space.sol --pkg=contracts --out $out_dir/base_content_space.go)
if [ $? -ne 0 ]; then
echo "error occured while creating go binding!"
else
echo "The go binding for base_content_space.sol is present at $out_dir/base_content_space.go"
fi




