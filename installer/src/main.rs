use std::env;
use std::process::Command;
use owo_colors::OwoColorize;

pub fn install() {
    println!("{}", "Installing botway 🤖".bright_cyan());

    let os = env::consts::OS;
    let arch_cmd = Command::new("uname").arg("-m").output().unwrap();
    let arch = String::from_utf8(arch_cmd.stdout).unwrap();

    // remove last line from arch
    let arch = arch.trim_end();

    let releases_api_url="https://github.com/abdfnx/botway/releases/download";
    let latest_version_cmd = Command::new("curl").arg("--silent").arg("https://get-latest.deno.dev/abdfnx/botway").output().unwrap();
    let latest_version = String::from_utf8(latest_version_cmd.stdout).unwrap();

    let mut name = "botway_linux_".to_owned() + &latest_version + "_amd64";

    match os {
        "linux" => {
            match &arch as &str {
                "x86_64" => {
                    name = "botway_linux_".to_owned() + &latest_version + "_amd64"
                }

                "i686" => {
                    name = "botway_linux_".to_owned() + &latest_version + "_386"
                }

                "i386" => {
                    name = "botway_linux_".to_owned() + &latest_version + "_386"
                }

                "arm64" => {
                    name = "botway_linux_".to_owned() + &latest_version + "_arm64"
                }

                "arm" => {
                    name = "botway_linux_".to_owned() + &latest_version + "_arm"
                }

                _ => {
                    println!("{}", "Unsupported architecture".bright_red());
                },
            }
        }

        "macos" => {
            match &arch as &str {
                // TODO: replace macos with darwin
                "x86_64" => {
                    name = "botway_macos_".to_owned() + &latest_version + "_amd64"
                }

                "arm64" => {
                    name = "botway_macos_".to_owned() + &latest_version + "_arm64"
                }

                _ => {
                    println!("{}", "Unsupported architecture".bright_red());
                },
            }
        }

        "freebsd" => {
            match &arch as &str {
                "x86_64" => {
                    name = "botway_freebsd_".to_owned() + &latest_version + "_amd64"
                }

                "i686" => {
                    name = "botway_freebsd_".to_owned() + &latest_version + "_386"
                }

                "i386" => {
                    name = "botway_freebsd_".to_owned() + &latest_version + "_386"
                }

                "arm64" => {
                    name = "botway_freebsd_".to_owned() + &latest_version + "_arm64"
                }

                "arm" => {
                    name = "botway_freebsd_".to_owned() + &latest_version + "_arm"
                }

                _ => {
                    println!("{}", "Unsupported architecture".bright_red());
                },
            }
        }

        _ => {
            println!("{}", "Unsupported operating system".bright_red());
            print!("You can open an issue at {}", "https://github.com/abdfnx/botway/issues".bright_yellow());
            println!(" to support your operating system.");
        }
    }

    let botway_url = releases_api_url.to_owned() + &"/".to_owned() + &latest_version + "/" + &name + ".zip";

    Command::new("wget").arg(botway_url).output().unwrap();
    Command::new("sudo").arg("chmod").arg("+x").arg(name.to_owned() + ".zip").output().unwrap();
    Command::new("unzip").arg(name.to_owned() + ".zip").output().unwrap();
    Command::new("rm").arg(name.to_owned() + ".zip").output().unwrap();

    // move botway to /usr/bin
    Command::new("sudo").arg("mv").arg(name.to_owned() + "/bin/botway").arg("/usr/bin/botway").output().unwrap();

    // chmod /usr/bin/botway
    Command::new("sudo").arg("chmod").arg("+x").arg("/usr/bin/botway").output().unwrap();

    // clean up
    Command::new("rm").arg("-rf").arg(name.to_owned()).output().unwrap();

    // check if botway is installed
    let botway_installed = Command::new("which").arg("botway").output().unwrap();

    if botway_installed.stdout.is_empty() {
        println!("{}", "Botway is not installed".bright_red());
    } else {
        println!("{}", "Botway is installed successfully".bright_green());
        println!("{}", "🙏 Thanks for installing Botway! If this is your first time using the CLI, be sure to run `botway help` first.".bright_green());
        println!("{}", "Stuck? Join our Discord https://dub.sh/bw-discord".bright_cyan());
    }
}

fn main() {
    match env::consts::OS {
        "macos" | "linux" | "freebsd" => {
            install();
        }

        "windows" => {
            let _ = enable_ansi_support::enable_ansi_support();

            print!("This installer is only for unix oses (MacOS/Linux), run ");
            print!("{}", "`iwr -useb https://dub.sh/bw-win | iex`".bright_cyan());
            println!(" command for windows.");
        }

        &_ => {
            println!("{}", "Unsupported operating system".bright_red());
        }
    }
}
