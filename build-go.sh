#!/bin/bash
set -Eeuo pipefail

#
# This build script regenerates go bindings for go-ethereum 1.10.19
# NOTES:
# * abigen does not anymore invoke solc (solc output must be explicitly provided).
# * unless a solc binary is found in folder '_bin/${platform}/solc/${version},
#   the script downloads the solc binary from 'binaries.soliditylang.org'
# * to add a new version: declare a new variable with the value of the version
#   and add it to the 'solc_versions' variable.


# build abigen
(
  echo "#### building abigen command"
  pushd cmds
  go build -o ../abigen ./abigen
  popd
)


#
# versions of solc that we need
#
solc_0_4_24="0.4.24"
solc_0_5_4="0.5.4"
solc_versions=(${solc_0_4_24} ${solc_0_5_4})


curr_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
bin_folder="_bin"
abigen_dir=$(realpath "$curr_dir")

platform=""
case $(uname) in
    Darwin*)
        platform=macosx-amd64
        ;;
    Linux*)
        platform=linux-amd64
        ;;
    *)
        echo "Not sure what environment this is.  Cannot setup solc"
        exit 1
        ;;
esac


solc_folder=${curr_dir}/${bin_folder}/${platform}/solc
mkdir -p ${solc_folder}

downloadSolc() {
  local version=$1
  local jq_cmd=".releases[\"${version}\"]"
  local binary

  # cat solc_macosx-amd64_list.json | jq '.releases["0.3.6"]'
  binary=$( curl -q -s "https://raw.githubusercontent.com/ethereum/solc-bin/gh-pages/${platform}/list.json" | jq "${jq_cmd}" )
  binary=${binary//'+'/'%2B'}
  binary=${binary%\"}
  binary=${binary#\"}

  mkdir ${solc_folder}/${version}
  curl -q -s -o ${solc_folder}/${version}/solc "https://binaries.soliditylang.org/${platform}/${binary}"
  chmod ugo+x ${solc_folder}/${version}/solc
}

# get solc binaries
for f in ${solc_versions[@]}; do
  if ! [ -f ${solc_folder}/${f}/solc ]; then
    downloadSolc ${f}
  fi
  # check solc
  "${solc_folder}/${f}/solc" --version | grep "Version: ${f}+"
done

# solc ./base_content_space.sol --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize | abigen --pkg=contracts --out build/base_content_space.go --combined-json -
abigen(){
    # 1: solc version
    # 2: contract file (.sol)
    # 3: package name
    # 4: output file - go bindings
    local version=$1
    solc_bin="${solc_folder}/${version}/solc"

    "${solc_bin}" "$2" --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize | "$abigen_dir/abigen" --pkg=${3} --out "${4}" --combined-json -
    ###./abigen --sol "${sol_dir}/${2}" --pkg=${3} --out "${out_dir}/${5}"
    ret=$?
    if [[ $ret -ne 0 ]]; then
        echo "FAILED  : error creating go binding for ${2}!"
        exit 1
    else
        echo "SUCCESS : go binding for ${2} generated at ${4}"
    fi
}

# checkout 'known' versions of contracts in folder _build_contracts_go
# and produce go-bindings in ../../contracts-go
mkdir -p _build_contracts_go
(
  cd _build_contracts_go

  if [ ! -d contracts ]; then
    echo "#### cloning contracts repo"
    git clone https://github.com/eluv-io/contracts.git
  fi
  cd contracts/
  git checkout develop
  git pull origin

  # recent contracts require solc 0.5.4
  (
    # shellcheck disable=SC2030
    echo "#### generating latest contracts"
    mkdir -p ../../contracts-go/contracts
    mkdir -p ../../contracts-go/tradable
    abigen $solc_0_5_4 base_content_space.sol contracts ../../contracts-go/contracts/base_content_space.go
    echo "#### generating latest contracts/tradable"
    abigen $solc_0_5_4 tradable/elv_tradable_full.sol tradable ../../contracts-go/tradable/elv_tradable_full.go
  )

  # older contracts require solc 0.4.24
  (
    git checkout "53ee85a858257d32c9650ff76f50e5ee2a1379d9"
    echo "#### generating contracts_20190331"
    mkdir -p ../../contracts-go/contracts_20190331
    abigen $solc_0_4_24 base_content_space.sol contracts_20190331 ../../contracts-go/contracts_20190331/base_content_space.go

    git checkout "5ed6b64537175315be6c79fdbfb9a3d35086344d"
    echo "#### generating contracts_20200206"
    mkdir -p ../../contracts-go/contracts_20200206
    abigen $solc_0_4_24 base_content_space.sol contracts_20200206 ../../contracts-go/contracts_20200206/base_content_space.go

    # 133cb66 is the revision used in the fabric but does not contain all events
    # git checkout "133cb66c6e67fc946c6b73d7f7db3bbb913d4859" # 2020-08-06
    git checkout "2536dfb9b29e394feedcc1c5c7d16f67d21a35bb"   # 2020-08-11
    echo "#### generating contracts_20200803"
    mkdir -p ../../contracts-go/contracts_20200803
    abigen $solc_0_4_24 base_content_space.sol contracts_20200803 ../../contracts-go/contracts_20200803/base_content_space.go
  )
)

# all good - now copy bindings to their 'legacy' locations
cd ${curr_dir}
rm -rf _build_contracts_go
mkdir -p build
mkdir -p build/20190331
mkdir -p build/20200206
mkdir -p build/20200803
mkdir -p build/tradable

cp contracts-go/contracts/base_content_space.go build/base_content_space.go
cp contracts-go/contracts_20190331/base_content_space.go build/20190331/base_content_space.go
cp contracts-go/contracts_20200206/base_content_space.go build/20200206/base_content_space.go
cp contracts-go/contracts_20200803/base_content_space.go build/20200803/base_content_space.go
cp contracts-go/tradable/elv_tradable_full.go build/tradable/elv_tradable_full.go
