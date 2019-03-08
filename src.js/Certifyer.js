const contract={"abi":[{"constant":true,"inputs":[{"name":"message_hash","type":"bytes32"},{"name":"sig","type":"bytes"}],"name":"recoverSignerFromMessageHash","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[{"name":"sig","type":"bytes"}],"name":"splitSignature","outputs":[{"name":"","type":"uint8"},{"name":"","type":"bytes32"},{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[{"name":"message","type":"bytes"},{"name":"sig","type":"bytes"}],"name":"recoverSignerFromMessage","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[{"name":"message","type":"bytes"}],"name":"messageHash","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[{"name":"i","type":"uint256"}],"name":"uint2str","outputs":[{"name":"","type":"bytes"}],"payable":false,"stateMutability":"pure","type":"function"}],"bytecode":"6105ef610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146060604052600436106100835763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416639c67eb248114610088578063a7bb5803146100fc578063aef840631461016a578063ed54c3b6146101f2578063f76f950e1461024a575b600080fd5b6100d3600480359060446024803590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506102cc95505050505050565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b61014260046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061034d95505050505050565b60405160ff909316835260208301919091526040808301919091526060909101905180910390f35b6100d360046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061038895505050505050565b61023860046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506103a395505050505050565b60405190815260200160405180910390f35b610255600435610485565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610291578082015183820152602001610279565b50505050905090810190601f1680156102be5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6000806000806102db8561034d565b919450925090506001868484846040516000815260200160405260405193845260ff9092166020808501919091526040808501929092526060840192909252608090920191516020810390808403906000865af1151561033a57600080fd5b5050602060405103519695505050505050565b600080600080600080865160411461036457600080fd5b6020870151925060408701519150606087015160001a979296509094509092505050565b600061039c610396846103a3565b836102cc565b9392505050565b60006103af8251610485565b826040517f19457468657265756d205369676e6564204d6573736167653a0a0000000000008152601a810183805190602001908083835b602083106104055780518252601f1990920191602091820191016103e6565b6001836020036101000a038019825116818451161790925250505091909101905082805190602001908083835b602083106104515780518252601f199092019160209182019101610432565b6001836020036101000a03801982511681845116179092525050509190910193506040925050505180910390209050919050565b61048d6105b1565b6000806104986105b1565b60008515156104dc5760408051908101604052600181527f3000000000000000000000000000000000000000000000000000000000000000602082015294506105a8565b859350600092505b83156104fb57600190920191600a840493506104e4565b826040518059106105095750595b818152601f19601f8301168101602001604052905091505060001982015b85156105a4576000198101907f01000000000000000000000000000000000000000000000000000000000000006030600a890601029083908151811061056957fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a86049550610527565b8194505b50505050919050565b602060405190810160405260008152905600a165627a7a72305820cee2f979aab619afb7bcebe95d811d930e10bf938d2e3f4a7dd6c92725c99fd60029"}; module.exports=contract;