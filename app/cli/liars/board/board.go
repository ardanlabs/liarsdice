// Package board handles the game board and all interactions.
package board

import (
	"errors"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

// Game positioning values for static content.
const (
	boardWidth    = 63
	boardHeight   = 18
	messageHeight = 12
	anteX         = 40
	anteY         = 8
	potX          = 40
	potY          = 9
	helpX         = 65
	configY       = 15
)

// Game positioning values for user input.
const (
	depositRowX = 11
	depositRowY = 19
	betRowX     = 13
	betRowY     = 10
)

// Game positioning values for changing values.
const (
	columnHeight = 2
	playersX     = 3
	betX         = 30
	balX         = 50
	myDiceX      = 13
	myDiceY      = 8
)

// =============================================================================

// player represents a player in the game.
type player struct {
	address string
	lastBet string
	balance string
	active  bool
}

// Board represents the game board and all its state.
type Board struct {
	screen   tcell.Screen
	style    tcell.Style
	cursor   string
	bets     []rune
	deposit  []rune
	messages []string
	players  []*player
	dice     []int
	ante     float64
}

// New contructs a game board.
func New() (*Board, error) {
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
		screen:   screen,
		style:    style,
		messages: make([]string, 3),
	}

	return &board, nil
}

// Shutdown tearsdown the game board.
func (b *Board) Shutdown() {
	b.screen.Fini()
}

// Init generates the initial game board and starts the event loop.
func (b *Board) Init() {
	b.screen.Clear()

	for i := 1; i < boardWidth; i++ {
		b.screen.SetContent(i, 1, '=', nil, b.style)
	}
	for i := 1; i < boardWidth; i++ {
		b.screen.SetContent(i, boardHeight, '=', nil, b.style)
	}
	for i := 2; i < boardHeight; i++ {
		b.screen.SetContent(1, i, '|', nil, b.style)
	}
	for i := 2; i < boardHeight; i++ {
		b.screen.SetContent(boardWidth-1, i, '|', nil, b.style)
	}

	for i := 1; i < boardWidth; i++ {
		b.screen.SetContent(i, messageHeight, '=', nil, b.style)
	}

	b.print(3, messageHeight, " Message Center ")
	b.print(playersX, columnHeight, "Players:")
	b.print(betX, columnHeight, "Last Bet:")
	b.print(balX, columnHeight, "  Balances:")
	b.print(myDiceX-9, myDiceY, "My Dice:")
	b.print(anteX-6, anteY, "Ante:")
	b.print(potX-6, potY, "Pot :")
	b.print(betRowX-9, betRowY, "My Bet :>")
	b.print(depositRowX-10, depositRowY, "Deposit :>")
	b.print(helpX, 2, "<1-6>    : set/increment bet")
	b.print(helpX, 3, "<delete> : decrement bet")
	b.print(helpX, 4, "<b>      : place bet")
	b.print(helpX, 5, "<l>      : call liar")
	b.print(helpX, 6, "<d>      : deposit funds")
	b.print(helpX, configY, "network  :")
	b.print(helpX, configY+1, "chainid  :")
	b.print(helpX, configY+2, "contract :")
	b.print(helpX, configY+3, "address  :")

	b.betMode()
}

// PrintMessage adds a message to the message center.
func (b *Board) PrintMessage(message string) {
	const width = boardWidth - 4
	msg := fmt.Sprintf("%-*s", width, message)

	b.messages[2] = b.messages[1]
	b.messages[1] = b.messages[0]
	b.messages[0] = msg

	b.print(3, messageHeight+2, b.messages[0])
	b.print(3, messageHeight+3, b.messages[1])
	b.print(3, messageHeight+4, b.messages[2])

	b.screen.Show()
}

// StartEventLoop starts a goroutine to handle keyboard input.
func (b *Board) StartEventLoop() chan struct{} {
	quit := make(chan struct{})

	go func() {
		for {
			ev := b.screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					close(quit)
					return

				case tcell.KeyDEL:
					b.delete()

				case tcell.KeyEnter:
					b.enter()

				case tcell.KeyRune:
					b.processKeyEvent(ev.Rune())
				}
			}
		}
	}()

	return quit
}

// AddPlayer adds a player to the list of players.
func (b *Board) AddPlayer(address string, balance string) {
	b.players = append(b.players, &player{address: address, balance: balance})
	b.printPlayers()
}

// NextPlayer sets the player as the active player.
func (b *Board) ActivePlayer(address string) {
	for i := range b.players {
		b.players[i].active = false
		if b.players[i].address == address {
			b.players[i].active = true
		}
	}
	b.printPlayers()
}

// UpdatePlayerBet updates the players last bet.
func (b *Board) UpdatePlayerBet(address string, lastBet string) {
	for i := range b.players {
		if b.players[i].address == address {
			b.players[i].lastBet = lastBet
			break
		}
	}
	b.printPlayers()
}

// SetDice captures the dice for the player and displays them.
func (b *Board) SetDice(dice []int) error {
	if len(dice) != 5 {
		return errors.New("five dice are required")
	}

	b.dice = dice
	b.printDice()

	return nil
}

// SetAnte displays the ante information.
func (b *Board) SetAnte(ante float64) {
	b.ante = ante
	b.printAnte()
}

// SetSettings draws the settings on the board.
func (b *Board) SetSettings(network string, chainID int, contractID string, address string) {
	b.print(helpX+11, configY, network)
	b.print(helpX+11, configY+1, fmt.Sprintf("%d", chainID))
	b.print(helpX+11, configY+2, contractID)
	b.print(helpX+11, configY+3, address)
}

