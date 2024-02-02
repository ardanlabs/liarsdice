-- Version: 1.01
-- Description: Create initial game tables
CREATE TABLE games
(
    game_id      UUID      NOT NULL,
    name         VARCHAR   NOT NULL,
    date_created TIMESTAMP NOT NULL,

    PRIMARY KEY (game_id)
);

CREATE TABLE game_state
(
    game_id          UUID      NOT NULL,
    round            INT       NOT NULL,
    status           VARCHAR   NOT NULL,
    player_last_out  VARCHAR   NULL,
    player_last_win  VARCHAR   NULL,
    player_turn      VARCHAR   NOT NULL,
    existing_players VARCHAR[] NOT NULL,

    PRIMARY KEY (game_id, round),
    FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
);

CREATE TABLE game_cups
(
    game_id    UUID    NOT NULL,
    round      INT     NOT NULL,
    player     VARCHAR NOT NULL,
    order_idx  INT     NOT NULL,
	outs       INT     NOT NULL,
    dice       INT[]   NOT NULL,

    PRIMARY KEY (game_id, round, player),
    FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
);

CREATE TABLE game_bets
(
    game_id   UUID    NOT NULL,
    round     INT     NOT NULL,
    bet_order INT     NOT NULL,
    player    VARCHAR NOT NULL,
    number    INT     NOT NULL,
    suit      INT     NOT NULL,

    PRIMARY KEY (game_id, round, bet_order, player),
    FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
);

CREATE TABLE game_balances
(
    game_id   UUID    NOT NULL,
    round     INT     NOT NULL,
    player    VARCHAR NOT NULL,
    amount    VARCHAR NOT NULL,

    PRIMARY KEY (game_id, round, player),
    FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
);
