#[cfg(botway_rs)]
#[allow(clippy::module_inception)]
pub mod botway_rs;

extern crate dirs;
use rustc_serialize::json::Json;
use snailquote::unescape;
use std::fs;
use std::fs::File;
use std::io::Read;
use std::path::Path;
use yaml_rust::YamlLoader;

fn get_home_dir() -> String {
    let home_dir = dirs::home_dir().unwrap();

    home_dir.to_str().unwrap().to_string()
}

fn return_path() -> String {
    let mut home_dir = get_home_dir();

    home_dir.push_str("/.botway/");
    home_dir.push_str("/botway.json");

    let path = Path::new(&home_dir);

    let mut file = File::open(path).unwrap();
    let mut data = String::new();

    file.read_to_string(&mut data).unwrap();

    format!("{}", data)
}

fn get_bot_info(value: &str) -> String {
    let bot_config =
        fs::read_to_string(".botway.yaml").expect("ERROR: Botway config file not found");

    let data = YamlLoader::load_from_str(&bot_config).unwrap();

    data[0]["bot"][value].as_str().unwrap().to_string()
}

// Get bot secrets
// Available bot secrets: token, app_id
//
// # Example
//
// ```
// let token = botway::get("token");
// let app_id = botway::get("app_id");
// ```
pub fn get(value_to_get: &str) -> String {
    let json = Json::from_str(&return_path()).unwrap();

    if value_to_get == "token" {
        unescape(
            &json
                .find_path(&["botway", "bots", &get_bot_info("name"), "bot_token"])
                .unwrap()
                .to_string(),
        )
        .unwrap()
    } else if value_to_get == "app_id" {
        unescape(
            &json
                .find_path(&["botway", "bots", &get_bot_info("name"), "bot_app_id"])
                .unwrap()
                .to_string(),
        )
        .unwrap()
    } else {
        "ERROR: Invalid value to get".to_string()
    }
}

// Get bot guild ids
// This function is only working with discord bots
//
// # Example
//
// ```
// let my_server_id = botway::get_guild_id("SERVER_NAME");
// ```
pub fn get_guild_id(server_name: &str) -> String {
    if get_bot_info("type") != "discord" {
        "ERROR: This function/feature is only working with discord bots.".to_string()
    } else {
        let json = Json::from_str(&return_path()).unwrap();

        unescape(
            &json
                .find_path(&[
                    "botway",
                    "bots",
                    &get_bot_info("name"),
                    "guilds",
                    server_name,
                    "server_id",
                ])
                .unwrap()
                .to_string(),
        )
        .unwrap()
    }
}