// =============================================================================

// processKeyEvent is the first line of processing for any key that is
// pressed during the game.
func (b *Board) processKeyEvent(r rune) {
	switch {
	case r == rune('d'):
		b.depositMode()

	case r == rune('b'):
		b.betMode()

	case (r >= rune('0') && r <= rune('9')) || r == rune('.'):
		b.value(r)

	default:
		b.screen.Beep()
	}
}

var i int

// enter is called to submit a bet or deposit.
func (b *Board) enter() {
	switch b.cursor {
	case "bet":
		b.betMode()

	case "deposit":
		b.depositMode()

	default:
		b.screen.Beep()
	}
}

// delete is called to remove the latest value from the bet or deposit.
func (b *Board) delete() {
	switch b.cursor {
	case "bet":
		b.subBet()

	case "deposit":
		b.subDeposit()

	default:
		b.screen.Beep()
	}
}

// value processes the keystroke based on the mode.
func (b *Board) value(r rune) {
	switch b.cursor {
	case "bet":
		if r >= rune('1') && r <= rune('6') {
			b.addBet(r)
			return
		}
		b.screen.Beep()

	case "deposit":
		b.addDeposit(r)

	default:
		b.screen.Beep()
	}
}

// =============================================================================

// depositMode puts the UI into the mode to accept deposit information and
// process a deposit.
func (b *Board) depositMode() {
	b.cursor = "deposit"
	b.deposit = []rune{}

	b.screen.ShowCursor(depositRowX+1, depositRowY)
	b.screen.SetContent(depositRowX, depositRowY, ' ', nil, b.style)
	b.screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)

	b.print(depositRowX, depositRowY, "                             ")
	b.print(betRowX, betRowY, "                             ")
}

// addDeposit takes the value selected on the keyboard and adds it to the
// deposit slice and screen.
func (b *Board) addDeposit(r rune) {
	if b.cursor != "deposit" {
		b.screen.Beep()
		return
	}

	x := depositRowX
	b.deposit = append(b.deposit, r)
	x += len(b.deposit)

	b.screen.ShowCursor(x+1, depositRowY)
	b.print(x, depositRowY, string(r))
}

// subDeposit removes a value from the deposit slice and screen.
func (b *Board) subDeposit() {
	if b.cursor != "deposit" || len(b.deposit) == 0 {
		b.screen.Beep()
		return
	}

	x := depositRowX
	x += len(b.deposit)
	b.deposit = b.deposit[:len(b.deposit)-1]

	b.screen.ShowCursor(x, depositRowY)
	b.print(x, depositRowY, " ")
}

// =============================================================================

// betMode puts the UI into the mode to accept bet information and
// process a bet.
func (b *Board) betMode() {
	b.cursor = "bet"
	b.bets = []rune{}

	b.screen.ShowCursor(betRowX+1, betRowY)
	b.screen.SetContent(betRowX, betRowY, ' ', nil, b.style)
	b.screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)

	b.print(depositRowX, depositRowY, "                             ")
	b.print(betRowX, betRowY, "                             ")
}

// addBet takes the value selected on the keyboard and adds it to the
// bet slice and screen.
func (b *Board) addBet(r rune) {
	if b.cursor != "bet" || (len(b.bets) > 0 && b.bets[0] != r) {
		b.screen.Beep()
		return
	}

	x := betRowX
	b.bets = append(b.bets, r)
	x += len(b.bets)

	b.screen.ShowCursor(x+1, betRowY)
	b.print(x, betRowY, string(r))
}

// subBet removes a value from the bet slice and screen.
func (b *Board) subBet() {
	if b.cursor != "bet" || len(b.bets) == 0 {
		b.screen.Beep()
		return
	}

	x := betRowX
	x += len(b.bets)
	b.bets = b.bets[:len(b.bets)-1]

	b.screen.ShowCursor(x, betRowY)
	b.print(x, betRowY, " ")
}

// =============================================================================

// printPlayers draws the players information on the screen.
func (b *Board) printPlayers() {
	const balWidth = 15
	var pot float64

	for i, p := range b.players {
		pot += b.ante

		addrY := columnHeight + 2 + i
		addr := fmt.Sprintf("%s...%s", p.address[:5], p.address[39:])
		bal := fmt.Sprintf("%*s", balWidth, p.balance)
		if p.active {
			b.print(playersX, addrY, "->")
		}
		b.print(playersX+4, addrY, addr)
		b.print(betX, addrY, p.lastBet)
		b.print(boardWidth-(balWidth+2), addrY, bal)
	}

	b.printPot(pot)
	b.screen.Show()
}

// printPot draws the pot information on the screen.
func (b *Board) printPot(pot float64) {
	b.print(potX, potY, fmt.Sprintf("$%.2f", pot))
}

// printDice draws the players dice on the screen.
func (b *Board) printDice() {
	dice := fmt.Sprintf("[%d][%d][%d][%d][%d]", b.dice[0], b.dice[1], b.dice[2], b.dice[3], b.dice[4])
	b.print(myDiceX, myDiceY, dice)
}

// printAnte draws the ante on the board.
func (b *Board) printAnte() {
	b.print(anteX, anteY, fmt.Sprintf("$%.2f", b.ante))
}

// print knows how to print a string on the screen.
func (b *Board) print(x, y int, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		b.screen.SetContent(x, y, c, comb, b.style)
		x += w
	}
	b.screen.Show()
}
