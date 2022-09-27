package new

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func updateHostServices(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.HostServiceChoice += 1
			if m.HostServiceChoice > 1 {
				m.HostServiceChoice = 2
			}

		case "k", "up":
			m.HostServiceChoice -= 1
			if m.HostServiceChoice < 0 {
				m.HostServiceChoice = 0
			}

		case "enter":
			m.HostService = true
			return m, frame()
		}
	}

	return m, nil
}

func updatePlatforms(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.PlatformChoice += 1
			if m.PlatformChoice > 3 {
				m.PlatformChoice = 3
			}

		case "k", "up":
			m.PlatformChoice -= 1
			if m.PlatformChoice < 0 {
				m.PlatformChoice = 0
			}

		case "enter":
			m.Platform = true
			return m, frame()
		}
	}

	return m, nil
}

func updateLangs(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.LangChoice += 1

			if m.PlatformChoice == 2 {
				if m.LangChoice > 2 {
					m.LangChoice = 2
				}
			} else if m.PlatformChoice == 3 {
				if m.LangChoice > 5 {
					m.LangChoice = 5
				} else {
					if m.PlatformChoice == 0 {
						if m.LangChoice > 15 {
							m.LangChoice = 15
						}
					} else {
						if m.LangChoice > 14 {
							m.LangChoice = 14
						}
					}
				}
			}

		case "k", "up":
			m.LangChoice -= 1

			if m.LangChoice < 0 {
				m.LangChoice = 0
			}

		case "enter":
			m.Lang = true

			return m, frame()
		}
	}

	return m, nil
}

func updatePMs(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.PMChoice += 1

			if m.LangChoice == 1 {
				if m.PlatformChoice == 2 {
					if m.PMChoice > 3 {
						m.PMChoice = 3
					}
				} else {
					if m.PMChoice > 0 {
						m.PMChoice = 0
					}
				}
			} else if m.LangChoice == 0 || m.LangChoice == 2 || m.LangChoice == 3 {
				if m.PMChoice > 3 {
					m.PMChoice = 3
				}
			} else {
				if m.PMChoice > 1 {
					m.PMChoice = 1
				}
			}

		case "k", "up":
			m.PMChoice -= 1
			if m.PMChoice < 0 {
				m.PMChoice = 0
			}

		case "enter":
			m.PM = true
			return m, frame()
		}
	}

	return m, nil
}

func buildBot(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	fmt.Println(finalView(m))

	l := ""

	if m.LangChoice == 0 {
		l = "Python"
	} else if m.LangChoice == 1 {
		if m.PlatformChoice == 2 {
			l = "Node"
		} else {
			l = "Go"
		}
	} else if m.LangChoice == 2 || m.LangChoice == 3 || m.LangChoice == 6 {
		l = "Node"
	} else if m.LangChoice == 4 {
		if m.PlatformChoice == 3 {
			l = "Node"
		} else {
			l = "Ruby"
		}
	} else if m.LangChoice == 5 {
		if m.PlatformChoice == 3 {
			l = "Java"
		} else {
			l = "Rust"
		}
	} else if m.LangChoice == 7 {
		l = "C#"
	} else if m.LangChoice == 8 {
		l = "Dart"
	} else if m.LangChoice == 9 {
		l = "PHP"
	} else if m.LangChoice == 10 || m.LangChoice == 11 {
		l = "Java"
	} else if m.LangChoice == 12 {
		l = "C++"
	} else if m.LangChoice == 13 {
		l = "Nim"
	} else if m.LangChoice == 14 {
		if m.PlatformChoice == 1 {
			l = "Swift"
		} else {
			l = "C"
		}
	} else if m.LangChoice == 15 {
		l = "Crystal"
	}

	NewBot(m, l, m.PlatformChoice, m.LangChoice)

	return m, tea.Quit
}
