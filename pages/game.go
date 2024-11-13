package pages

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
)

func GameScreen(renderer *lipgloss.Renderer, pty ssh.Pty) gameScreenModel {
	const (
		boardWidth  = 9
		boardHeight = 9
	)
	model := gameScreenModel{
		boardWidth:   boardWidth,
		boardHeight:  boardHeight,
		whitePawnPos: pawnPos{0, boardWidth / 2},
		blackPawnPos: pawnPos{boardHeight - 1, boardWidth / 2},
	}
	for i := 0; i < boardHeight; i++ {
		model.fences = append(model.fences, []fenceState{})
		for j := 0; j < boardWidth; j++ {
			model.fences[i] = append(model.fences[i], fenceState{false, false, false, false})
		}
	}
	return model
}

type pawnPos struct {
	row, col int
}

type fenceState struct {
	fenceLeft, fenceRight, fenceUp, fenceDown bool
}

type gameScreenModel struct {
	boardWidth  int
	boardHeight int

	fences       [][]fenceState
	whitePawnPos pawnPos
	blackPawnPos pawnPos
}

func (m gameScreenModel) Init() tea.Cmd {
	return nil
}

func (m gameScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m gameScreenModel) View() string {
	board := "╔═══"
	for i := 0; i < m.boardWidth-1; i++ {
		board += "╦═══"
	}
	board += "╗\n"

	for i := 0; i < m.boardHeight; i++ {
		board += "║ "
		for j := 0; j < m.boardWidth; j++ {
			if m.whitePawnPos.row == i && m.whitePawnPos.col == j {
				board += "W"
			} else if m.blackPawnPos.row == i && m.blackPawnPos.col == j {
				board += "B"
			} else {
				board += " "
			}
			if j == m.boardWidth-1 {
				continue
			}
			if m.fences[i][j].fenceRight || m.fences[i][j+1].fenceLeft {
				board += " ║ "
			} else {
				board += " │ "
			}
		}
		board += " ║\n"
		if i == m.boardHeight-1 {
			continue
		}
		board += "╠"
		for j := 0; j < m.boardWidth; j++ {
			if m.fences[i][j].fenceDown || m.fences[i+1][j].fenceUp {
				board += "═══"
			} else {
				board += "───"
			}
			if j == m.boardWidth-1 {
				continue
			}
			board += "┼"
		}
		board += "╣\n"
	}
	board += "╚═══"
	for i := 0; i < m.boardWidth-1; i++ {
		board += "╩═══"
	}
	board += "╝\n"
	return board
}
