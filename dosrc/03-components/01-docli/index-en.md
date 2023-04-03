`DOCLI` is a command line tool to pre-process website content. Its features include:
- Check website content for errors.
- Build fulltext index.

## Installation

Firstly, <a href="https://go.dev/doc/install" target="_blank">download</a> and install Go. Then install `DOCLI` with the following command:
```shell
go install github.com/btnguyen2k/docms/docli@latest
```

> `DOCLI` is installed to directory `$GOPATH/bin`. You may need to add the directory to system's PATH.

## Usage

**Get help**

Command `docli help` shows list of available commands, and `docli help command` shows help for a specific command. For example:

```shell
$ docli help
NAME:
   docli - Pre-process and build content for DO CMS

USAGE:
   docli [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   build, b  Build DO CMS data
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

**Build website content**

Website content is pre-processed with command `docli build`. For example:

```shell
docli build --src dosrc --out dodata
```

Command's options:

|Option|Required|Defaut value|Description|
|---|---|---|---|
|`--src source_dir`|false|`dosrc`|source directory where website content is stored|
|`--out output_dir`|false|`dodata`|output directory where pre-processed data is written to|
|`--purge`|false||purge output directory before procesing|
