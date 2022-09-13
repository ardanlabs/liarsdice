// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

// abc
contract Bank {

    // Owner represents the address who deployed the contract.
    address public Owner; // 123

    // API represents the address of the contract allowed to interact with this Bank.
    address public API; // def

    // accountBalances represents the amount of money an account has available.
    mapping (address => uint256) private accountBalances;

    // EventLog provides support for external logging.
    event EventLog(string value);

    // =========================================================================

    // constructor is called when the contract is deployed.
    constructor() {
        Owner = msg.sender;
    }

    // =========================================================================
    // Owner Only Calls

    // onlyOwner can be used to restrict access to a function for only the owner.
    modifier onlyOwner {
        if (msg.sender != Owner) revert();
        _;
    }

    // onlyAPI can be used to restrict access to a function for only the API.
    modifier onlyAPI {
        if (msg.sender != API) revert();
        _;
    }

    // UpdateAPI allows the owner to authorize a new API contract.
    function UpdateAPI(address api) onlyOwner public {
        API = api;
    }

    // AccountBalance returns the current account's balance.
    function AccountBalance(address account) onlyOwner view public returns (uint) {
        return accountBalances[account];
    }

    // AllBalances returns an ordered list of balances for the provided accounts.
    function AllBalances(address[] accounts) onlyAPI view public returns (uint[]) {
        uint[] calldata output = new uint[];
        for (uint i = 0; i < accounts.length; i++) {
            output.push(accountBalances[accounts[i]]);
        }
        return output;
    }

    // MultiUpdate will update an account's balance with a provided linked list
    // of update amounts.
    function MultiUpdate(address[] accounts, int[] amounts) onlyAPI public {
        for (uint i = 0; i < accounts.length; i++) {
            accountBalances[accounts[i]] += amounts[i];
        }
    }

    // =========================================================================
    // Account Only Calls

    // Balance returns the balance of the caller.
    function Balance() view public returns (uint) {
        return accountBalances[msg.sender];
    }

    // Deposit the given amount to the account balance.
    function Deposit() payable public {
        accountBalances[msg.sender] += msg.value;
        emit EventLog(string.concat("deposit[", Error.Addrtoa(msg.sender), "] balance[", Error.Itoa(accountBalances[msg.sender]), "]"));
    }

    // Withdraw the given amount from the account balance.
    function Withdraw() payable public {
        address payable account = payable(msg.sender);

        if (accountBalances[msg.sender] == 0) {
            revert("not enough balance");
        }

        uint256 amount = accountBalances[msg.sender];
        account.transfer(amount);        
        accountBalances[msg.sender] = 0;

        emit EventLog(string.concat("withdraw[", Error.Addrtoa(msg.sender), "] amount[", Error.Itoa(amount), "]"));
    }
}

// def
contract BankAPI {

    // Owner represents the address who deployed the contract.
    address public Owner; // 123

    // Bank represents the address of the bank store.
    address public Bank;

    // =========================================================================

    // constructor is called when the contract is deployed.
    constructor(address bank) {
        Owner = msg.sender;
        Bank = bank;
    }

    // =========================================================================
    // Owner Only Calls

    // onlyOwner can be used to restrict access to a function for only the owner.
    modifier onlyOwner {
        if (msg.sender != Owner) revert();
        _;
    }

    // Reconcile settles the accounting for a game that was played.
    function Reconcile(address winner, address[] calldata losers, uint256 anteWei, uint256 gameFeeWei) onlyOwner public {

        // The accounts array will contain all accounts involved in
        // reconciliation, including the Owner and winner.
        //
        // Owner is known to be at index 0
        // winner is known to be at index 1
        //
        // Constants may be preferable to refer to these addresses, but have
        // been omitted here for brevity's sake.
        address[] memory accounts = new address[Owner, winner];
        for (uint i=0; i < losers.length; i++) {
            accounts.push(losers[i]);
        }

        // Retrieve all balances from the Bank for the given set of accounts
        uint[] memory balances = Bank.AllBalances(accounts);

        // Create an updates array to track changes to balances for update later
        int[] memory updates = new int[];


        // Add the ante for each player to the pot. The initialization is
        // for the winner's ante.
        uint256 pot = anteWei;
        for (uint i = 2; i < accounts.length; i++) {
            if (balances[i] < anteWei) {
                emit EventLog(string.concat("account balance ", Error.Itoa(balances[i]), " is less than ante ", Error.Itoa(anteWei)));
                pot += balances[i];
                updates[i] -= balances[i];
            } else {
                pot += anteWei;
                updates[i] -= anteWei;
            }
        }

        emit EventLog(string.concat("ante[", Error.Itoa(anteWei), "] gameFee[", Error.Itoa(gameFeeWei), "] pot[", Error.Itoa(pot), "]"));

        // This should not happen but check to see if the pot is 0 because none
        // of the losers had an account balance.
        if (pot == 0) {
            revert("no pot was created based on account balances");
        }

        // This should not happen but check there is enough in the pot to cover
        // the game fee.
        if (pot < gameFeeWei) {
            updates[0] += pot;
            emit EventLog(string.concat("pot less than fee: winner[0] owner[", Error.Itoa(pot), "]"));
            return;
        }

        // Take the game fee from the pot and give the winner the remaining pot
        // and the owner the game fee.
        pot -= gameFeeWei;
        updates[1] += pot;
        updates[0] += gameFeeWei;
        Bank.MultiUpdate(accounts, updates);
        emit EventLog(string.concat("winner[", Error.Itoa(pot), "] owner[", Error.Itoa(gameFeeWei), "]"));
    }
}
