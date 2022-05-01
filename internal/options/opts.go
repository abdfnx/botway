package options

type RootOptions struct {
	Version bool
}

type InitOptions struct {
	Docker bool
}

type NewOptions struct {
	BotName string
}

type TokenAddOptions struct {
	BotName  string
	Discord  bool
	Slack    bool
	Telegram bool
}
