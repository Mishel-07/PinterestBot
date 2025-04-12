# PinterestBot

Welcome to **PinterestBot** your friendly and efficient Telegram bot with exciting features!

This bot can:
- Download images from Pinterest links
- Perform Pinterest searches
- Fetch wallpapers
- Access Google and Bing images

---

## Support

- **Try it**: [Sample Bot](https://t.me/ImgRobot)
- **Support**: [Join Group](https://t.me/XBOTSUPPORTS)

---

## Commands 

```
/pinterest <query> - Search and download Pinterest images
/wallpaper <query> - Get a wallpaper from wallpaper.com
/img <query>       - Search images using Bing
```

You can also use these features via **inline mode**:
- Type `@YourBotUsername <query>` in any chat

---

## Features

- Download images from direct Pinterest links
- Search Pinterest and return top image results
- Access wallpapers from wallpaper.com
- Bing image search support
- Full inline mode support for image search

---

## Environment Variables

Set the following in your environment:

```
TOKEN   - Your bot token from BotFather
PORT    - Optional: defaults to 8080
WEBHOOK - Optional: set to "false" to disable
URL     - Optional: If you're hosting on a Koyeb or Render-like server and the bot occasionally stops, you can use this to set the Render/Koyeb URL. This is specifically for web-only support servers. If there are no requests for 2 minutes, the server may stop. You can enable WEBHOOK to prevent this behavior.
```
