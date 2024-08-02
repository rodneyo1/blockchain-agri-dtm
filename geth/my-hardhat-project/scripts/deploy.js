async function main() {
  const [deployer] = await ethers.getSigners();

  console.log("Deploying contracts with the account:", deployer.address);

  const TransactionContract = await ethers.getContractFactory("TransactionContract");
  const transactionContract = await TransactionContract.deploy();

  console.log("TransactionContract deployed to:", transactionContract.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
