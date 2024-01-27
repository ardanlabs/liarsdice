-- Version: 1.01
-- Description: Create initial game tables
CREATE TABLE games
(
    game_id VARCHAR    NOT NULL,
    name    VARCHAR    NOT NULL,

    PRIMARY KEY (game_id)
);

CREATE TABLE game_status
(
    game_id         VARCHAR NOT NULL,
    iteration       INT     NOT NULL,
    status          VARCHAR NOT NULL,
    player_last_out VARCHAR NULL,
    player_last_win VARCHAR NULL,
    player_turn     VARCHAR NOT NULL,
    round           INT     NOT NULL,

    PRIMARY KEY (game_id, iteration)
    FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
)

CREATE TABLE game_cups
(
    game_id    VARCHAR NOT NULL,
    iteration  INT     NOT NULL,
    player     VARCHAR NOT NULL,
    order_idx  INT     NOT NULL,
	outs       INT     NOT NULL,

    PRIMARY KEY (game_id, iteration, player)
    FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
)

CREATE TABLE game_dice
(
    game_id   VARCHAR NOT NULL,
    iteration INT     NOT NULL,
    player    VARCHAR NOT NULL,
    dice      INT     NOT NULL,

    PRIMARY KEY (game_id, iteration, player)
    FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
)

CREATE TABLE game_existing_players
(
    game_id   VARCHAR NOT NULL,
    iteration INT     NOT NULL,
    player    VARCHAR NOT NULL,

    PRIMARY KEY (game_id, iteration, player)
    FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
)

CREATE TABLE game_bets
(
    game_id   VARCHAR NOT NULL,
    iteration INT     NOT NULL,
    player    VARCHAR NOT NULL,
    number    INT     NOT NULL,
    suit      INT     NOT NULL,

    PRIMARY KEY (game_id, iteration, player)
    FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
)

CREATE TABLE game_balances
(
    game_id   VARCHAR NOT NULL,
    iteration INT     NOT NULL,
    player    VARCHAR NOT NULL,
    balance   VARCHAR NOT NULL,

    PRIMARY KEY (game_id, iteration, player)
    FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
)
