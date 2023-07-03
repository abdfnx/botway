package new

import "fmt"

func platformsView(m model) string {
	c := m.PlatformChoice

	tpl := "Which platform do you want to build your bot for?\n\n"
	tpl += "%s\n\n"
	tpl += subtle.Render("j/k, up/down: select") + dot + subtle.Render("enter: choose") + dot + subtle.Render("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		checkbox("Discord", c == 0),
		checkbox("Telegram", c == 1),
		checkbox("Slack", c == 2),
		checkbox("Twitch", c == 3),
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
	} else if m.PlatformChoice == 3 {
		c = "Twitch"
	}

	tpl := "Choose language/framework for your " + c + " bot\n\n"
	tpl += "%s\n\n"
	tpl += subtle.Render("j/k, up/down: select") + dot + subtle.Render("enter: choose") + dot + subtle.Render("q, esc: quit")

	var n = func() string {
		if m.PlatformChoice == 2 {
			return fmt.Sprintf(
				"%s\n%s",
				checkbox("Node.js", l == 1),
				checkbox("Node.js (TypeScript)", l == 2),
			)
		} else {
			return fmt.Sprintf(
				"%s\n%s\n%s",
				checkbox("Golang", l == 1),
				checkbox("Node.js", l == 2),
				checkbox("Node.js (TypeScript)", l == 3),
			)
		}
	}

	langs := fmt.Sprintf(
		"%s\n%s",
		checkbox("Python", l == 0),
		n(),
	)

	if m.PlatformChoice != 2 && m.PlatformChoice != 3 {
		l14 := "C"

		if m.PlatformChoice == 1 {
			l14 = "Swift"
		}

		langs += fmt.Sprintf(
			"\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s",
			checkbox("Ruby", l == 4),
			checkbox("Rust", l == 5),
			checkbox("Deno", l == 6),
			checkbox("C#", l == 7),
			checkbox("Dart", l == 8),
			checkbox("PHP", l == 9),
			checkbox("Kotlin", l == 10),
			checkbox("Java", l == 11),
			checkbox("C++", l == 12),
			checkbox("Nim", l == 13),
			checkbox(l14, l == 14),
		)
	}

	if m.PlatformChoice == 3 {
		langs += fmt.Sprintf(
			"\n%s\n%s",
			checkbox("Deno", l == 4),
			checkbox("Java", l == 5),
		)
	}

	if m.PlatformChoice == 0 {
		langs += fmt.Sprintf(
			"\n%s",
			checkbox("Crystal", l == 15),
		)
	}

	return fmt.Sprintf(tpl, langs)
}

func pmsView(m model) string {
	pm := m.PMChoice

	l := ""

	if m.LangChoice == 0 {
		l = "Python"
	} else if m.LangChoice == 1 {
		l = "Golang"
	} else if m.LangChoice == 2 {
		l = "Node.js"
	} else if m.LangChoice == 3 {
		l = "Node.js (TypeScript)"
	} else if m.LangChoice == 4 {
		if m.PlatformChoice == 3 {
			l = "Deno"
		} else {
			l = "Ruby"
		}
	} else if m.LangChoice == 5 {
		if m.PlatformChoice == 3 {
			l = "Java"
		} else {
			l = "Rust"
		}
	} else if m.LangChoice == 6 {
		l = "Deno"
	} else if m.LangChoice == 7 {
		l = "C#"
	} else if m.LangChoice == 8 {
		l = "Dart"
	} else if m.LangChoice == 9 {
		l = "PHP"
	} else if m.LangChoice == 10 {
		l = "Kotlin"
	} else if m.LangChoice == 11 {
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

	tpl := "Choose your favorite package manager for " + l + "\n\n"
	tpl += "%s\n\n"
	tpl += subtle.Render("j/k, up/down: select") + dot + subtle.Render("enter: choose") + dot + subtle.Render("q, esc: quit")

	langs := ""
	nodePms := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		checkbox("npm", pm == 0),
		checkbox("yarn", pm == 1),
		checkbox("pnpm", pm == 2),
		checkbox("bun", pm == 3),
	)
	rubyPM := checkbox("bundler", pm == 0)

	if m.LangChoice == 0 {
		langs += fmt.Sprintf(
			"%s\n%s\n%s",
			checkbox("pip", pm == 0),
			checkbox("pipenv", pm == 1),
			checkbox("poetry", pm == 2),
		)
	} else if m.LangChoice == 1 {
		if m.PlatformChoice == 2 {
			langs += nodePms
		} else {
			langs += checkbox("go mod", pm == 0)
		}
	} else if m.LangChoice == 2 || m.LangChoice == 3 {
		langs += nodePms
	} else if m.LangChoice == 4 {
		if m.PlatformChoice == 3 {
			langs += checkbox("deno", pm == 0)
		} else {
			langs += rubyPM
		}
	} else if m.LangChoice == 5 {
		if m.PlatformChoice == 3 {
			langs += checkbox("gradle", pm == 0)
		} else {
			langs += fmt.Sprintf(
				"%s\n%s",
				checkbox("cargo", pm == 0),
			)
		}
	} else if m.LangChoice == 6 {
		langs += checkbox("deno", pm == 0)
	} else if m.LangChoice == 7 {
		langs += checkbox("dotnet", pm == 0)
	} else if m.LangChoice == 8 {
		langs += checkbox("pub", pm == 0)
	} else if m.LangChoice == 9 {
		langs += checkbox("composer", pm == 0)
	} else if m.LangChoice == 10 || m.LangChoice == 11 {
		langs += checkbox("gradle", pm == 0)
	} else if m.LangChoice == 12 {
		langs += checkbox("cmake", pm == 0)
	} else if m.LangChoice == 13 {
		langs += checkbox("nimble", pm == 0)
	} else if m.LangChoice == 14 {
		if m.PlatformChoice == 1 {
			langs += checkbox("swift package", pm == 0)
		} else {
			langs += checkbox("continue", pm == 0)
		}
	} else if m.LangChoice == 15 {
		langs += checkbox("shards", pm == 0)
	}

	return fmt.Sprintf(tpl, langs)
}

