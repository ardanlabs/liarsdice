// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

contract LiarsDice {

    // game holds all game related data.
    struct game {
        uint created_at;
        bool finished;
        uint256 pot;
        uint256 ante;
    }

    // Owner represents the address who deployed the contract.
    address public Owner;

    // Game represents the current game.
    game public Game;

    // playerbalance represents the amount of coins a player have available.
    mapping (address => uint) private playerbalance;

    // EventLog provides support for external logging.
    event EventLog(string value);

    // EventPlaceAnte is an event to indicate a bet was performed.
    event EventPlaceAnte(address player, string uuid, uint amount);

    // EventNewGame is an event to indicate a new game was created.
    event EventNewGame(string uuid);

    // onlyOwner can be used to restrict access to a function for only the owner.
    modifier onlyOwner {
        if (msg.sender != Owner) revert();
        _;
    }

    // constructor is called when the contract is deployed.
    constructor() {
        Owner = msg.sender;
    }

    // NewGame creates a new game.
    function NewGame() public {
        Game = game(block.timestamp, false, 0, 5);
    }

    // PlaceAnte adds the amount to the game pot and removes from player balance.
    function PlaceAnte() public {

        // Check if game is finshed.
        if (Game.finished) {
            revert(string.concat("game is not available anymore"));
        }

        // Check if player has enough balance to pay the game ante.
        if (playerbalance[msg.sender] < Game.ante) {
            revert(string.concat("not enough balance to join the game, it requires ", Error.Itoa(Game.ante)));
        }

        // Remove game ante from player's balance.
        playerbalance[msg.sender] -= Game.ante;

        // Increase game pot.
        Game.pot += Game.ante;

        emit EventLog(string.concat("player ", Error.Addrtoa(msg.sender), " joined the game"));
        emit EventLog(string.concat("current game pot ", Error.Itoa(Game.pot)));
    }

    // GameEnd transfers the game pot amount to the player and finish the game.
    function GameEnd(address player) onlyOwner public {

        // Finish the game so it does not accept any more players.
        Game.finished = true;

        // Move the pot amount to the player's balance.
        playerbalance[player] += Game.pot;

        emit EventLog(string.concat("game is over with a pot of ", Error.Itoa(Game.pot), " LDC. The winner is ", Error.Addrtoa(player)));
    }

    // GameAnte returns the game pot amount.
    function GameAnte() onlyOwner public returns (uint) {
        emit EventLog(string.concat("game current pot: ", Error.Itoa(Game.pot)));
        return Game.pot;
    }

    // Deposit the given amount to the player balance.
    function Deposit() payable public {
        playerbalance[msg.sender] += msg.value;
        emit EventLog(string.concat("deposit: ", Error.Addrtoa(msg.sender), " - ", Error.Itoa(msg.value)));
    }

    // Withdraw the given amount from the player balance.
    // TODO: we still need to find a way to transfer the balance amount to the
    // player's wallet.
    function Withdraw(uint256 amount) public {

        // Check if player has enough balance.
        if (playerbalance[msg.sender] < amount) {
            revert("not enough balance");
        }

        playerbalance[msg.sender] -= amount;
        emit EventLog(string.concat("withdraw: ", Error.Addrtoa(msg.sender), " - ", Error.Itoa(amount)));
    }
}