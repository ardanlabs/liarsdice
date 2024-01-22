-- Version: 1.01
-- Description: Create table games
CREATE TABLE games
(
    game_id              VARCHAR    NOT NULL,
    status               VARCHAR    NOT NULL,
    round                INT        NOT NULL,
    ante_usd             MONEY      NOT NULL,
    player_turn          INT,
    player_last_out      VARCHAR,
    player_last_win      VARCHAR,

    PRIMARY KEY (game_id)
);

-- Version: 1.02
-- Description: Create table game_players
CREATE TABLE game_players
(
    game_id VARCHAR NOT NULL,
    player  VARCHAR NOT NULL,
    order   INT     NOT NULL,

    PRIMARY KEY (game_id, player)
    FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
)

-- Version: 1.03
-- Description: Create table game_player_bets
CREATE TABLE game_player_bets
(
    game_id    VARCHAR   NOT NULL,
    player     VARCHAR   NOT NULL,
    round      INT       NOT NULL,
    bet_number INT       NOT NULL,
    number     INT       NOT NULL,
    suite      INT       NOT NULL,

    PRIMARY KEY (game_id, player, round, bet_number)
    FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
    FOREIGN KEY (game_players) REFERENCES game_players(game_id, player) ON DELETE CASCADE
);

-- Version: 1.04
-- Description: Create table game_player_cups
CREATE TABLE game_player_cups
(
    game_id    VARCHAR NOT NULL,
    player     VARCHAR NOT NULL,
    round      INT     NOT NULL,
    dice       INT[]   NOT NULL,

    PRIMARY KEY (game_id, player, round),
    FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
    FOREIGN KEY (game_players) REFERENCES game_players(game_id, player) ON DELETE CASCADE
);

-- Version: 1.04
-- Description: Create table game_player_outs
CREATE TABLE game_player_outs
(
    game_id    VARCHAR NOT NULL,
    player     VARCHAR NOT NULL,
    outs       INT     NOT NULL,

    PRIMARY KEY (game_id, player),
    FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
    FOREIGN KEY (game_players) REFERENCES game_players(game_id, player) ON DELETE CASCADE
);