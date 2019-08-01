## Issue Finder

Do you want to contribute to open source projects but can't seem to find fresh issues?
Use this tool to search github for issues with specific labels in a specific language

#### Installation
```bash
go get github.com/phillipahereza/issue_finder
```



#### Usage

[![asciicast](https://asciinema.org/a/HwURcOP2S6G0g8QKcxl7XZD9n.svg)](https://asciinema.org/a/HwURcOP2S6G0g8QKcxl7XZD9n)

```
$ issues help

NAME:
   issues - Need to contribute to OpenSource? Find fresh issues

USAGE:
   issues [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --language value, -l value  search for issues in this language
   --label value, -b value     Search for issues with this label (default: "good-first-issue")
   --days value, -d value      Search for issues created within the last n days (default: 30)
   --help, -h                  show help
   --version, -v               print the version
   
$ issues -l go -b help-wanted -d 10

```

