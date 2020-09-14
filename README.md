# Processor

What an original name, right?

## Structure

- `cmd/processor` : http handlers that will execute the request on the backend
- `cmd/alan` : cli tool to interact with the application (just to avoid using `curl` or `httpie`)
- `cmd/bot`* : discord bot to interact with the app. Simmilar to the cli, but thru discord.

