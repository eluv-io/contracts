#In order to link with the libraries, solc should be used as follow:

solc da_file_to_compile.sol --bin --libraries libraries --optimize

Note: 
	"--optimize" cuts the size of the binary by 40%, which is important otherwise we are above limit for our main contract

	"--libraries libraries" indicates that the library linking to use are contained in the libraries file. 
	Each line in that file, contains a linking and provides the address at which that library was deployed
