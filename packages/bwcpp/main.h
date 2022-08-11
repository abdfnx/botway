#pragma once
#include <{{.BotName}}/{{.BotName}}.h>
#include <sstream>

using namespace std;
using json = nlohmann::json;

json config;
string homeDir = getenv("HOME");
string slash = "/";

#ifdef _WIN32
    homeDir = getenv("HOMEPATH");
    slash = "\\";
#endif

ifstream configfile(homeDir + slash + ".botway" + slash + "botway.json");

string Get(string botName, string value) {
    configfile >> config;

    if (value.find("token") != string::npos) {
        return config["botway"]["bots"][botName]["bot_token"];
    } else if (value.find("id") != string::npos) {
        return config["botway"]["bots"][botName]["bot_app_id"]; 
    }

    return config["botway"]["bots"][botName][value];
}

string GetGuildId(string botName, string serverName) {
    configfile >> config;

    return config["botway"]["bots"][botName]["guilds"][serverName]["server_id"];
}
