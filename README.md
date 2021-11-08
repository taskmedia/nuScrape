# nuScrape

[![releases](https://img.shields.io/github/v/release/taskmedia/nuScrape?style=flat-square)](https://github.com/taskmedia/nuScrape/releases/latest)
[![docs](https://img.shields.io/badge/docs-pkg.go.dev-blue?style=flat-square)](https://pkg.go.dev/github.com/taskmedia/nuScrape)
[![golang version](https://img.shields.io/github/go-mod/go-version/taskmedia/nuScrape?style=flat-square)](https://golang.org/dl/#stable)
<br />
[![codecoverage](https://img.shields.io/codecov/c/github/taskmedia/nuScrape?style=flat-square)](https://app.codecov.io/gh/taskmedia/nuScrape)
![code size](https://img.shields.io/github/languages/code-size/taskmedia/nuScrape?style=flat-square)
<br />
[![issues](https://img.shields.io/github/issues/taskmedia/nuScrape?style=flat-square)](https://github.com/taskmedia/nuScrape/issues)
[![pull requests](https://img.shields.io/github/issues-pr/taskmedia/nuScrape?style=flat-square)](https://github.com/taskmedia/nuScrape/pulls)
<br />
[![twitter](https://img.shields.io/twitter/follow/taskmediaDE?style=social)](https://twitter.com/taskmediaDE)

This application will generate a REST endpoint to fetch data from [nuLiga](https://bhv-handball.liga.nu/).
The service will parse the table of nuLiga to a JSON object.

## Start application

You will be able to start the application directly with golang:

```bash
go run cmd/nuScrape/nuScrape.go
```

Another option would be running the application in a Docker container:

```bash
docker run \
  --name nuscrape \
  -p 8080:8080 \
  taskmedia/nuscrape:latest
```

## Example API call

To get a gesamtspielplan you can request following endpoint: http://localhost:8080/rest/v1/gesamtspielplan/2021_22/AV/281103

The response will be e.g.:

```json
[
  {
      "date": "2021-10-09T19:15:00Z",
      "team": {
        "home": "TV Memmingen",
        "guest": "TSV Ottobeuren III"
      },
      "goal": {
        "home": 28,
        "guest": 27
      },
      "location": 27031301,
      "id": 27031301,
      "annotation": {
        "date": "",
        "result": ""
      },
      "report": 7013920,
      "referee": null
  },
  ...
]
```
