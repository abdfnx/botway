<div align="center">
  <h1>BW++</h1>
	<p>
		C++ client package for Botway
	</p>
</div>

## Usage

> after creating a new c++ botway project, you need to use your tokens to connect with your bot.

```c++
#include <dpp/dpp.h>
#include <botname/botname.h>
#include <sstream>
#include "botway/botway.h"

using namespace std;

int main(int argc, char const *argv[]) {
    dpp::cluster bot(Get("botname", "token"));
...
```
