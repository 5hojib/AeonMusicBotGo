# AeonMusicBotGo

A Telegram music search bot built with Go and deployed on Heroku.

[![Deploy on Heroku](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/yourusername/AeonMusisBotGo)

## Features

- Search music across multiple Telegram channels
- Inline query support
- User and bot client integration

## Prerequisites

- Telegram API ID and Hash
- Telegram bot token
- User session string

## Environment Variables

The following environment variables need to be set:

| Variable       | Description                     | Required |
|----------------|---------------------------------|----------|
| `APP_ID`       | Your Telegram API ID            | Yes      |
| `API_HASH`     | Your Telegram API Hash          | Yes      |
| `USER_SESSION` | Your Telegram user session      | Yes      |
| `BOT_TOKEN`    | Your Telegram bot token         | Yes      |

## Deployment

### 1. Deploy to Heroku (One-Click)

Click the "Deploy on Heroku" button above and fill in the required environment variables.

### 2. Manual Deployment with Heroku CLI

```bash
# Install Heroku CLI
# https://devcenter.heroku.com/articles/heroku-cli

# Login to Heroku
heroku login
heroku container:login

# Create a new Heroku app
heroku create your-app-name --stack=container

# Set environment variables
heroku config:set APP_ID=your_app_id \
                  API_HASH=your_api_hash \
                  USER_SESSION=your_user_session \
                  BOT_TOKEN=your_bot_token

# Deploy
git push heroku main

# Start the worker process
heroku ps:scale worker=1

# View logs
heroku logs --tail