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
$ ohp list

 --------------------------------------------
 |  name  |    tag     |       url          |
 | google | search     | https://google.com |
 | amazon | e-commarce | https://amazon.com |
 --------------------------------------------

```

### insert

record new page

```
$ ohp insert -n yahoo -t search -u https://yahoo.com
$ ohp list

 --------------------------------------------
 |  name  |    tag     |       url          |
 | yahoo  | search     | https://yahoo.com  |
 | google | search     | https://google.com |
 | amazon | e-commarce | https://amazon.com |
 --------------------------------------------
```

※ can't allocate multiple tags or urls in a row

### open

open recorded url with matching tag or name flag on browser

```sh

$ ohp list

 --------------------------------------------
 |  name  |    tag     |       url          |
 | yahoo  | search     | https://yahoo.com  |
 | google | search     | https://google.com |
 | amazon | e-commarce | https://amazon.com |
 --------------------------------------------

$ ohp open -n Google
-> open browser and open "https://google.com"
$ ohp open -t search
-> open browser and open "https://google.com" and "https://yahoo.com"
```


### delete

delete page with matching tag or name flag

```sh
$ ohp list
 --------------------------------------------
 |  name  |    tag     |       url          |
 | yahoo  | search     | https://yahoo.com  |
 | google | search     | https://google.com |
 | amazon | e-commarce | https://amazon.com |
 --------------------------------------------

$ ohp delete -t search
$ ohp list

 --------------------------------------------
 |  name  |    tag     |       url          |
 | amazon | e-commarce | https://amazon.com |
 --------------------------------------------

$ ohp delete -n amazon

$ ohp list
-> none

```
