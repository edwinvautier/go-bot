# go-bot

![Go](https://github.com/edwinvautier/go-bot/workflows/Go/badge.svg?branch=main)

A bot to interact with in natural language.

Made with :

- [go](https://github.com/golang/go)
- [dicordgo](https://github.com/bwmarrin/discordgo)
- [wit AI](https://wit.ai/)
- [youtube api v3](https://developers.google.com/youtube/v3)
- [OpenWeather API](https://openweathermap.org)

## Demo

![google search](./assets/google.gif)

![music search](./assets/marley.gif)

## Setup

First you need to created your `.env` file (you can use the .env.dist file).
The `DISCORD_TOKEN` variable is the discord bot token. It's needed in order to make the bot work.

You can get a token by following [this link](https://discord.com/developers/applications/)

We use wit.ai to understand natural language, a token is also needed as `WIT_TOKEN`

We also use youtube api to search for musics or videos, a token is needed as `YOUTUBE_TOKEN` and can be created in google's applications admin console.

Finally, to get the weather informations our API uses OpenWeather. [Get your API token](https://openweathermap.org/api) and save it in .env as `OWN_API_KEY`.

```sh
  docker-compose up --build
```

## Tests

You can run tests suites by using the following command :

```sh
  go test -v ./<package-name>
```

---

## Usage

Once the bot is invited to discord, you can talk to it simply by prefixing your messages with : `assistant,`.

The bot is able to :

- find musics, ex: `assistant, fais moi écouter du Mickael Jackson`
- give the weather, ex: `assistant, fait il beau à Paris`
- search on google, ex: `assistant, comment faire un noeud de chaise`

---

## Contributing

We follow a [code of conduct](CODE_OF_CONDUCT.md), if you wish to contribute on this project, we strongly advise you to read it.

### Branch naming convention

- You branch should have a name that reflects it's purpose.

- It should use the same guidelines as [COMMIT_CONVENTIONS](COMMIT_CONVENTIONS.md) (`feat`, `fix`, `build`, `perf`, `docs`), followed by an underscore (`_`) and a very quick summary of the subject in [kebab case][1].

    Example: `feat_add-image-tag-database-relation`.

### Pull requests (PR)

Pull requests in this project follow two conventions, you will need to use the templates available in the [ISSUE_TEMPLATE](.github/ISSUE_TEMPLATE) folder :

- Adding a new feature should use the [FEATURE_REQUEST](.github/ISSUE_TEMPLATE/FEATURE_REQUEST.md) template.
- Reporting a bug should use the [BUG_REPORT](.github/ISSUE_TEMPLATE/BUG_REPORT.md) template.

If your pull request is still work in progress, please add "WIP: " (Work In Progress) in front of the title, therefor you inform the maintainers that your work is not done, and we can't merge it.

The naming of the PR should follow the same rules as the [COMMIT_CONVENTIONS](COMMIT_CONVENTIONS.md)

### Continuous Integration (CI)

A CI pipeline is configured for this project and is accessible in the `.github/workflows/go-ci.yaml` file.

The pipeline will run 3 different jobs:

- Dependencies check
- Linter
- Tests

The pipeline will be triggered automatically when creating a new **Pull Request** and on each **push** on it. It will also be triggered on push on `main` branch.

---

## Contributors

<table align="center">
  <tr>
    <td align="center">
    <a href="https://github.com/jasongauvin">
      <img src="https://avatars1.githubusercontent.com/u/41618366?s=400&u=b970ed03cbb921ce1312ef86b39093e4fa0be7e3&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Jason Gauvin</b></sub>
    </a>
    </td>
    <td align="center">
    <a href="https://github.com/JackMaarek/">
      <img src="https://avatars3.githubusercontent.com/u/28316928?s=400&u=3cdfb5b0683245ad333a39cfca3a5251f3829824&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Jacques Maarek</b></sub>
    </a>
    </td>
    <td align="center">
    <a href="https://github.com/edwinvautier">
      <img src="https://avatars3.githubusercontent.com/u/35581502?s=460&u=d9096f90151f35552d9adcd57bacaee366f0aaef&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Edwin Vautier</b></sub>
    </a>
    </td>
  </tr>
</table>
