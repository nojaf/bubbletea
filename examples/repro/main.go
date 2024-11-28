package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type model struct {
	v string
}

func NewModel() model {
	return model{v: ""}
}

func (m model) Init() (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:
		log.Println("key pressed:", msg.String())

		switch msg.String() {
		case "ctrl+c", "enter":
			return m, tea.Quit
		default:
			m.v += msg.String()
		}
	}

	return m, nil
}

func (m model) View() string {
	return "> " + m.v
}

func main() {
	file := "debug.log"

	// Check if the file exists
	if _, err := os.Stat(file); err == nil {
		// File exists, attempt to remove it
		if err := os.Remove(file); err != nil {
			// Handle error during removal
			panic(err)
		}
	} else if !os.IsNotExist(err) {
		// Handle unexpected error during Stat
		panic(err)
	}

	f, err := tea.LogToFile(file, "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	p := tea.NewProgram(NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	p2 := tea.NewProgram(NewModel())
	if _, err := p2.Run(); err != nil {
		fmt.Printf("Alas, there's been an error 2: %v", err)
		os.Exit(1)
	}
}
