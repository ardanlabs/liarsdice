// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

contract LiarsDice {

    // game holds all game related data.
    struct game {
        uint created_at;
        bool finished;
        uint256 pot;
    }

    // Owner represents the address who deployed the contract.
    address public Owner;

    // playerbalance represents the amount of coins a player have available.
    mapping (address => uint) public playerbalance;

    // Games represents a list of all games.
    mapping (string => game) public games;

    // EventLog provides support for external logging.
    event EventLog(string value);

    // EventPlaceAnte is an event to indicate a bet was performed.
    event EventPlaceAnte(address player, string uuid, uint amount);

    // EventNewGame is an event to indicate a new game was created.
    event EventNewGame(string uuid);

    // constructor is called when the contract is deployed.
    constructor() {
        Owner = msg.sender;
    }

    // NewGame creates a new game with the given uuid and default values.
    function NewGame(string memory uuid) public {
        games[uuid] = game(block.timestamp, false, 0);
        emit EventNewGame(uuid);
    }

    // PlaceAnte adds the amount to the game pot and removes from player balance.
    function PlaceAnte(string memory uuid, uint256 amount, uint256 minimum) public {
        address player = msg.sender;

        // Check if game is finshed.
        if (games[uuid].finished) {
            revert(string.concat("game ", uuid, " is not available anymore"));
        }

        // Check if player has enough balance.
        if (playerbalance[player] < minimum) {
            revert("not enough balance to place a bet");
        }

        playerbalance[player] -= amount;
        games[uuid].pot += amount;

        emit EventLog(string.concat("player ", Error.Addrtoa(player), " placed a bet of ", Error.Itoa(amount), " LDC on game ", uuid));
        emit EventLog(string.concat("current game pot ", Error.Itoa(games[uuid].pot)));
        emit EventPlaceAnte(player, uuid, amount);
    }

    // GameEnd transfers the game pot amount to the player and resets the pot.
    function GameEnd(address player, string memory uuid) public {
        playerbalance[player] += games[uuid].pot;
        games[uuid].finished = true;
        
        emit EventLog(string.concat("game ", uuid, " is over with a pot of ", Error.Itoa(games[uuid].pot), " LDC. The winner is ", Error.Addrtoa(player)));
    }

    //===========================================================================

    // [For testing purposes]
    // deposits the given amount to the player balance.
    function deposit(address player, uint256 amount) public {
        playerbalance[player] += amount;
    }

    // [For testing purposes]
    // withdraws the given amount to the player balance.
    function withdraw(address player, uint256 amount) public {
        playerbalance[player] -= amount;
    }
}