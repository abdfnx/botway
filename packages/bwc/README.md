<div align="center">
  <h1>BWC</h1>
	<p>
		C client package for Botway
	</p>
</div>

## Usage

> after creating a new c botway project, you need to use your tokens to connect with your bot.

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "botway.h"
#include "concord/discord.h"

...

int main(int argc, char *argv[])
{
    const char *botway_dir = path_join(getenv("HOME"), ".botway");
    const char *config_file = path_join(botway_dir, "botway-c-config.json");

    ccord_global_init();
    struct discord *client = discord_config_init(config_file);

    discord_set_on_ready(client, &on_ready);
    discord_set_on_command(client, "ping", &on_ping);
    discord_set_on_command(client, "pong", &on_pong);

    print_usage();

    discord_run(client);

    discord_cleanup(client);
    ccord_global_cleanup();
}
```
