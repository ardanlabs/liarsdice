// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

contract Bank {

    // Owner represents the address who deployed the contract.
    address public Owner;

    // playerBalance represents the amount of money a player have available.
    mapping (address => uint256) private playerBalance;

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

    // Reconcile settles the accounting for a game that was played.
    function Reconcile(address winningPlayer, address[] calldata losers, uint256 ante, uint256 gameFee) onlyOwner public {

        // Build the pot for the winner based on taking the ante from
        // each player that lost.
        uint256 pot;
        for (uint i = 0; i < losers.length; i++) {
            if (playerBalance[losers[i]] < ante) {
                emit EventLog(string.concat("player balance ", Error.Itoa(playerBalance[losers[i]]), " is less than ante ", Error.Itoa(ante)));
                pot += playerBalance[losers[i]];
                playerBalance[losers[i]] = 0;                
            } else {
                pot += ante;
                playerBalance[losers[i]] -= ante;
            }
        }

        emit EventLog(string.concat("ante[", Error.Itoa(ante), "] gameFee[", Error.Itoa(gameFee), "] pot[", Error.Itoa(pot), "]"));

        // This should not happen but check to see if the pot is 0 because none
        // of the losers had a player balance.
        if (pot == 0) {
            revert("no pot was created based on player balances");
        }

        // This should not happen but check there is enough in the pot to cover
        // the game fee.
        if (pot < gameFee) {
            playerBalance[Owner] += pot;
            emit EventLog(string.concat("winningPlayer[0] owner[", Error.Itoa(pot), "]"));
            return;
        }

        // Take the game fee from the pot and give the winner the remaining pot
        // and the owner the game fee.
        pot -= gameFee;
        playerBalance[winningPlayer] += pot;
        playerBalance[Owner] += gameFee;
        emit EventLog(string.concat("winningPlayer[", Error.Itoa(pot), "] owner[", Error.Itoa(gameFee), "]"));
    }

    // PlayerBalance returns the current players balance.
    function PlayerBalance(address player) onlyOwner view public returns (uint) {
        return playerBalance[player];
    }

    // =========================================================================
    // Player Wallet Calls

    // Deposit the given amount to the player balance.
    function Deposit() payable public {
        playerBalance[msg.sender] += msg.value;
        emit EventLog(string.concat("deposit: ", Error.Addrtoa(msg.sender), " - ", Error.Itoa(playerBalance[msg.sender])));
    }

    // Withdraw the given amount from the player balance.
    function Withdraw() payable public {
        address payable player = payable(msg.sender);

        if (playerBalance[msg.sender] == 0) {
            revert("not enough balance");
        }

        player.transfer(playerBalance[msg.sender]);        
        playerBalance[msg.sender] = 0;

        emit EventLog(string.concat("withdraw: ", Error.Addrtoa(msg.sender)));
    }
}