func hostServicesView(m model) string {
	c := m.HostServiceChoice

	tpl := "Which hosting service do you want to use?\n\n"
	tpl += "%s\n\n"
	tpl += subtle.Render("j/k, up/down: select") + dot + subtle.Render("enter: choose") + dot + subtle.Render("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s",
		checkbox("Railway (railway.app)", c == 0),
		checkbox("Render (render.com)", c == 1),
	)

	return fmt.Sprintf(tpl, choices)
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
	var platform, lang, pm, hostService string

	switch m.PlatformChoice {
	case 0:
		platform = "Discord"

	case 1:
		platform = "Telegram"

	case 2:
		platform = "Slack"

	case 3:
		platform = "Twitch"
	}

	switch m.HostServiceChoice {
	case 0:
		hostService = "Railway"

	case 1:
		hostService = "Render"
	}

	switch m.LangChoice {
	case 0:
		lang = "Python"

		switch m.PMChoice {
		case 0:
			pm = "pip"

		case 1:
			pm = "pipenv"

		case 2:
			pm = "poetry"
		}

	case 1:
		if m.PlatformChoice == 2 {
			lang = "Node.js"

			switch m.PMChoice {
			case 0:
				pm = "npm"

			case 1:
				pm = "yarn"

			case 2:
				pm = "pnpm"

			case 3:
				pm = "bun"
			}
		} else {
			lang = "Golang"
			pm = "go mod"
		}

	case 2:
	case 3:
		lang = "Node.js"

		if m.LangChoice == 3 {
			lang = "Node.js (TypeScript)"
		}

		switch m.PMChoice {
		case 0:
			pm = "npm"

		case 1:
			pm = "yarn"

		case 2:
			pm = "pnpm"

		case 3:
			pm = "bun"
		}

	case 4:
		if m.PlatformChoice == 3 {
			lang = "Deno"
			pm = "deno"
		} else {
			lang = "Ruby"
			pm = "bundler"
		}

	case 5:
		if m.PlatformChoice == 3 {
			lang = "Java"
			pm = "gradle"
		} else {
			lang = "Rust"

			switch m.PMChoice {
			case 0:
				pm = "cargo"
			}
		}

	case 6:
		lang = "Deno"
		pm = "deno"

	case 7:
		lang = "C#"
		pm = "dotnet"

	case 8:
		pm = "pub"
		lang = "Dart"

	case 9:
		lang = "PHP"
		pm = "composer"

	case 10:
		lang = "Kotlin"
		pm = "gradle"

	case 11:
		lang = "Java"
		pm = "gradle"

	case 12:
		lang = "C++"
		pm = "cmake"

	case 13:
		lang = "Nim"
		pm = "nimble"

	case 14:
		if m.PlatformChoice == 1 {
			lang = "Swift"
			pm = "swift package"
		} else {
			lang = "C"
		}

	case 15:
		lang = "Crystal"
		pm = "shards"
	}

	msg := "\n🤖 Noice, you're going to build a " + prim.Render(platform) + " bot via " + prim.Render(lang)
	hMsg := " and " + prim.Render(hostService) + " as hosting service"

	if lang == "C" {
		return msg + hMsg + "\n"
	} else {
		return msg + " with " + prim.Render(pm) + " package manager" + hMsg + "\n"
	}
}
