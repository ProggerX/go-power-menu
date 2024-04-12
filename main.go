package main

import (
    "fmt"
	"os/exec"
	"github.com/fatih/color"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices []string
	cursor [2]int
	is_confirm int
	l [2]int
}

func do(what int) {
	switch what {
	case 0:
		(exec.Command("sh", "lockscreen")).Run()
	case 1:
		(exec.Command("systemctl", "poweroff")).Run()
	case 2:
		(exec.Command("systemctl", "reboot")).Run()
	}
}

func initialModel() model {
	return model {
		choices: []string{
			"Lock screen (will run \"sh lockscreen\")",
			"Shut down computer",
			"Reboot computer",
		},
		cursor: [2]int{0, 1},
		is_confirm: 0,
		l: [2]int{2, 1},
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
            if m.cursor[m.is_confirm] > 0 {
                m.cursor[m.is_confirm]--
            }

        case "down", "j":
            if m.cursor[m.is_confirm] < m.l[m.is_confirm] {
                m.cursor[m.is_confirm]++
            }
		case "enter", "right", "l":
			if m.is_confirm == 0 {
				m.is_confirm = 1
			} else {
				if m.cursor[m.is_confirm] == 1 {
					m.is_confirm = 0
					m.cursor[1] = 1
				} else {
					do(m.cursor[0])
					return m, tea.Quit
				}
			}
		case "left", "h":
			m.is_confirm = 0
			m.cursor[1] = 1
		}
	}

    return m, nil
}

func (m model) View() string {
	var s string
	if m.is_confirm == 0 {
	    s = color.WhiteString("What to do?\n\n")
	
	    for i, choice := range m.choices {
	        cursor := " " 
	        if m.cursor[0] == i {
	            cursor = ">" 
				s += fmt.Sprintf(color.BlueString("%s %s\n", cursor, choice))
	        } else {
	        	s += fmt.Sprintf(color.WhiteString("%s %s\n", cursor, choice))
			}
	    }
	} else {
	    s = color.WhiteString("Are you sure?\n\n")
	
	    for i, choice := range []string{"Yes", "No"} {
	        cursor := " " 
	        if m.cursor[1] == i {
	            cursor = ">" 
				s += fmt.Sprintf(color.BlueString("%s %s\n", cursor, choice))
	        } else {
	        	s += fmt.Sprintf(color.WhiteString("%s %s\n", cursor, choice))
			}
	    }
	}

    return s
}

func (m model) Init() tea.Cmd {
	return nil
}

func main() {
    powermenu := tea.NewProgram(initialModel())
	powermenu.Run()
}
