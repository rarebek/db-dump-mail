# PostgreSQL Backup Tool

This tool creates PostgreSQL database backups and can send them via email and/or Telegram bot.

## Features

- Creates PostgreSQL database dumps using pg_dump
- Sends backups via email (Gmail SMTP)
- Sends backups via Telegram bot (optional)
- Supports both delivery methods simultaneously
- Configurable via environment variables

## Setup

### Telegram Bot Setup (Optional)

1. Create a new bot by messaging [@BotFather](https://t.me/botfather) on Telegram
2. Send `/newbot` and follow the instructions
3. Copy the bot token provided
4. Get your chat ID by:
   - Messaging your bot
   - Visiting `https://api.telegram.org/bot<YOUR_BOT_TOKEN>/getUpdates`
   - Look for the "chat":{"id": value in the response

## Installation

1. Get dependencies:
```bash
go mod tidy
```

2. Build and place binary:
```bash
go build -o pgbackup
sudo cp pgbackup /usr/local/bin/
```

3. Configure environment:
```bash
# Copy the example file and edit it
cp .env.example .env
# Edit .env with your database and notification settings
```

4. Place .env file:
```bash
sudo mkdir /etc/pgbackup
sudo cp .env /etc/pgbackup/
```

5. Open crontab:
```bash
crontab -e
```

6. Add one of these lines:
```bash
# Daily at 2 AM (cd is needed because .env file is in /etc/pgbackup)
0 2 * * * cd /etc/pgbackup && /usr/local/bin/pgbackup

# Every 6 hours
# 0 */6 * * * cd /etc/pgbackup && /usr/local/bin/pgbackup

# Weekly (Sunday 1 AM)
# 0 1 * * 0 cd /etc/pgbackup && /usr/local/bin/pgbackup
```
