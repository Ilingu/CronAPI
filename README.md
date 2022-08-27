# CronAPI 🌐

#### ➡️ Simple Go Api that Handle Crons Execution

## Made With:

1. **Elegance** ✅
2. `Go` ⚡⚡ 🤍
3. `gorilla/mux` 🦍
4. `robfig/cron` 🔩

## ⚡ Usage

1.  Self Host this API in your own server, You'll Need:
    - Go (to build the api, or you can directly go to the Github's Releases)
    - `SERVER_KEY` **env var** prehashed (For Auth)
    - And/Or an `PORT` **env var** if your not using a serverless provider _(e.g: Heroku, Railway, AWS Thing for Serverless for which I absolutely didn't forgot the name...)_
2.  Send an `POST` Req at this route: `/addJob`:

    - ```json
      {
        /* BODY */
        "Frequency": "@every 1s", // Cron Frenquency
        "CallbackUrl": "http://localhost/receiver" // The Callback Url (that will be called as an webhook)
      }
      ```
    - ```json
        { "Authorization": SERVER_KEY } /* HEADER */
      ```
      \_PS: the `CallbackUrl` is the **Unique ID** of a CronJob since there can be only one `CallbackUrl` to call\*

3.  **Validate Each Webhooks Request**, each time the Cron Job is triggered, the _CronAPI_ will call your `CallbackUrl` with a `POST` request, the `CallbackUrl` must be an **valid** endpoint that should respond to the _CronAPI_ with the status code: `200 OK` and the Header: `Continue: true`. `Continue` means that the Cron should continue its execution. If by some reason these two requirements are not returned to the _CronAPI_ the Cron execution will **STOP** definitively.

4.  **[BONUS]**: If at some point you want to delete/stop a cron, you can send a `DELETE` req at `/delJob` route with this query: `?identifier=<CronID>`. _Note that `<CronID>` is the same as `CallbackUrl` of the CronJob targetted since `CallbackUrl` is the **Unique ID** of a CronJob, and by that fact it's an URL so you **MUST** encode it (e.g: in JS it'd be the function `encodeURIComponent()`_
