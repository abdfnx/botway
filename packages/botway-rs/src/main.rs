pub mod botway {
    use std::fs;
    use yaml_rust::{YamlLoader};
    use rustc_serialize::json::Json;
    use std::fs::File;
    use std::io::Read;
    use snailquote::unescape;

    fn get_bot_info(value: &str) -> String {
        let bot_config = fs::read_to_string(".botway.yaml")
            .expect("ERROR: Botway config file not found");

        let data = YamlLoader::load_from_str(&bot_config).unwrap();

        data[0]["bot"][value].as_str().unwrap().to_string()
    }

    pub fn get(value_to_get: &str) -> String {
        if get_bot_info("lang") != "rust" {
            "ERROR: Botway is not configured for rust".to_string()
        } else {
            let mut file = File::open("botway.json").unwrap();
            let mut data = String::new();
            file.read_to_string(&mut data).unwrap();
            let json = Json::from_str(&data).unwrap();

            if value_to_get == "token" {
                unescape(&json.find_path(&["botway", "bots", &get_bot_info("name"), "bot_token"]).unwrap().to_string()).unwrap()
            } else if value_to_get == "app_id" {
                unescape(&json.find_path(&["botway", "bots", &get_bot_info("name"), "bot_app_id"]).unwrap().to_string()).unwrap()
            } else {
                "ERROR: Invalid value to get".to_string()
            }
        }
    }

    pub fn get_guild_id(server_name: &str) -> String {
        if get_bot_info("lang") != "rust" {
            "ERROR: Botway is not configured for rust".to_string()
        } else if get_bot_info("type") != "discord" {
            "ERROR: This function/feature is only working with discord bots.".to_string()
        } else {
            let mut file = File::open("botway.json").unwrap();
            let mut data = String::new();
            file.read_to_string(&mut data).unwrap();
            let json = Json::from_str(&data).unwrap();

            unescape(&json.find_path(&["botway", "bots", &get_bot_info("name"), "guilds", server_name, "server_id"]).unwrap().to_string()).unwrap()
        }
    }
}

fn main() {
    println!("{}", botway::get("token"));
    println!("{}", botway::get_guild_id("secman"));
}
