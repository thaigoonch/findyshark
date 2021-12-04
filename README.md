# findyshark

A recursive-searching CLI tool.

## Table of Contents

- [Installation](##Installation)
- [Config](##Config)
- [Usage](##Usage)

## Installation

TBD

## Config

Create a config file at `$HOME/.findyshark.yaml`

Set the `ignore` variable to a value or a comma-separated list of values. These should specify files and directories you want findyshark to exclude in searches. 

e.g. to ignore all `vendor`, `.git` directories and `go.sum`, `go.mod` files:

``` bash
ignore: .*/vendor/*,*/.git/*,*/go.sum,*/go.mod
```

Note: do not put a star as the first character. The CLI will not read in the ignores if you do. If you need to do so, put a `.` before it like in the example above.

## Usage

Runs out of pwd.

To run:
``` bash
findyshark [flags]
```

Flags:
``` bash
      --config string      path to config file (default is $HOME/.findyshark.yaml)
  -e, --extension string   search in specified file extension; e.g. txt
  -h, --help               help for findyshark
  -i, --insensitive        search case-insensitive
```

Example output:
``` 
        _________         .    .
       (..       \_    ,  |\  /|    +-+-+-+-+-+-+            dooo
        \       0  \  /|  \ \/ /    |findy|shark|                    doo doo
         \______    \/ |   \  /     +-+-+-+-+-+-+      doo                        da-doo
           vvvvv       |  /   | 
            ^^^^       \_/    )     Using config file: /home/katherine-nieman/.findyshark.yaml
             `\_   )))       /
             / /\_   \ /    /       Ignoring: .*/vendor/*,*/.git/*,*/go.sum,*/go.mod
             |/   \___\|___/
                       v   

find: 
github
╭─────────────────────────────┬─────┬───────────────────────────────────────────────╮
│ FILENAME                    │ LN# │ RESULT                                        │
├─────────────────────────────┼─────┼───────────────────────────────────────────────┤
│ ./main.go                   │ 26  │ // from github                                │
├─────────────────────────────┼─────┼───────────────────────────────────────────────┤
│ ./my_vendor/myTermstxt      │ 1   │ find, file, vendor, shark, github             │
├─────────────────────────────┼─────┼───────────────────────────────────────────────┤
│ ./app/table.go              │ 8   │     "github.com/jedib0t/go-pretty/v6/table"   │
├─────────────────────────────┼─────┼───────────────────────────────────────────────┤
│ ./app/table.go              │ 9   │     "github.com/jedib0t/go-pretty/v6/text"    │
├─────────────────────────────┼─────┼───────────────────────────────────────────────┤
│ ./cmd/root.go               │ 24  │     homedir "github.com/mitchellh/go-homedir" │
├─────────────────────────────┼─────┼───────────────────────────────────────────────┤
│ ./cmd/root.go               │ 25  │     "github.com/spf13/cobra"                  │
├─────────────────────────────┼─────┼───────────────────────────────────────────────┤
│ ./cmd/root.go               │ 26  │     "github.com/spf13/viper"                  │
├─────────────────────────────┼─────┼───────────────────────────────────────────────┤
│ FOUND 7 RESULTS IN 4 FILES. │     │                                               │
╰─────────────────────────────┴─────┴───────────────────────────────────────────────╯
```
