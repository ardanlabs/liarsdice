// Package board handles the game board and all interactions.
package board

import (
	"fmt"

	"github.com/ardanlabs/liarsdice/app/cli/liars/engine"
	"github.com/gdamore/tcell/v2"
)

// Game positioning values for static content.
const (
	boardWidth    = 63
	boardHeight   = 19
	messageHeight = 12
	anteX         = 40
	anteY         = 8
	potX          = 40
	potY          = 9
	helpX         = 65
	statusY       = 14
)

// Game positioning values for bet input.
const (
	betRowX = 13
	betRowY = 10
)

// Game positioning values for player rows.
const (
	columnHeight = 2
	playersX     = 3
	outX         = 22
	betX         = 35
	balX         = 50
	myDiceX      = 13
	myDiceY      = 8
)

// Unicode characters for the game board.
const (
	hozTopRune = '\u2580'
	hozBotRune = '\u2584'
	verRune    = '\u2588'
)

var words = []string{"", "one's", "two's", "three's", "four's", "five's", "six's"}

// =============================================================================

// Board represents the game board and all its state.
type Board struct {
	accountID  string
	engine     *engine.Engine
	config     engine.Config
	screen     tcell.Screen
	style      tcell.Style
	bets       []rune
	messages   []string
	lastStatus engine.Status
	modalUp    bool
	modalMsg   string
}

// New contructs a game board and renders the board.
func New(engine *engine.Engine, accountID string) (*Board, error) {
	config, err := engine.Configuration()
	if err != nil {
		return nil, fmt.Errorf("get game configuration: %w", err)
	}

	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)

	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("new screen: %w", err)
	}

	if err := screen.Init(); err != nil {
		return nil, fmt.Errorf("init: %w", err)
	}

	style := tcell.StyleDefault
	style = style.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

	board := Board{
		accountID: accountID,
		config:    config,
		engine:    engine,
		screen:    screen,
		style:     style,
		messages:  make([]string, 5),
	}

	if err := board.printBoard(); err != nil {
		return nil, fmt.Errorf("init: %w", err)
	}

	return &board, nil
}

// Shutdown tearsdown the game board.
func (b *Board) Shutdown() {
	b.screen.Fini()
}

// Run starts a goroutine to handle terminal events. This is a
// blocking call.
func (b *Board) Run() chan struct{} {
	return b.pollEvents()
}

// Events handles any events from the websocket. This function should be
// registered with any code receiving the web socket events.
func (b *Board) Events(event string, address string) {
	b.webEvents(event, address)
}
