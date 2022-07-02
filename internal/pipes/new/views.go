package new

import "fmt"

func platformsView(m model) string {
	c := m.PlatformChoice

	tpl := "Which platform do you want to build your bot for?\n\n"
	tpl += "%s\n\n"
	tpl += subtle.Render("j/k, up/down: select") + dot + subtle.Render("enter: choose") + dot + subtle.Render("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n%s",
		checkbox("Discord", c == 0),
		checkbox("Telegram", c == 1),
		checkbox("Slack", c == 2),
	)

	return fmt.Sprintf(tpl, choices)
}

func langsView(m model) string {
	l := m.LangChoice

	c := ""

	if m.PlatformChoice == 0 {
		c = "Discord"
	} else if m.PlatformChoice == 1 {
		c = "Telegram"
	} else if m.PlatformChoice == 2 {
		c = "Slack"
	}

	tpl := "Choose language/framework for your " + c + " bot\n\n"
	tpl += "%s\n\n"
	tpl += subtle.Render("j/k, up/down: select") + dot + subtle.Render("enter: choose") + dot + subtle.Render("q, esc: quit")

	var n = func() string {
		if m.PlatformChoice == 2 {
			return checkbox("Node.js", l == 1)
		} else {
			return fmt.Sprintf(
				"%s\n%s",
				checkbox("Golang", l == 1),
				checkbox("Node.js", l == 2),
			)
		}
	}

	langs := fmt.Sprintf(
		"%s\n%s",
		checkbox("Python", l == 0),
		n(),
	)

	if m.PlatformChoice != 2 {
		langs += fmt.Sprintf(
			"\n%s\n%s\n%s",
			checkbox("Ruby", l == 3),
			checkbox("Rust", l == 4),
			checkbox("Deno", l == 5),
		)
	}

	if m.PlatformChoice == 0 {
		langs += fmt.Sprintf(
			"\n%s\n%s\n%s",
			checkbox("C#", l == 6),
			checkbox("Crystal", l == 7),
			checkbox("Dart", l == 8),
		)
	}

	return fmt.Sprintf(tpl, langs)
}

func pmsView(m model) string {
	pm := m.PMCoice

	l := ""

	if m.LangChoice == 0 {
		l = "Python"
	} else if m.LangChoice == 1 {
		l = "Golang"
	} else if m.LangChoice == 2 {
		l = "Node.js"
	} else if m.LangChoice == 3 {
		l = "Ruby"
	} else if m.LangChoice == 4 {
		l = "Rust"
	} else if m.LangChoice == 5 {
		l = "Deno"
	} else if m.LangChoice == 6 {
		l = "C#"
	} else if m.LangChoice == 7 {
		l = "Crystal"
	} else if m.LangChoice == 8 {
		l = "Dart"
	}

	tpl := "Choose your favorite package manager for " + l + "\n\n"
	tpl += "%s\n\n"
	tpl += subtle.Render("j/k, up/down: select") + dot + subtle.Render("enter: choose") + dot + subtle.Render("q, esc: quit")

	langs := ""
	nodePms := fmt.Sprintf(
		"%s\n%s\n%s",
		checkbox("npm", pm == 0),
		checkbox("yarn", pm == 1),
		checkbox("pnpm", pm == 2),
	)
	rubyPM := checkbox("bundler", pm == 0)

	if m.LangChoice == 0 {
		langs += fmt.Sprintf(
			"%s\n%s",
			checkbox("pip", pm == 0),
			checkbox("pipenv", pm == 1),
		)
	} else if m.LangChoice == 1 {
		if m.PlatformChoice == 2 {
			langs += nodePms
		} else {
			langs += checkbox("go mod", pm == 0)
		}
	} else if m.LangChoice == 2 {
		if m.PlatformChoice == 2 {
			langs += rubyPM
		} else {
			langs += nodePms
		}
	} else if m.LangChoice == 3 {
		langs += rubyPM
	} else if m.LangChoice == 4 {
		langs += fmt.Sprintf(
			"%s\n%s",
			checkbox("cargo", pm == 0),
			checkbox("fleet", pm == 1),
		)
	} else if m.LangChoice == 5 {
		langs += checkbox("deno", pm == 0)
	} else if m.LangChoice == 6 {
		langs += checkbox("dotnet", pm == 0)
	} else if m.LangChoice == 7 {
		langs += checkbox("shards", pm == 0)
	} else if m.LangChoice == 8 {
		langs += checkbox("pub", pm == 0)
	}

	return fmt.Sprintf(tpl, langs)
}

// func tokenView(m model) string {
// 	instructionsIn := `# Setup Discord Bot Token

// > Follow Instructions at [abdfnx/botway/discussions](https://github.com/abdfnx/botway/discussions/4)` + "\n\n"

// 	instructions, err := glamour.Render(instructionsIn, "dark")

// 	if err != nil {
// 		return err.Error()
// 	}

// 	tpl := instructions

// 	return tpl
// }

func finalView(m model) string {
	var platform, lang, pm string

	switch m.PlatformChoice {
		case 0:
			platform = "Discord"

		case 1:
			platform = "Telegram"

		case 2:
			platform = "Slack"
	}

	switch m.LangChoice {
		case 0:
			lang = "Python"

			switch m.PMCoice {
				case 0:
					pm = "pip"

				case 1:
					pm = "pipenv"
			}

		case 1:
			if m.PlatformChoice == 2 {
				lang = "Node.js"

				switch m.PMCoice {
					case 0:
						pm = "npm"

					case 1:
						pm = "yarn"

					case 2:
						pm = "pnpm"
				}
			} else {
				lang = "Golang"
				pm = "go mod"
			}

		case 2:
			lang = "Node.js"

			switch m.PMCoice {
				case 0:
					pm = "npm"

				case 1:
					pm = "yarn"

				case 2:
					pm = "pnpm"
			}

		case 3:
			lang = "Ruby"
			pm = "bundler"

		case 4:
			lang = "Rust"

			switch m.PMCoice {
				case 0:
					pm = "cargo"

				case 1:
					pm = "fleet"
			}

		case 5:
			lang = "Deno"
			pm = "deno"

		case 6:
			lang = "C#"
			pm = "dotnet"
		
		case 7:
			lang = "Crystal"
			pm = "shards"

		case 8:
			lang = "Dart"
			pm = "pub"
	}

	return "\n🤖 Noice, you're going to build a " + prim.Render(platform)  + " bot via " + prim.Render(lang) +  " with " + prim.Render(pm) + " package manager\n"
}
