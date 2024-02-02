// Package gamedb contains game related CRUD functionality.
package gamedb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ardanlabs/liarsdice/business/core/game"
	"github.com/ardanlabs/liarsdice/business/data/sqldb"
	"github.com/ardanlabs/liarsdice/foundation/logger"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Store manages the set of APIs for user database access.
type Store struct {
	log *logger.Logger
	db  sqlx.ExtContext
}

// NewStore constructs the api for data access.
func NewStore(log *logger.Logger, db *sqlx.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

// Create adds a game to the db.
func (s *Store) Create(ctx context.Context, g *game.Game) error {

	// NEED A TRANSACTION HERE

	d := struct {
		ID          uuid.UUID `db:"game_id"`
		Name        string    `db:"name"`
		DateCreated time.Time `db:"date_created"`
	}{
		ID:          g.ID(),
		Name:        g.ID().String(),
		DateCreated: g.DateCreated(),
	}

	q := `
    INSERT INTO games
        (game_id, name, date_created)
    VALUES
        (:game_id, :name, :date_created)`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, d); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	q = `
    INSERT INTO game_state
        (game_id, round, status, player_last_out, player_last_win, player_turn, existing_players)
    VALUES
		(:game_id, :round, :status, :player_last_out, :player_last_win, :player_turn, :existing_players)`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, toDBState(g.State())); err != nil {
		return fmt.Errorf("namedexeccontext-state: %w", err)
	}

	return nil
}

// InsertRound adds a new set of state records for a given round to the db.
func (s *Store) InsertRound(ctx context.Context, state game.State) error {
	dbState := toDBState(state)

	// NEED A TRANSACTION HERE

	q := `
    INSERT INTO game_state
        (game_id, round, status, player_last_out, player_last_win, player_turn, existing_players)
    VALUES
		(:game_id, :round, :status, :player_last_out, :player_last_win, :player_turn, :existing_players)`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, dbState); err != nil {
		return fmt.Errorf("namedexeccontext-state: %w", err)
	}

	q = `
    INSERT INTO game_cups
        (game_id, round, player, order_idx, outs, dice)
    VALUES
		(:game_id, :round, :player, :order_idx, :outs, :dice)`

	for _, dbCup := range dbState.Cups {
		if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, dbCup); err != nil {
			return fmt.Errorf("namedexeccontext-cups: %w", err)
		}
	}

	q = `
    INSERT INTO game_bets
        (game_id, round, bet_order, player, number, suit)
    VALUES
		(:game_id, :round, :bet_order, :player, :number, :suit)`

	for _, dbBet := range dbState.Bets {
		if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, dbBet); err != nil {
			return fmt.Errorf("namedexeccontext-bets: %w", err)
		}
	}

	q = `
    INSERT INTO game_balances
        (game_id, round, player, amount)
    VALUES
		(:game_id, :round, :player, :amount)`

	for _, dbBalance := range dbState.Balances {
		if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, dbBalance); err != nil {
			return fmt.Errorf("namedexeccontext-balances: %w", err)
		}
	}

	return nil
}

// QueryStateByID gets the specified game state from the database.
func (s *Store) QueryStateByID(ctx context.Context, gameID uuid.UUID, round int) (game.State, error) {
	data := struct {
		ID    string `db:"game_id"`
		Round int    `db:"round"`
	}{
		ID:    gameID.String(),
		Round: round,
	}

	const q = `
	SELECT
        g.game_id,
		g.name,
    	g.date_created,
		s.round,
		s.status,
		s.player_last_out,
		s.player_last_win,
		s.player_turn,
		s.existing_players
	FROM
		games AS g
	JOIN
		game_state AS s ON g.game_id = s.game_id
	WHERE
		g.game_id = :game_id AND
		s.round = :round`

	var dbState dbState
	if err := sqldb.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbState); err != nil {
		if errors.Is(err, sqldb.ErrDBNotFound) {
			return game.State{}, fmt.Errorf("namedquerystruct: %w", game.ErrNotFound)
		}
		return game.State{}, fmt.Errorf("namedquerystruct: %w", err)
	}

	return toCoreState(dbState)
}
