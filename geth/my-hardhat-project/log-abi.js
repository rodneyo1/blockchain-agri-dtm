const fs = require("fs");
const path = require("path");

async function main() {
    // Define the path to the compiled contract artifacts
    const contractName = "TransactionContract"; // Replace with your contract name
    const artifactsPath = path.join(__dirname, "artifacts", "contracts", `${contractName}.sol`, `${contractName}.json`);
    
    // Read the compiled contract artifacts
    const artifacts = JSON.parse(fs.readFileSync(artifactsPath, "utf8"));

    // Extract the ABI
    const abi = artifacts.abi;

    // Log the ABI to the console
    console.log("Contract ABI:", JSON.stringify(abi, null, 2));
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});

