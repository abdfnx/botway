const { App, LogLevel } = require("@slack/bolt");
const botway = require("botway.js");

const app = new App({
  socketMode: true,
  token: botway.GetToken(),
  signingSecret: botway.GetSigningSecret(),
  appToken: botway.GetAppId(),
  logLevel: LogLevel.DEBUG,
});

app.command("/hello", async ({ body, ack, say }) => {
  await ack();

  await say(`Hi ${body.user_name}`);
});

app.event("app_mention", async ({ event, say }) => {
  await say("What's up?");
});

(async () => {
  await app.start(process.env.PORT || 8080);

  console.log("⚡️ Bolt app is running!");
})();
