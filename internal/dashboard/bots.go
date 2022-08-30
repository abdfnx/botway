package dashboard

import (
	"github.com/abdfnx/botway/constants"
	"github.com/tidwall/gjson"
)

var (
	bots       = gjson.Get(string(constants.BotwayConfig), "botway.bots_names")
	bots_count = gjson.Get(string(constants.BotwayConfig), "botway.bots_names.#").Int()
)

func (b Bubble) botInfo(infoToGet string) string {
	v := ""

	bots.ForEach(func(i, value gjson.Result) bool {
		if b.bubbles.primaryPaginator.Cursor == int(i.Int()) && !b.nav.listCursorHide {
			v = value.String()
		}

		return true
	})

	bot := gjson.Get(string(constants.BotwayConfig), "botway.bots."+v)

	if infoToGet == "token" {
		return gjson.Get(bot.String(), "bot_token").String()
	} else if infoToGet == "name" {
		return v
	}

	return gjson.Get(bot.String(), infoToGet).String()
}
