# findyshark

A recursive-searching CLI tool for Linux

## Table of Contents

- [Installation](#installation)
- [Config](#config)
- [Usage](#usage)

## Installation

### Install Dependencies

#### Install Go

e.g. after verifying Go isn't already installed:

``` bash
wget https://golang.org/dl/go1.17.3.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.17.3.linux-amd64.tar.gz
```

Add Go env variables to `~/.bash_profile`

e.g. in `~/.bash_profile`:

```  bash
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

``` bash
sudo source ~/.bash_profile
```

Verify Go is installed; e.g.

``` bash
$ go version
go version go1.17.3 linux/amd64
```

#### Install Other Dependencies

Install `make`; e.g.
``` bash
sudo yum install make
```

Install `goxc`; e.g.
``` bash
sudo go get -v github.com/laher/goxc
```

Install `shc`; e.g.
``` bash
sudo yum install http://repo.okay.com.mx/centos/7/x86_64/release/okay-release-1-1.noarch.rpm
sudo yum install shc
```

Install GCC; e.g.
``` bash
sudo yum group install "Development Tools"
sudo yum install man-pages
```

### Install `findyshark`

Clone this repo; e.g.
``` bash
git clone https://github.com/thaigoonch/findyshark.git
 ```
 
Build and install

```
make install
```

Verify `findyshark` is installed; e.g.
``` bash
$ findyshark version
findyshark '0.0.2'
```

Note: Do not move the installed binary files out of the `'$(GOPATH)/bin/` location, or the program will not work. If you need to put them elsewhere, make sym links.

## Config

Optionally, create a config file at `$HOME/.findyshark.yaml`

Set the `ignore` variable to a value or a comma-separated list of values. These should specify files and directories you want findyshark to exclude in

e.g. to ignore all `vendor`, `.git` directories and `go.sum`, `go.mod` files, in `$HOME/.findyshark.yaml`:

``` bash
ignore: .*/vendor/*,*/.git/*,*/go.sum,*/go.mod
```

Note: do not put a `*` as the first character. The CLI will not read in the ignores if you do. If you need to do so, put a `.` before it like in the example above.

## Usage

Runs out of pwd.

### To run:
``` bash
  findyshark [flags]
  findyshark [command]
```

#### Available Commands:

  `completion`  generate the autocompletion script for the specified shell
  
  `help`        Help about any command
  
  `version`     Show findyshark version
  

#### Flags:

`--config` string      path to config file (default is $HOME/.findyshark.yaml)

`-e`, `--extension` string   search in specified file extension; e.g. `txt`

`-h`, `--help`               help for findyshark

`-i`, `--insensitive`        search case-insensitive


### Example output:
```
        _________         .    .
       (..       \_    ,  |\  /|    +-+-+-+-+-+-+            dooo
        \       0  \  /|  \ \/ /    |findy|shark|                    doo doo
         \______    \/ |   \  /     +-+-+-+-+-+-+      doo                        da-doo
           vvvvv       |  /   |
            ^^^^       \_/    )     Using config file: /home/it3510/.findyshark.yaml
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

