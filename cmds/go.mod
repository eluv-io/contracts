module github.com/eluv-io/contracts/cmds

go 1.13

require (
	// contracts @latest -> used by ethverify abi-diff cmd
	github.com/eluv-io/contracts v0.0.0-20201105023949-91a07d95aad7
	github.com/ethereum/go-ethereum v1.9.25
	github.com/mattn/go-colorable v0.1.2
	github.com/mattn/go-isatty v0.0.9
	github.com/mitchellh/go-homedir v1.1.0
	github.com/olekukonko/tablewriter v0.0.2-0.20190409134802-7e037d187b0c
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.4.0
	gopkg.in/src-d/go-git.v4 v4.13.1
)
