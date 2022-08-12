#!/bin/bash
set -Eeuo pipefail

if [ -z "${2-}" ]; then
  echo "Usage: $0 SOLC_0_4_24_DIR SOLC_0_5_4_DIR"
  exit 1
fi

solc_0_4_24_dir=$(realpath "$1")
solc_0_5_4_dir=$(realpath "$2")

# check solc
"$solc_0_4_24_dir/solc" --version | grep "Version: 0.4.24+"
"$solc_0_5_4_dir/solc" --version | grep "Version: 0.5.4+"

# build abigen
(
  echo "#### building abigen command"
  cd cmds
  go build -o ../abigen ./abigen
)

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
    hash -d solc 2>/dev/null || : # clear cache for solc location
    # shellcheck disable=SC2030
    PATH="$solc_0_5_4_dir:$PATH"
    export PATH

    echo "#### generating latest contracts"
    ../../abigen --sol base_content_space.sol --pkg contracts --out ../../contracts-go/contracts/base_content_space.go
    echo "#### generating latest contracts/tradable"
    ../../abigen --sol tradable/elv_tradable_full.sol --pkg tradable --out ../../contracts-go/tradable/elv_tradable_full.go
  )

  # older contracts require solc 0.4.24
  (
    hash -d solc 2>/dev/null || : # clear cache for solc location
    PATH="$solc_0_4_24_dir:$PATH"
    export PATH

    git checkout "53ee85a858257d32c9650ff76f50e5ee2a1379d9"
    echo "#### generating contracts_20190331"
    ../../abigen --sol base_content_space.sol --pkg contracts_20190331 --out ../../contracts-go/contracts_20190331/base_content_space.go

    git checkout "5ed6b64537175315be6c79fdbfb9a3d35086344d"
    echo "#### generating contracts_20200206"
    ../../abigen --sol base_content_space.sol --pkg contracts_20200206 --out ../../contracts-go/contracts_20200206/base_content_space.go

    git checkout "133cb66c6e67fc946c6b73d7f7db3bbb913d4859"
    echo "#### generating contracts_20200803"
    ../../abigen --sol base_content_space.sol --pkg contracts_20200803 --out ../../contracts-go/contracts_20200803/base_content_space.go
  )
)
