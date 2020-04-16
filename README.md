
[![Release](https://img.shields.io/github/v/release/clementlecorre/mail-to-telegram)](https://github.com/clementlecorre/mail-to-telegram/releases/latest)
[![Go version](https://img.shields.io/github/go-mod/go-version/clementlecorre/mail-to-telegram/master)](https://golang.org/doc/devel/release.html)
[![Docker](https://img.shields.io/docker/pulls/cl3m3nt/mail-to-telegram)](https://hub.docker.com/r/cl3m3nt/mail-to-telegram)

## About

`mail-to-telegram` listens to your mail (imap) and sends the message on telegram. The email fetch is event-based. 
You can customize the processing on the received email 

## Install

### Docker

```bash
docker pull clementlecorre/mail-to-telegram
docker container run --rm clementlecorre/mail-to-telegram --help
```
