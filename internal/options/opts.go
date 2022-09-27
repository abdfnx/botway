package options

type RootOptions struct {
	Version bool
}

type InitOptions struct {
	CopyFile bool
	NoRepo   bool
}

type CommonOptions struct {
	BotName string
}

type NewOptions struct {
	NoRepo    bool
	RepoName  string
	IsPrivate bool
	IsBlank   bool
}

type TokenAddOptions struct {
	BotName  string
	Discord  bool
	Slack    bool
	Telegram bool
	Twitch   bool
}

type LoginOptions struct {
	Railway bool
	Render  bool
}
