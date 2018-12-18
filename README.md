#In order to link with the libraries, solc should be used as follow:

solc da_file_to_compile.sol --bin --libraries libraries --optimize

Note: 
	"--optimize" cuts the size of the binary by 40%, which is important otherwise we are above limit for our main contract

	"--libraries libraries" indicates that the library linking to use are contained in the libraries file. 
	Each line in that file, contains a linking and provides the address at which that library was deployed

### Deploy base_content_space.sol using geth node

*  generate binary and abi file for the contract to be deployed
````
solc --abi base_content_space.sol --optimize > build/base_content_space.abi
solc --bin base_content_space.sol --optimize > build/base_content_space.bin
````
* to get content space binary and abi part
````
more build/base_content_space.bin
more build/base_content_space.abi
````
###### Note : copy only the content space contract's abi and bin.

##### In geth node,
* Initialize the variables
````
var contentSpace = eth.contract(<CONTENTS_OF_ABI_FILES>)
var bytecode = '0x<CONTENTS_OF_BIN_FILE>'
````
* Deploy the contract
````
var deploy = {from:eth.coinbase, data:bytecode, gas:80000000}

personal.unlockAccount(eth.coinbase)

var contentSpaceInstance = contentSpace.new(deploy)
````
* to get the address of the deployed content space
````
contentSpaceInstance.address
````

the address can then be converted to ispc_XXX fromat
