package discord

import (
	"log"
	"os"
	"os/exec"

	"github.com/abdfnx/looker"
)

func LinuxPackageManagerPython() string {
	return `
		pkgs="libffi-dev python-dev"
		declare -A osInfo;

		osInfo[/etc/redhat-release]=yum
		osInfo[/etc/arch-release]=pacman
		osInfo[/etc/gentoo-release]=emerge
		osInfo[/etc/debian_version]=apt-get
		osInfo[/etc/alpine-release]=apk

		for f in ${!osInfo[@]}
		do
			if [[ -f $f ]];then
				if [ "${osInfo[$f]}" == "yum" ]; then
					sudo yum install -y $pkgs
				elif [ "${osInfo[$f]}" == "pacman" ]; then
					sudo pacman -S $pkgs
				elif [ "${osInfo[$f]}" == "emerge" ]; then
					sudo emerge -pv $pkgs
				elif [ "${osInfo[$f]}" == "apt-get" ]; then
					sudo apt-get install -y $pkgs
				elif [ "${osInfo[$f]}" == "apk" ]; then
					sudo apk add $pkgs
				fi
			fi
		done
	`
}

func InstallCommandPython() {
	cmd := exec.Command("bash", "-c", LinuxPackageManagerPython())

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

func LinuxPackageManagerRust() string {
	return `
		declare -A osInfo;

		osInfo[/etc/redhat-release]=yum
		osInfo[/etc/arch-release]=pacman
		osInfo[/etc/gentoo-release]=emerge
		osInfo[/etc/debian_version]=apt-get
		osInfo[/etc/alpine-release]=apk

		for f in ${!osInfo[@]}
		do
			if [[ -f $f ]];then
				if [ "${osInfo[$f]}" == "yum" ]; then
					sudo yum install -y opus ffmpeg
				elif [ "${osInfo[$f]}" == "pacman" ]; then
					sudo pacman -S opus base-devel youtube-dl
				elif [ "${osInfo[$f]}" == "emerge" ]; then
					sudo emerge -pv opus ffmpeg
				elif [ "${osInfo[$f]}" == "apt-get" ]; then
					sudo apt-get install -y libopus-dev ffmpeg build-essential autoconf automake libtool m4 youtube-dl
				elif [ "${osInfo[$f]}" == "apk" ]; then
					sudo apk add libsodium ffmpeg opus autoconf automake libtool m4 youtube-dl
				fi
			fi
		done
	`
}

func InstallCommandRust() {
	cmd := exec.Command("bash", "-c", LinuxPackageManagerRust())

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

func InstallCommandRuby() {
	_, err := looker.LookPath("brew")

	if err != nil {
		log.Fatal("error: brew is not installed")
	} else {
		cmd := exec.Command("bash", "-c", "brew install opus libsodium")

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}
}
