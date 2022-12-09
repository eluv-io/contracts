#!/bin/bash
set -Eeuo pipefail

# This script makes sure the required versions of solc are available.
# * unless a solc binary is found in folder '_bin/${platform}/solc/${version},
#   the script downloads the solc binary from 'binaries.soliditylang.org'
# * to add a new version: declare a new variable with the value of the version
#   and add it to the 'solc_versions' variable.

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
