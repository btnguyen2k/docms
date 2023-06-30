`DOCLI` is a command line tool to pre-process website content. Its features include:
- Check content metadata for errors.
- Build fulltext index.

## Installation

Firstly, [download](https://go.dev/doc/install) and install Go. Then install `DOCLI` with the following command:
```shell
go install github.com/btnguyen2k/docms/docli@latest
```

You can also install a specific version (*):
```shell
go install github.com/btnguyen2k/docms/docli@cli-v0.3.1.3
```

```bs-alert warning

`DOCLI` is installed to directory `$GOPATH/bin`. You may need to add the directory to system's PATH.

(*) List of available versions can be found [here](https://github.com/btnguyen2k/docms/tags).
```

## Usage

### Get help

Command `docli help` shows list of available commands, and `docli help command` shows help for a specific command. For example:

```shell
$ docli help
NAME:
   docli - DO CMS website content preprocessing tool

USAGE:
   docli [global options] command [command options] [arguments...]

VERSION:
   0.3.1.1

AUTHOR:
   Thanh Nguyen <btnguyen2k (at) gmail (dot) com>

COMMANDS:
   build, b           Preprocess website content, ready for DO CMS runtime
   new, n, create, c  Helper to create assets with default metadata
   touch, t           Touch document metadata file to update timestamp
   help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

COPYRIGHT:
   Copyright (c) 2023 - DO CMS
```

### Pre-process website content

Website content is pre-processed with command `docli build`. For example:

```shell
docli build --src dosrc --out dodata
```

`docli help build` shows help of `build` command:

```shell
$ docli help build
NAME:
   docli build - Preprocess website content, ready for DO CMS runtime

USAGE:
   docli build [command options] [arguments...]

OPTIONS:
   --src value  source directory
   --out value  output directory
   --purge      purge output directory if not empty (default: false)
   --help, -h   show help
```

Command's options:

|Option|Required|Defaut value|Description|
|---|---|---|---|
|`--src value`|false|`dosrc`|source directory where website content is stored|
|`--out value`|false|`dodata`|output directory where pre-processed data is written to|
|`--purge`|false||purge output directory before procesing|

### Helper to create content metadata

Starting from version `cli-v0.1.1`, `DO CLI` provides a few useful tools to help creating website content metadata via command `docli new`:

- `docli new [global options] site     [command options] [arguments...]`: create metadata for site content
- `docli new [global options] topic    [command options] [arguments...]`: create metadata for topic
- `docli new [global options] document [command options] [arguments...]`: create metadata for document

**Global options**

The `new` command has the following global options, applying to all sub-commands:

|Global Option|Required|Defaut value|Description|
|---|---|---|---|
|`--dir value`|false|`doscr`|directory where metadata files will be written to|
|`--override`|false||override if destination exists|

```bs-alert warning

Global options must be placed _before_ the subcommand (e.g. right after `new`).
```

**Command `new site`**

Helps creating master metadata for the website:

```shell
docli new --override --dir doscr site --name "My Blog"
```

Command's options:

|Option|Required|Defaut value|Description|
|---|---|---|---|
|`--name value`|true||(short) name of the website|
|`--icon value`|false|`fas fa-globe`|icon of the website (support [FontAwesome icons](https://fontawesome.com/search?m=free))|
|<span class="text-nowrap">`--languages value`</span>|false|`en:english,default:en`|content's supported languages, in format `<code1:label1>[,<code2:label2>...],<default:code>`|
|`--mode value`|false|`document`|site's mode, current valid values are `document` and `blog`|
|`--author.name`|false|site's default author's name|
|`--author.email`|false|site's default author's email|
|`--author.avatar`|false|site's default author's avatar url|

```bs-alert warning

Language codes are recommended to follow [ISO 639-1 codes](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes).
```

**Command `new topic`**

Helps creating topic metadata:

```shell
docli new --override --dir doscr topic --id topic1
```

Command's options:

|Option|Required|Defaut value|Description|
|---|---|---|---|
|`--id value`|true||topic's unique id|
|<span class="text-nowrap">`--icon value`</span>|false|`fas fa-book`|icon of the topic (support [FontAwesome icons](https://fontawesome.com/search?m=free))|
|`--hidden`|false| |if specified, the topic is hidden from listing on the frontend|
|`--img value`|false| |url of topic's entry image|

```bs-alert warning

Topic's id must be unique, recommended to contain only lower-cased letters (`a-z`), digits (`0-9`) and hyphens (`-`).

Name of the directory that stores topic data will be automatically prefixed with a number and dash (format `\d+-`). The number is used to determine the topic order when rendering.
```

**Command `new document`**

Helps creating document metadata:

```shell
docli new --override --dir doscr document --id doc1 --topic topic1
```

Command's options:

|Option|Required|Defaut value|Description|
|---|---|---|---|
|`--topic value`|true||id of the topic where the document belongs to|
|`--id value`|true||document's unique id (within the topic)|
|`--icon value`|false|`fas fa-file`|icon of the document (support [FontAwesome icons](https://fontawesome.com/search?m=free))|
|<span class="text-nowrap">`--use-timestamp`</span>|false||use current timestamp to prefix document directory name (format `yyyyMMddHHmm`), see below (*)|
|`--img value`|false| |url of document's entry image|
|`--page value`|false| |set document's role as a special page|
|`--style value`|false| |set document's special style|
|`--author.name`|false|document's author's name|
|`--author.email`|false|document's author's email|
|`--author.avatar`|false|document's author's avatar url|

```bs-alert warning

Document's id must be unique (within the topic), recommended to contain only lower-cased letters (`a-z`), digits (`0-9`) and hyphens (`-`).

(*) Name of the directory that stores document data will be automatically prefixed with a number and dash (format `\d+-`). The number is used to determine the topic order when rendering. If option `--use-timestamp` is supplied, the current timestamp (format `yyyyMMddHHmm`) will be used as the prefix number.
```

### Other utilities

- `docli touch --topic <topic-id> --id <doc-id>`: update document's `last-updated` timestamp.
