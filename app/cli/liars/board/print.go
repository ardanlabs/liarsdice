package board

import (
	"fmt"
	"strings"

	"github.com/ardanlabs/liarsdice/app/cli/liars/engine"
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

// printBoard generates the initial game board and starts the event loop.
func (b *Board) printBoard() error {
	status, err := b.engine.QueryStatus()
	if err != nil {
		return err
	}

	b.screen.Clear()

	b.drawGameBoard(false)
	b.printLables()
	b.printSettings()

	b.print(helpX, 1, "<1-6>+   : set bet")
	b.print(helpX, 2, "<del>    : remove bet number")
	b.print(helpX, 3, "<l>      : call liar")
	b.print(helpX, 4, "<n>      : new game")
	b.print(helpX, 5, "<j>      : join game")
	b.print(helpX, 6, "<s>      : start game")

	b.print(helpX, statusY-6, "status   :")
	b.print(helpX, statusY-5, "round    :")
	b.print(helpX, statusY-4, "lastbet  :")
	b.print(helpX, statusY-3, "lastwin  :")
	b.print(helpX, statusY-2, "lastlose :")
	b.print(helpX, statusY, "engine   :")
	b.print(helpX, statusY+1, "blockchn :")
	b.print(helpX, statusY+2, "chainid  :")
	b.print(helpX, statusY+3, "contract :")
	b.print(helpX, statusY+4, "account  :")

	b.bets = []rune{}
	b.screen.ShowCursor(betRowX+1, betRowY)
	b.screen.SetContent(betRowX, betRowY, ' ', nil, b.style)
	b.screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
	b.print(betRowX, betRowY, "                 ")

	b.printStatus(status)

	return nil
}

// printStatus display the status information.
func (b *Board) printStatus(status engine.Status) {

	// Save this status for modal support.
	b.lastStatus = status

	// Print the current game status and round.
	b.print(helpX+11, statusY-6, fmt.Sprintf("%-10s", status.Status))
	b.print(helpX+11, statusY-5, fmt.Sprintf("%d   ", status.Round))

	// Show the account who last won and lost.
	if status.LastWinAcctID != "" {
		b.print(helpX+11, statusY-3, b.fmtAddress(status.LastWinAcctID))
		b.print(helpX+11, statusY-2, b.fmtAddress(status.LastOutAcctID))
	}

	// Show the last bet.
	if len(status.Bets) > 0 {
		bet := status.Bets[len(status.Bets)-1]
		betStr := fmt.Sprintf("%d %-10s", bet.Number, words[bet.Suite])
		b.print(helpX+11, statusY-4, betStr)
	} else {
		b.print(helpX+11, statusY-4, "                 ")
	}

	var pot float64

	// Clear the player lines.
	for i := 0; i < 5; i++ {
		addrY := columnHeight + 1 + i
		b.print(playersX, addrY, fmt.Sprintf("%-*s", boardWidth-4, " "))
		b.print(myDiceX, myDiceY, fmt.Sprintf("%-20s", " "))
	}

	// Print the player lines.
	for i, cup := range status.Cups {
		pot += status.AnteUSD

		// Players Column.
		addrY := columnHeight + 2 + i
		accountID := b.fmtAddress(cup.AccountID)
		b.print(playersX+3, addrY, accountID)

		// Outs.
		b.print(outX, addrY, fmt.Sprintf("%d", cup.Outs))

		// Show the active player.
		if i == status.CurrentCup {
			b.print(playersX, addrY, "->")
			b.print(playersX+3, addrY, accountID)
		} else {
			b.print(playersX, addrY, "  ")
		}

		// Last Bets.
		if cup.LastBet.Number != 0 {
			bet := fmt.Sprintf("%d %-10s", cup.LastBet.Number, words[cup.LastBet.Suite])
			b.print(betX, addrY, bet)
		} else {
			b.print(betX, addrY, "                 ")
		}

		// Balance Column.
		const balWidth = 15
		bal := fmt.Sprintf("%*s", balWidth, "$"+status.Balances[i])
		b.print(boardWidth-(balWidth+2), addrY, bal)

		// Show the dice for the connected account.
		if strings.EqualFold(cup.AccountID, b.accountID) {
			if cup.Dice[0] != 0 {
				dice := fmt.Sprintf("[%d][%d][%d][%d][%d]", cup.Dice[0], cup.Dice[1], cup.Dice[2], cup.Dice[3], status.Cups[i].Dice[4])
				b.print(myDiceX, myDiceY, dice)
			}
		}
	}

	// Show the ante and pot information.
	b.print(anteX, anteY, fmt.Sprintf("$%.2f", status.AnteUSD))
	b.print(potX, potY, fmt.Sprintf("$%.2f", pot))

	// Handle active player screen changes.
	if len(status.CupsOrder) > 0 {
		if status.CupsOrder[status.CurrentCup] == b.accountID {
			for x, r := range b.bets {
				b.print(betRowX+x+1, betRowY, string(r))
			}
			b.screen.ShowCursor(betRowX+len(b.bets)+1, betRowY)
			b.drawGameBoard(true)
		} else {
			b.screen.HideCursor()
			b.drawGameBoard(false)
		}
	}

	// Hide the cursor to show the game is over.
	if status.Status == "gameover" {
		b.screen.HideCursor()
	}

	// If the model was up, show it again.
	if b.modalUp {
		b.showModal(b.modalMsg)
	}

	b.screen.Show()
}

// printLables places the labels on the board.
func (b *Board) printLables() {
	b.print(playersX, columnHeight, "Players:")
	b.print(outX, columnHeight, "Outs:")
	b.print(betX, columnHeight, "Last Bet:")
	b.print(balX, columnHeight, "  Balances:")
	b.print(myDiceX-9, myDiceY, "My Dice:")
	b.print(anteX-6, anteY, "Ante:")
	b.print(potX-6, potY, "Pot :")
	b.print(potX-6, potY+1, "Bet :")
	b.print(betRowX-9, betRowY, "My Bet :>")
}

// printSettings draws the settings on the board.
func (b *Board) printSettings() {
	b.print(helpX+11, statusY, b.engine.URL())
	b.print(helpX+11, statusY+1, b.network)
	b.print(helpX+11, statusY+2, fmt.Sprintf("%d", b.chainID))
	b.print(helpX+11, statusY+3, b.fmtAddress(b.contractID))
	b.print(helpX+11, statusY+4, b.fmtAddress(b.accountID))
}

// PrintMessage adds a message to the message center.
func (b *Board) printMessage(message string, beep bool) {
	const width = boardWidth - 4
	msg := fmt.Sprintf("%-*s", width, message)
	if len(msg) > 58 {
		msg = msg[:58]
	}

	b.messages[4] = b.messages[3]
	b.messages[3] = b.messages[2]
	b.messages[2] = b.messages[1]
	b.messages[1] = b.messages[0]
	b.messages[0] = msg

	b.print(3, messageHeight+1, b.messages[0])
	b.print(3, messageHeight+2, b.messages[1])
	b.print(3, messageHeight+3, b.messages[2])
	b.print(3, messageHeight+4, b.messages[3])
	b.print(3, messageHeight+5, b.messages[4])

	if beep {
		b.screen.Beep()
	}

	b.screen.Show()

	if strings.Contains(message, "rolldice") ||
		strings.Contains(message, "bet") {
		return
	}

	b.showModal(message)
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

// fmtAddress provides a shortened version of an address.
func (*Board) fmtAddress(address string) string {
	if len(address) != 42 {
		return address
	}
	return fmt.Sprintf("%s..%s", address[:5], address[39:])
}

// =============================================================================

// drawGameBoard draws the game box.
func (b *Board) drawGameBoard(white bool) {
	x := 1
	y := 1
	width := boardWidth
	height := boardHeight

	style := b.style
	if white {
		style = style.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	} else {
		style = style.Background(tcell.ColorBlack).Foreground(tcell.ColorGrey)
	}

	// This places the message bar.
	for i := 1; i < boardWidth; i++ {
		b.screen.SetContent(i, messageHeight, hozTopRune, nil, style)
	}

	// This places the outer lines.
	for h := y; h < height; h++ {
		for w := x; w < width; w++ {
			if h == y {
				b.screen.SetContent(w, h, hozTopRune, nil, style)
			}
			if h == height-1 {
				b.screen.SetContent(w, h, hozBotRune, nil, style)
			}
			if w == x || w == width-1 {
				b.screen.SetContent(w, h, verRune, nil, style)
			}
		}
	}

	b.screen.Show()
}

// drawBox draws an empty box on the screen.
func (b *Board) drawBox(x int, y int, width int, height int) {
	style := b.style
	style = style.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

	for h := y; h < height; h++ {
		for w := x; w < width; w++ {
			b.screen.SetContent(w, h, ' ', nil, b.style)
		}
	}

	for h := y; h < height; h++ {
		for w := x; w < width; w++ {
			if h == y {
				b.screen.SetContent(w, h, hozTopRune, nil, style)
			}
			if h == height-1 {
				b.screen.SetContent(w, h, hozBotRune, nil, style)
			}
			if w == x || w == width-1 {
				b.screen.SetContent(w, h, verRune, nil, style)
			}
		}
	}

	b.screen.Show()
}
