package pages

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
)

func RootScreen(renderer *lipgloss.Renderer, pty ssh.Pty) rootScreenModel {
	return rootScreenModel{
		currentScreen: HomeScreen(renderer, pty),
	}
}

type rootScreenModel struct {
	currentScreen tea.Model
}

func (m rootScreenModel) Init() tea.Cmd {
	return m.currentScreen.Init()
}

func (m rootScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.currentScreen.Update(msg)
}

func (m rootScreenModel) View() string {
	return m.currentScreen.View()
}

func (m rootScreenModel) switchScreen(model tea.Model) (tea.Model, tea.Cmd) {
	return model, model.Init()
}
