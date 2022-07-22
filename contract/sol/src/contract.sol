// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

contract LiarsDice {

    // gameInfo holds all game related data.
    struct gameInfo {
        bool playing;
        uint256 pot;
    }

    // =========================================================================

    // Owner represents the address who deployed the contract.
    address public Owner;

    // game represents the only game we current allow to be played.
    gameInfo private game;

    // fees collected by the smart contract for the owner.
    uint256 private fees;

    // playerBalance represents the amount of money a player have available.
    mapping (address => uint256) private playerBalance;

    // EventLog provides support for external logging.
    event EventLog(string value);

    // =========================================================================

    // onlyOwner can be used to restrict access to a function for only the owner.
    modifier onlyOwner {
        if (msg.sender != Owner) revert();
        _;
    }

    // constructor is called when the contract is deployed.
    constructor() {
        Owner = msg.sender;
        game = gameInfo(false, 0);
    }

    // =========================================================================
    // Owner Only Calls

    // PlaceAnte adds the amount to the game pot and removes from player balance. Each
    // player pays the gas fees for this call.
    function PlaceAnte(address player, uint256 ante, uint256 gasFee) onlyOwner public {
        if (!game.playing) {
            revert("game is not available anymore");
        }

        uint256 totalPrice = ante + gasFee;

        if (playerBalance[player] < totalPrice) {
            revert(string.concat("not enough balance to join the game, it requires ", Error.Itoa(totalPrice)));
        }

        playerBalance[player] -= totalPrice;

        game.pot += ante;
        fees += gasFee;

        emit EventLog(string.concat("player: ", Error.Addrtoa(msg.sender), " joined the game"));
        emit EventLog(string.concat("current game pot: ", Error.Itoa(game.pot)));
    }

    // GameEnd transfers the game pot amount to the winning player and they pay
    // any gas fees.
    function GameEnd(address winningPlayer, uint256 gasFee) onlyOwner public {
        game.playing = false;

        playerBalance[winningPlayer] += game.pot;
        fees += gasFee;

        emit EventLog(string.concat("game is over with a pot of ", Error.Itoa(game.pot), ". The winner is ", Error.Addrtoa(winningPlayer)));

        game.playing = false;
        game.pot = 0;
    }

    // GamePot returns the game pot amount.
    function GamePot() onlyOwner view public returns (uint) {
        return game.pot;
    }

    // PlayerBalance returns the current players balance.
    function PlayerBalance(address player) onlyOwner view public returns (uint) {
        return playerBalance[player];
    }

    // function Players() onlyOwner view public returns (mapping (address => uint256)) {
    //     return playerBalance;
    // }

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