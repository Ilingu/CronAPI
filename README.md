# CronAPI 🌐

#### ➡️ Simple Go Api that Handle Crons Execution

## Made With:

1. **Elegance** ✅
2. `Go` ⚡⚡ 🤍
3. `gorilla/mux` 🦍
4. `robfig/cron` 🔩

## ⚡ Usage

1. Self Host this API in your own server, You'll Need:
   - Go (to build the api, or you can directly go to the Github's Releases)
   - `SERVER_KEY` **env var** prehashed (For Auth)
   - And/Or an `PORT` **env var** if your not using a serverless provider _(e.g: Heroku, Railway, AWS Thing for Serverless for which I absolutely didn't forgot the name...)_
2. Sand an POST Req at `https://whatever.whatever/addJob` with

```json
{
  "Frequency": "@every 1s", // Cron Frenquency
  "CallbackUrl": "http://localhost/receiver" // The Callback Url (that will be called as an webhook)
}
```

3. **Validate Each Webhooks Request**, the _CronAPI_ will call your `CallbackUrl` with a `POST` request, the `CallbackUrl` must be an **valid** endpoint that should respond to the _CronAPI_ with the status code: `200` and the Header: `Continue: true`. `Continue` means that the Cron should continue its execution, if by some reason this is not returned to the _CronAPI_ the Cron execution will **STOP** definitively.
