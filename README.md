
### Description

Bot's name:
- covidBot
- matthausen_covid_19_bot

You can find the bot at `t.me/matthausen_covid_19_bot` on Telegram.
Type `/help` to get instructions on how to use it.

The main command is `covid <COUNTRY_NAME>`, e.g. `covid china` or `covid uk`



### Deployment 

Expose your bot to the internet with the telegram setWebHook method

curl -F "url=https://covid19-telegram-bot.appspot.com" https://api.telegram.org/bot<TELEGRAM_TOKEN>/setWebhook
