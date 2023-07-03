package new

import (
	"fmt"
	"runtime"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/internal/options"
	"github.com/charmbracelet/lipgloss"
)

var (
	prim   = lipgloss.NewStyle().Foreground(constants.PRIMARY_COLOR)
	subtle = lipgloss.NewStyle().Foreground(constants.GRAY_COLOR)
	dot    = lipgloss.NewStyle().Foreground(constants.GRAY_COLOR).SetString(" • ").String()
	opts   = options.CommonOptions{
		BotName: "",
	}
)

func checkbox(label string, checked bool) string {
	if checked {
		return prim.Render("[✔] " + label)
	}

	return fmt.Sprintf("[ ] %s", label)
}

func HostServiceName(m model) string {
	if m.HostServiceChoice == 0 {
		return "railway"
	} else if m.HostServiceChoice == 1 {
		return "render"
	}

	return ""
}

func HostService(m model) string {
	if m.HostServiceChoice == 0 {
		return "railway.app"
	} else if m.HostServiceChoice == 1 {
		return "render.com"
	}

	return "# Specify the host service here"
}

func BotType(m model) string {
	if m.PlatformChoice == 0 {
		return "discord"
	} else if m.PlatformChoice == 1 {
		return "telegram"
	} else if m.PlatformChoice == 2 {
		return "slack"
	} else if m.PlatformChoice == 3 {
		return "twitch"
	}

	return "# You need to specify a platform (discord, telegram, slack, twitch)"
}

var blankLangMessage = "# You need to specify a language (python, go, nodejs, nodejs (typescript) ruby, rust, deno, csharp, dart, php, kotlin, java, crystal, c++, nim, c)"

func BotLang(m model) string {
	if m.LangChoice == 0 {
		return "python"
	} else if m.LangChoice == 1 {
		if m.PlatformChoice == 2 {
			return "nodejs"
		} else {
			return "go"
		}
	} else if m.LangChoice == 2 {
		return "nodejs"
	} else if m.LangChoice == 3 {
		return "typescript"
	} else if m.LangChoice == 4 {
		if m.PlatformChoice == 3 {
			return "deno"
		} else {
			return "ruby"
		}
	} else if m.LangChoice == 5 {
		if m.PlatformChoice == 3 {
			return "java"
		} else {
			return "rust"
		}
	} else if m.LangChoice == 6 {
		return "deno"
	} else if m.LangChoice == 7 {
		return "csharp"
	} else if m.LangChoice == 8 {
		return "dart"
	} else if m.LangChoice == 9 {
		return "php"
	} else if m.LangChoice == 10 {
		return "kotlin"
	} else if m.LangChoice == 11 {
		return "java"
	} else if m.LangChoice == 12 {
		return "cpp"
	} else if m.LangChoice == 13 {
		return "nim"
	} else if m.LangChoice == 14 {
		if m.PlatformChoice == 1 {
			return "swift"
		} else {
			return "c"
		}
	} else if m.LangChoice == 15 {
		return "crystal"
	}

	return blankLangMessage
}

func BotStartCmd(m model) string {
	nodeCmd := BotPM(m) + " start"
	denoCmd := "deno task run"

	if m.LangChoice == 0 && m.PMChoice == 0 {
		if runtime.GOOS == "windows" {
			return `py .\src\main.py`
		} else {
			return `python3 ./src/main.py`
		}
	} else if m.LangChoice == 0 && m.PMChoice == 1 {
		if runtime.GOOS == "windows" {
			return `pipenv run py .\src\main.py`
		} else {
			return `pipenv run python3 ./src/main.py`
		}
	} else if m.LangChoice == 0 && m.PMChoice == 2 {
		if runtime.GOOS == "windows" {
			return `poetry run .\src\main.py`
		} else {
			return `poetry run ./src/main.py`
		}
	} else if m.LangChoice == 1 {
		if m.PlatformChoice == 2 {
			return nodeCmd
		} else {
			return "go run src/main.go"
		}
	} else if m.LangChoice == 2 {
		return nodeCmd
	} else if m.LangChoice == 3 {
		return nodeCmd
	} else if m.LangChoice == 4 {
		if m.PlatformChoice == 3 {
			return denoCmd
		} else {
			return "bundle exec ruby src/main.rb"
		}
	} else if m.LangChoice == 5 {
		if m.PlatformChoice == 3 {
			if runtime.GOOS == "windows" {
				return `.\gradlew.bat run`
			} else {
				return "./gradlew run"
			}
		} else {
			return "cargo run src/main.rs"
		}
	} else if m.LangChoice == 6 {
		return denoCmd
	} else if m.LangChoice == 7 {
		return "dotnet run"
	} else if m.LangChoice == 8 {
		return "dart run src/main.dart"
	} else if m.LangChoice == 9 {
		return "php src/main.php"
	} else if m.LangChoice == 10 || m.LangChoice == 11 {
		if runtime.GOOS == "windows" {
			return `.\gradlew.bat run`
		} else {
			return "./gradlew run"
		}
	} else if m.LangChoice == 12 {
		if runtime.GOOS == "windows" {
			return `.\run.ps1`
		} else {
			return "cd build; make -j; ./" + opts.BotName
		}
	} else if m.LangChoice == 14 {
		if m.PlatformChoice == 1 {
			return "swift run"
		} else {
			if runtime.GOOS == "windows" {
				return `.\run.ps1`
			} else {
				return "gcc src/main.c -o bot -pthread -ldiscord -lcurl; ./bot"
			}
		}
	} else if m.LangChoice == 13 {
		return "nim c -r src/main.nim"
	} else if m.LangChoice == 15 {
		return "crystal run src/main.cr"
	}

	return "# Write your start command here"
}

