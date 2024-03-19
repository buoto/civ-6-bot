civ-6-bot
========

Webhook server for play by cloud notifications hosted on Vercel.

The webhook sends JSONs in format ([source](https://www.reddit.com/r/civ/comments/aqwwui/play_by_cloud_making_webhook_work_with_discord/)):
```json
{ "value1":"[game name]", "value2":"[player name]", "value3":"[game turn number]" }
```
