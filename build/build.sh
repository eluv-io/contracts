#!/bin/sh

read -p "Enter sol_file:" fsol
read -p "Enter output file name:" fout #Exmaple: content.go

# to get path to contracts folder
contracts_dir="$( cd "$( dirname "${BASH_SOURCE[1]}" )" && pwd )"

# to get path to myworkspace
ws_dir="$( cd "$( dirname $contracts_dir)" && pwd)"

# check if the GOPATH is set to content-fabric
gopath_dir="$( cd "$ws_dir/content-fabric" && pwd )"
if [[ $GOPATH != *"$ws_dir/content-fabric"* ]]; then
export GOPATH=$GOPATH:$gopath_dir
fi

# generate abigen
if [ ! -f $gopath_dir/bin/abigen ] # check if abigen exists
then
cd $gopath_dir/src/eluvio/vendor/github.com/ethereum/go-ethereum/cmd/abigen
go install .
fi

if [ ! -d $contracts_dir/abigen ]
then
mkdir $contracts_dir/abigen
fi


# use abigen
cd $gopath_dir
$(./bin/abigen --sol $contracts_dir/$fsol --pkg=contracts --out $contracts_dir/abigen/$fout)
if [ $? -ne 0 ]; then
echo "error occured while creating go binding!"
else
echo "The go binding for $fsol is present at $contracts_dir/abigen/$fout"
fi





