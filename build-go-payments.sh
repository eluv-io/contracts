#!/bin/bash
set -Eeuo pipefail

if [ -z "${1-}" ]; then
  echo "Usage: $0 SOLC_0_8_13_DIR"
  exit 1
fi

solc_0_8_13_dir=$(realpath "$1")

# check solc
"$solc_0_8_13_dir/solc" --version | grep "Version: 0.8.13+"

# build abigen
(
  echo "#### building abigen command"
  cd cmds
  go build -o ../abigen ./abigen
  cd ..
)


(
  base_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
  echo "BASE_DIR $base_dir"
  # recent contracts require solc 0.8.13
  (
    hash -d solc 2>/dev/null || : # clear cache for solc location
    # shellcheck disable=SC2030
    PATH="$solc_0_8_13_dir:$PATH"
    export PATH

    payments_out_dir="${base_dir}/dist/payments"
    payments_sol_file="${base_dir}/src/payments/ERC20Payments.sol"
    payments_abigen_dir="${base_dir}/contracts-go/payments"

    elv_token_out_dir="${base_dir}/dist/tradable"
    elv_token_sol_file="${base_dir}/src/tradable/ElvToken.sol"
    elv_token_abigen_dir="${base_dir}/contracts-go/tradable"

    solc  @openzeppelin/=lib/openzeppelin-contracts/ ${payments_sol_file}  --abi --hashes --optimize -o ${payments_out_dir} --overwrite
    solc  @openzeppelin/=lib/openzeppelin-contracts/ ${payments_sol_file}  --bin --optimize -o ${payments_out_dir} --overwrite

    solc  @openzeppelin/=lib/openzeppelin-contracts/ ${elv_token_sol_file}  --abi --hashes --optimize -o ${elv_token_out_dir} --overwrite
    solc  @openzeppelin/=lib/openzeppelin-contracts/ ${elv_token_sol_file}  --bin --optimize -o ${elv_token_out_dir} --overwrite

    echo "#### generating latest src/payments/ERC20Payments.sol"
    ./abigen --abi "${payments_out_dir}/ERC20Payments.abi" --bin "${payments_out_dir}/ERC20Payments.bin" --pkg payments --out "${payments_abigen_dir}/erc20_payments.go"

    echo "#### generating latest src/tradable/ElvToken.sol"
    ./abigen --abi "${elv_token_out_dir}/ElvToken.abi" --bin "${elv_token_out_dir}/ElvToken.bin" --pkg tradable --out "${elv_token_abigen_dir}/ElvToken.go"

  )
)
