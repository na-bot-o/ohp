[![CircleCI](https://circleci.com/gh/na-bot-o/ohp.svg?style=svg)](https://circleci.com/gh/na-bot-o/ohp)

# OHP

ohp is command to open recorded page(url) on browser

## Architecture

recorded page is managed in ${HOME}/.ohp file(default)
ohp command reads and writes to this file.

## Install

1. Download this repository
2. execute `make` command
   
※ you need to install go

## Usage

### list

open recorded page list

```sh
ohp list
```

### insert

record new page

```
ohp insert -n Google -t search -u https://google.com
```

※ can't allocate multiple tags or urls in a row

### open

open recorded url with matching tag or name flag on browser

```sh
ohp open -n Google
ohp open -t search
```


### delete

delete page with matching tag or name flag

```sh
ohp delete -n Google
ohp delete -t search
```