func BotPM(m model) string {
	var nodePMs = func() string {
		if m.PMChoice == 1 {
			return "yarn"
		} else if m.PMChoice == 2 {
			return "pnpm"
		} else if m.PMChoice == 3 {
			return "bun"
		}

		return "npm"
	}

	if m.LangChoice == 0 && m.PMChoice == 0 {
		return "pip"
	} else if m.LangChoice == 0 && m.PMChoice == 1 {
		return "pipenv"
	} else if m.LangChoice == 0 && m.PMChoice == 2 {
		return "poetry"
	} else if m.LangChoice == 1 {
		return "go mod"
	} else if m.LangChoice == 2 || m.LangChoice == 3 {
		nodePMs()
	} else if m.LangChoice == 4 {
		if m.PlatformChoice == 3 {
			return "deno"
		} else {
			return "bundler"
		}
	} else if m.LangChoice == 5 {
		if m.PlatformChoice == 3 {
			return "gradle"
		} else {
			return "cargo"
		}
	} else if m.LangChoice == 6 {
		return "deno"
	} else if m.LangChoice == 7 {
		return "dotnet"
	} else if m.LangChoice == 8 {
		return "pub"
	} else if m.LangChoice == 9 {
		return "composer"
	} else if m.LangChoice == 10 || m.LangChoice == 11 {
		return "gradle"
	} else if m.LangChoice == 12 {
		return "cmake"
	} else if m.LangChoice == 13 {
		return "nimble"
	} else if m.LangChoice == 14 {
		if m.PlatformChoice == 1 {
			return "swift"
		} else {
			return "continue"
		}
	} else if m.LangChoice == 15 {
		return "shards"
	}

	return "# Specify the package manager here"
}

func CSharpGitIgnore() string {
	return `## Ignore Visual Studio temporary files, build results, and
## files generated by popular Visual Studio add-ons.

# User-specific files
*.suo
*.user
*.userosscache
*.sln.docstates

# User-specific files (MonoDevelop/Xamarin Studio)
*.userprefs

# Build results
[Dd]ebug/
[Dd]ebugPublic/
[Rr]elease/
[Rr]eleases/
build/
bld/
[Bb]in/
[Oo]bj/

# Visual Studo 2015 cache/options directory
.vs/

# MSTest test Results
[Tt]est[Rr]esult*/
[Bb]uild[Ll]og.*

# NUNIT
*.VisualState.xml
TestResult.xml

# Build Results of an ATL Project
[Dd]ebugPS/
[Rr]eleasePS/
dlldata.c

*_i.c
*_p.c
*_i.h
*.ilk
*.meta
*.obj
*.pch
*.pdb
*.pgc
*.pgd
*.rsp
*.sbr
*.tlb
*.tli
*.tlh
*.tmp
*.tmp_proj
*.log
*.vspscc
*.vssscc
.builds
*.pidb
*.svclog
*.scc

# Chutzpah Test files
_Chutzpah*

# Visual C++ cache files
ipch/
*.aps
*.ncb
*.opensdf
*.sdf
*.cachefile

# Visual Studio profiler
*.psess
*.vsp
*.vspx

# TFS 2012 Local Workspace
$tf/

# Guidance Automation Toolkit
*.gpState

# ReSharper is a .NET coding add-in
_ReSharper*/
*.[Rr]e[Ss]harper
*.DotSettings.user

# JustCode is a .NET coding addin-in
.JustCode

# TeamCity is a build add-in
_TeamCity*

# DotCover is a Code Coverage Tool
*.dotCover

# NCrunch
_NCrunch_*
.*crunch*.local.xml

# MightyMoose
*.mm.*
AutoTest.Net/

# Web workbench (sass)
.sass-cache/

# Installshield output folder
[Ee]xpress/

# DocProject is a documentation generator add-in
DocProject/buildhelp/
DocProject/Help/*.HxT
DocProject/Help/*.HxC
DocProject/Help/*.hhc
DocProject/Help/*.hhk
DocProject/Help/*.hhp
DocProject/Help/Html2
DocProject/Help/html

# Click-Once directory
publish/

# Publish Web Output
*.[Pp]ublish.xml
*.azurePubxml
# TODO: Comment the next line if you want to checkin your web deploy settings
# but database connection strings (with potential passwords) will be unencrypted
*.pubxml
*.publishproj

# NuGet Packages
*.nupkg
# The packages folder can be ignored because of Package Restore
**/packages/*
# except build/, which is used as an MSBuild target.
!**/packages/build/
# Uncomment if necessary however generally it will be regenerated when needed
#!**/packages/repositories.config

# Windows Azure Build Output
csx/
*.build.csdef

# Windows Store app package directory
AppPackages/

# Others
*.[Cc]ache
ClientBin/
~$*
*~
*.dbmdl
*.dbproj.schemaview
*.pfx
*.publishsettings
node_modules/
bower_components/

# RIA/Silverlight projects
Generated_Code/

# Backup & report files from converting an old project file
# to a newer Visual Studio version. Backup files are not needed,
# because we have git ;-)
_UpgradeReport_Files/
Backup*/
UpgradeLog*.XML
UpgradeLog*.htm

# SQL Server files
*.mdf
*.ldf

# Business Intelligence projects
*.rdl.data
*.bim.layout
*.bim_*.settings

# Microsoft Fakes
FakesAssemblies/

# Node.js Tools for Visual Studio
.ntvs_analysis.dat

# Visual Studio 6 build log
*.plg

# Visual Studio 6 workspace options file
*.opt

#Custom
project.lock.json
*.pyc
/.editorconfig

\.idea/

# Codealike UID
codealike.json`
}
