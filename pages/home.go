package pages

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
)

func HomeScreen(renderer *lipgloss.Renderer, pty ssh.Pty) homeModel {
	return homeModel{
		width:  pty.Window.Width,
		height: pty.Window.Height,

		renderer: renderer,
		pty:      pty,
	}
}

type homeModel struct {
	width  int
	height int

	renderer *lipgloss.Renderer
	pty      ssh.Pty
}

func (m homeModel) Init() tea.Cmd {
	return nil
}

func (m homeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
        case "n":
            return rootScreenModel{}.switchScreen(GameScreen(m.renderer, m.pty))
		}
	}
	return m, nil
}

func (m homeModel) View() string {
	return "Welcome to Quoriterm!!!\n\nPress n for a new game....."
}
