[![CircleCI](https://circleci.com/gh/na-bot-o/ohp.svg?style=svg)](https://circleci.com/gh/na-bot-o/ohp)

# OHP

ohp is command to open recorded page(url) on browser

## Architecture

recorded page is managed in ${HOME}/.ohp file(default)
ohp command reads and writes to this file.

## Install

1. Download this repository
2. execute `make` command
   
※ you need to set up go env in advance

## Usage

### list

open recorded page list

```sh
ohp list

//result image
// ---------------------------------------
// |  name  |  tag   |       url         |
// | google | search | https://google.com|
// ---------------------------------------

```

### insert

record new page

```
//flag 
//  -n page name
//  -t tag name
//  -u url
ohp insert -n Google -t search -u https://google.com
```

※ can't allocate multiple tags or urls in a row

### open

open recorded url with matching tag or name flag on browser

```sh
//flag
//  -n page name
//  -t tag name
ohp open -n Google
ohp open -t search
```


### delete

delete page with matching tag or name flag

```sh
//flag
//  -n page name
//  -t tag name
ohp delete -n Google
ohp delete -t search
```