# Prerequisites

Build solc. The version must be 0.5.4

```shell
# Install solc using brew (instructions from soliditylang.org)
git clone https://github.com/ethereum/homebrew-ethereum.git
cd homebrew-ethereum
# Find commit hash for 0.5.4 at https://github.com/ethereum/homebrew-ethereum/commits/master/solidity.rb
git checkout a2645750428cc17b1d9a3c07d8dba798346a385a
brew unlink solidity
brew install solidity.rb

# OR Install solc using solc-select (Python)
pip3 install solc-select
solc-select install 0.5.4
solc-select use 0.5.4
```

Build abigen
```shell
cd $GOPATH/pkg/mod/github.com/ethereum/go-ethereum@v1.9.11/cmd/abigen
go install .
```

# Build
```shell
# Add abigen to PATH
export PATH=$PATH:$GOPATH/bin

./build.sh
```

In order to link with the libraries, solc should be used as follows:
```
solc my_contract.sol --bin --libraries libraries --optimize
```

Notes:
* "--optimize" cuts the size of the binary by 40%, which is important - otherwise some contracts are over the size limit
* "--libraries libraries" indicates that the library linking to use are contained in the libraries file. Each line in that file contains a link and provides the address at which that library was deployed


## Version History

- tag `2.0.0` - Base contracts corresponding to the Content Fabric 'auth v2' version
- tag `3.0.0` - Base contracts correspondnig to the Content Fabric 'auth v3' version
