<h1 align="center">
  <img src="https://github.com/enfipy/cronpub/blob/master/assets/logo.svg" width="300"/>
  <p align="center" style="font-size: 0.5em">Automated publishing bot</p>
</h1>

Project `cronpub` - automated publishing bot based on Telegram and Redis

## Usage:

To begin development:

1. Alter `settings.yaml` file `telegram` token and scraper links
2. Then just run:

```bash
docker-compose up --build
```

Or just fetch docker container from my [docker hub account](https://hub.docker.com/r/enfipy/cronpub)

## Todo:

1. Turn off/on scraper by altering `settings.yaml`
2. Add other social networks
