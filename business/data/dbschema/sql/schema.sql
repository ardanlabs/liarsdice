-- Version: 1.01
-- Description: Create table games
CREATE TABLE games
(
    game_id               VARCHAR    NOT NULL,
    status                VARCHAR    NOT NULL,
    round                 INT        NOT NULL,
    ante_usd              MONEY      NOT NULL,
    players               VARCHAR[],
    existing_players      VARCHAR[],
    player_balances_gwei  VARCHAR[],
    player_turn           INT,
    player_last_out       VARCHAR,
    player_last_win       VARCHAR,

    PRIMARY KEY (game_id)
);

-- Version: 1.02
-- Description: Create table bets
CREATE TABLE bets
(
    bet_id     BIGSERIAL NOT NULL,
    game_id    VARCHAR   NOT NULL,
    account_id VARCHAR   NOT NULL,
    number     INT       NOT NULL,
    suite      INT       NOT NULL,

    PRIMARY KEY (bet_id),
    FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
);

-- Version: 1.03
-- Description: Create index on bets
CREATE UNIQUE INDEX bets_idx1 ON bets (game_id, account_id);

-- Version: 1.04
-- Description: Create table cups
CREATE TABLE cups
(
    game_id    VARCHAR NOT NULL,
    account_id VARCHAR NOT NULL,
    order_idx  INT     NOT NULL,
    last_bet   BIGINT  NOT NULL,
    outs       INT     NOT NULL,
    dice       INT[]   NOT NULL,

    PRIMARY KEY (game_id, account_id),
    FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
    FOREIGN KEY (last_bet) REFERENCES bets(bet_id) ON DELETE CASCADE
);