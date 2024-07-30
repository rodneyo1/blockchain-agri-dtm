// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract TransactionContract {
    struct Transaction {
        bytes32 previousHash;
        uint256 timestamp;
        uint256 nonce;
        string farmerName;
        string farmerPhoneNumber;
        string customerName;
        uint256 amount;
        string customerPhoneNumber;
    }

    Transaction[] public transactions;
    uint256 public transactionCount;

    event TransactionLogged(
        bytes32 previousHash,
        uint256 timestamp,
        uint256 nonce,
        string farmerName,
        string farmerPhoneNumber,
        string customerName,
        uint256 amount,
        string customerPhoneNumber
    );

    function logTransaction(
        bytes32 _previousHash,
        uint256 _timestamp,
        uint256 _nonce,
        string memory _farmerName,
        string memory _farmerPhoneNumber,
        string memory _customerName,
        uint256 _amount,
        string memory _customerPhoneNumber
    ) public {
        Transaction memory newTransaction = Transaction({
            previousHash: _previousHash,
            timestamp: _timestamp,
            nonce: _nonce,
            farmerName: _farmerName,
            farmerPhoneNumber: _farmerPhoneNumber,
            customerName: _customerName,
            amount: _amount,
            customerPhoneNumber: _customerPhoneNumber
        });

        transactions.push(newTransaction);
        transactionCount++;

        emit TransactionLogged(
            _previousHash,
            _timestamp,
            _nonce,
            _farmerName,
            _farmerPhoneNumber,
            _customerName,
            _amount,
            _customerPhoneNumber
        );
    }

    function getTransaction(uint256 index) public view returns (
        bytes32 previousHash,
        uint256 timestamp,
        uint256 nonce,
        string memory farmerName,
        string memory farmerPhoneNumber,
        string memory customerName,
        uint256 amount,
        string memory customerPhoneNumber
    ) {
        Transaction memory t = transactions[index];
        return (
            t.previousHash,
            t.timestamp,
            t.nonce,
            t.farmerName,
            t.farmerPhoneNumber,
            t.customerName,
            t.amount,
            t.customerPhoneNumber
        );
    }

    function getTransactionCount() public view returns (uint256) {
        return transactionCount;
    }
}
