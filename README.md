
### Deep Fictitious

Ficticious chat GPT implementation.
Not sure where I’ve been reading the new buzzword “deep-fake” from, but I couldn’t appropriate it hence the improvisation.

This is for people who want to deploy, maintain and control their own version of a telegram bot built on openAI’s APIs. If you’d rather not go through the stress, use [Kainene](https://medium.com/r/?url=https%3A%2F%2Fsavant.holeyfox.co%2F), please. It’s miles better.

**A few things worthy(?) of note**

-   This uses in-memory(i.e. fake) storage to store conversation context. Little did I know that openAI has no time to handle that for you  :s . At the moment, it stores said context for about 12 minutes after the last message. 
-   It currently uses the `gpt-3.5-turbo`, but depending on your `rizz`, you might want to [change this](https://medium.com/r/?url=https%3A%2F%2Fgithub.com%2Fyouthtrouble%2Fcongenial-goggles%2Fblob%2F71dbf12594eaf71bff9e1d5b7d83ad17a92e77fc%2Fgpt%2FopenAI.go%23L58).

**How to run**

You need to copy the contents of the `.env.example` file into a new `.env` file.

-   `TELEGRAM_BOT_TOKEN` : You need to create a new bot on Telegram and generate new keys for the said bot. [Here’s](https://medium.com/r/?url=https%3A%2F%2Fsendpulse.ng%2Fknowledge-base%2Fchatbot%2Ftelegram%2Fcreate-telegram-chatbot) a good guide to help with this.
-   `OPENAI_API_KEY` : OpenAI provides API keys on the user settings on the web app. [Read more](https://medium.com/r/?url=https%3A%2F%2Fwww.windowscentral.com%2Fsoftware-apps%2Fhow-to-get-an-openai-api-key)…

Run `go run main.go` .

**Issues**

There’s currently a repeated occurrence of the bot not responding after some time. My suspicion is that it might not be a result of the code; rather, it’s associated with the free Render instance I currently have it hosted on. I might be wrong, and it’s totally open to fixes and contributions.

**Contributions**

Please contribute 🧎🏽

[Easter Egg](https://medium.com/r/?url=https%3A%2F%2Fgithub.com%2Fyouthtrouble%2Fcongenial-goggles%2Fblob%2F71dbf12594eaf71bff9e1d5b7d83ad17a92e77fc%2Fgpt%2FopenAI.go%23L40)