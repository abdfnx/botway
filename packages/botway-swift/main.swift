import Foundation

struct BW: Codable {
    let botway: Botway
    let user: User
}

struct Botway: Codable {
    let bots: Bots
    let botsNames: [String]

    enum CodingKeys: String, CodingKey {
        case bots
        case botsNames = "bots_names"
    }
}

struct Bots: Codable {
    let {{.BotName}}: {{.BotName}}
}

struct {{.BotName}}: Codable {
    let type, path, lang, botToken: String
    let botAppID: String?

    enum CodingKeys: String, CodingKey {
        case type, path, lang
        case botToken = "bot_token"
        case botAppID = "bot_app_id"
    }
}

struct User: Codable {
    let dockerID, githubUsername: String

    enum CodingKeys: String, CodingKey {
        case dockerID = "docker_id"
        case githubUsername = "github_username"
    }
}

let botwayConfigPath = NSString(string: NSHomeDirectory())
    .appendingPathComponent(NSString(string: ".botway")
    .appendingPathComponent(NSString(string: "botway.json")
    as String
))

let config = try? Data(contentsOf: URL(fileURLWithPath: botwayConfigPath))

let botwayConfig = try? JSONDecoder().decode(BW.self, from: config!)

public func GetToken() -> String {
    return botwayConfig!.botway.bots.{{.BotName}}.botToken
}

public func GetAppId() -> String {
    return botwayConfig!.botway.bots.{{.BotName}}.botAppID!
}
