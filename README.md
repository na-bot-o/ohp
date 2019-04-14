[![CircleCI](https://circleci.com/gh/na-bot-o/ohp.svg?style=svg)](https://circleci.com/gh/na-bot-o/ohp)

# OHP

ohp is command to open recorded page(url) on browser

## Architecture

recorded page is managed in ${HOME}/.ohp file(default)
ohp command reads and writes to this file.

## Usage

### list

open recorded page list

### insert

record new page

### open

open recorded url with matching tag or name flag on browser


### delete

delete page list with matching tag or name flag

### update

update page list info (name, tag, or url)