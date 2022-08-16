package options

type RootOptions struct {
	Version bool
}

type InitOptions struct {
	Docker bool
}

type CommonOptions struct {
	BotName string
}

type NewOptions struct {
	CreateRepo     bool
	RepoName       string
	IsPrivate      bool
	IsBlank        bool
}

type TokenAddOptions struct {
	BotName  string
	Discord  bool
	Slack    bool
	Telegram bool
}

type LoginOptions struct {
	Railway bool
	Render  bool
}
