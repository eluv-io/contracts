const fs = require("fs");
const Path = require("path");
const solc = require("solc").setupMethods(require("./soljson-v0.4.21+commit.dfe3193c"));
const ethers = require("ethers");
const glob = require("glob");

// Create minimal JS containing ABI and bytecode from compiled contract
const BuildContracts = () => {
  let sources = {};
  glob.sync(Path.join(__dirname, "..", "src", "**", "*.sol"))
    .forEach(path => {
      if(!path.endsWith("Migrations.sol")) {
        sources[Path.basename(path)] = fs.readFileSync(path).toString()
      }
    });

  const compilationResult = solc.compile({sources}, 1);

  let topicMap = {};
  Object.keys(compilationResult.contracts).map((contractKey => {
    const contractName = contractKey.split(":")[1];
    const abi = JSON.parse(compilationResult.contracts[contractKey].interface);
    const contractData = {
      abi,
      bytecode: compilationResult.contracts[contractKey].bytecode
    };

    // Write out contract info
    const contractJS = "const contract=" + JSON.stringify(contractData) + "; module.exports=contract;";
    fs.writeFileSync(Path.join(__dirname, "..", "src.js", contractName + ".js"), contractJS);

    // Map event signature hash (aka "topic") to its interface
    const events = abi.filter(entry => entry.type === "event");
    events.forEach(event => {
      const signature = `${event.name}(${event.inputs.map(input => input.type).join(",")})`;
      const signatureHash = ethers.utils.keccak256(ethers.utils.toUtf8Bytes(signature));
      if(!topicMap[signatureHash] || contractName.startsWith("Base")) {
        // Base contracts take priority of there are methods with identical signatures
        topicMap[signatureHash] = {abi: [event], contract: contractName};
      }
    });
  }));

  const topicMapJS = "const topics=" + JSON.stringify(topicMap) + "; module.exports=topics;";
  fs.writeFileSync(Path.join(__dirname, "..", "src.js", "topics", "Topics" + ".js"), topicMapJS);
};

BuildContracts();
