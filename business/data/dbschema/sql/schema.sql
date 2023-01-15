-- Version: 1.01
-- Description: Create table games
CREATE TABLE games
(
    id               VARCHAR PRIMARY KEY,
    status           VARCHAR,
    current_cup      INT,
    round            INT,
    ante_usd         MONEY,
    org_order        VARCHAR[],
    balance_gwei     FLOAT8,
    last_out_acct_id VARCHAR,
    last_win_acct_id VARCHAR,
    current_acct_id  VARCHAR,
    cups_order       VARCHAR[],
    balance          VARCHAR[]
);

-- Version: 1.02
-- Description: Create table bets
CREATE TABLE bets
(
    id         BIGSERIAL UNIQUE NOT NULL,
    game_id    VARCHAR,
    account_id VARCHAR,
    number     INT,
    suite      INT,
    PRIMARY KEY (game_id, account_id)
);

-- Version: 1.03
-- Description: Create table cups
CREATE TABLE cups
(
    game_id    VARCHAR REFERENCES games (id),
    account_id VARCHAR,
    order_idx  INT,
    last_bet   BIGINT REFERENCES bets (id),
    outs       INT,
    dice       INT[],
    PRIMARY KEY (game_id, account_id)
);