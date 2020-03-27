# Draco

[![Build Status](https://travis-ci.org/calini/draco.svg?branch=master)](https://travis-ci.org/calini/draco)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fcalini%2Fdraco.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fcalini%2Fdraco?ref=badge_shield)

A vulnerability management system for open source projects

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fcalini%2Fdraco.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fcalini%2Fdraco?ref=badge_large)

## Snyk Token
For the Snyk integration you will need to fetch an API token. In order to add it to the project, take the following steps:

1. Get an API Token for your organisation from [snyk.io](https://snyk.io)
2. Create a `.env` file at the root of the project
3. Paste in the SNYK_TOKEN as such: `SNYK_TOKEN=<SNYK_TOKEN here>`

## Building it
To build this project just run
```sh
make build
```

## Running it
To run this project directly 
```sh
make run
```