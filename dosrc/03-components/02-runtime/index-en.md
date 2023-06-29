`DO CMS runtime` is responsible for rendering website content. It also offers other features such as _i18n_ and _fulltext search_.

For your convenience, `DO CMS runtime` has already been containerized as a [Docker image](https://hub.docker.com/r/btnguyen2k/docmsruntime).

The following command spins up an instance of `DO CMS runtime` from Docker image:

```shell
docker run --name docms \
    -p 8000:8000 \
    -v /path/to/my/dodata:/app/dodata \
    btnguyen2k/docmsruntime
```

```bs-alert info flex
<i class="fas fa-circle-info fa-2xl me-3"></i>
The Docker image offer the following tags:

`:stable` or `:version` (i.e. `:1.2.3`): stable releases<br/>
`:latest` or `:yyyymmdd` or `:yyyymmdd-hash`: usable releases, may have some glitches but always include latest features up to the build time.
```

## Environment Variables

### HTTP_LISTEN_ADDR/HTTP_LISTEN_PORT

Specify listen address & port number. Default value for HTTP_LISTEN_ADDR is `0.0.0.0` and HTTP_LISTEN_PORT is `8000`. The following command makes `DO CMS runtime` to listen on `127.0.0.1:3000`:
```shell
docker run --name docms \
    -p 8000:3000 \
    -v /path/to/my/dodata:/app/dodata \
    -e HTTP_LISTEN_ADDR=127.0.0.1 -e HTTP_LISTEN_PORT=3000 \
    btnguyen2k/docmsruntime
```

### GOOGLE_TAG

Specify Google Tag ID for Google Analytics, example:
```shell
docker run --name docms \
    -p 8000:8000 \
    -e GOOGLE_TAG=G-ABC123D4E5 \
    btnguyen2k/docmsruntime
```

```bs-alert info

`GOOGLE_TAG` is supported since runtime `v0.3.1` or Docker image `:20230426`.
```

### DOCMS_DATA_DIR

Specify the root directory where website content is located. Default value is `./dodata`. The following command makes `DO CMS runtime` to load website content from directory `/mydata`:
```shell
docker run --name docms \
    -p 8000:8000 \
    -v /path/to/my/dodata:/mydata \
    -e DOCMS_DATA_DIR=/mydata \
    btnguyen2k/docmsruntime
```

### FRONTEND_TEMPLATE

Specify the GUI template to be used for the frontend. Currently there are 3 available templates `bootstrap`, `coderdocs` and `prettydocs` (default template is `bootstrap`). The following command switches the template to `coderdocs`:
```shell
docker run --name docms \
    -p 8000:8000 \
    -e FRONTEND_TEMPLATE=coderdocs \
    btnguyen2k/docmsruntime
```

**Bootstrap** template:
```bs-cards cols-3 cols-sm-1 equals lightbox=Bootstrap
    [[bs-card
    -img:bootstrap1.png
    -title:Home page
    ]]

    [[bs-card
    -img:bootstrap2.png
    -title:Topic page
    ]]

    [[bs-card
    -img:bootstrap3.png
    -title:Document page
    ]]
```

**CoderDocs** template:
```bs-cards cols-3 cols-sm-1 equals lightbox=CoderDocs
    [[bs-card
    -img:coderdocs1.png
    -title:Home page
    ]]

    [[bs-card
    -img:coderdocs2.png
    -title:Topic page
    ]]

    [[bs-card
    -img:coderdocs3.png
    -title:Document page
    ]]
```

**PrettyDocs** template:
```bs-cards cols-3 cols-sm-1 equals lightbox=PrettyDocs
    [[bs-card
    -img:prettydocs1.png
    -title:Home page
    ]]

    [[bs-card
    -img:prettydocs2.png
    -title:Topic page
    ]]

    [[bs-card
    -img:prettydocs3.png
    -title:Document page
    ]]
```

```bs-alert info

New template `blogy` has been added since `v0.3.1` (or Docker image `:20230511`). `blogy` template is suitable for blog-style website (site was [created](../docli/#helper-to-create-content-metadata) with `--mode blog`).
```

### MIME_TYPES

Specify allowed MINE types for media files attached to documents, value must be a valid JSON-encoded map, for example:
```shell
docker run --name docms \
    -p 8000:8000 \
    -e MIME_TYPES='{".jpg": "image/jpeg", ".gif": "image/gif"}' \
    btnguyen2k/docmsruntime
```

```bs-alert info

`MIME_TYPES` is supported since runtime `v0.4.0` or Docker image `:20230628`.
```

The following MIME types are configured by default:
```json
  ## Media MIME types
  media_mime {
    # photo files
    ".jpg":  "image/jpeg",
    ".jpeg": "image/jpeg",
    ".png":  "image/png",
    ".gif":  "image/gif",
    ".svg":  "image/svg+xml",
    # music files
    ".aac":  "audio/aac",
    ".mp3":  "audio/mp3",
    ".wav":  "audio/wav",
    ".wma":  "audio/wma",
    # video files
    ".mp4":  "audio/mp4",
    ".mov":  "audio/mov",
    ".avi":  "audio/avi",
    # documents
    ".xml":  "application/xml",
    ".json": "application/json",
    ".pdf":  "application/pdf",
    ".csv":  "application/csv",
    ".tsv":  "application/tsv",
    ".doc":  "application/doc",
    ".docx": "application/docx",
    ".xls":  "application/xls",
    ".xlsx": "application/xlsx",
    ".ppt":  "application/ppt",
    ".pptx": "application/pptx",
    ".odm":  "application/odm",
    ".odt":  "application/odt",
    ".odp":  "application/odp",
    ".ods":  "application/ods",
  }
```

### MIME_TYPES_ADD

Add MINE types to the existing list, value must be a valid JSON-encoded map, for example:
```shell
docker run --name docms \
    -p 8000:8000 \
    -e MIME_TYPES_ADD='{".mp4": "video/mp4"}' \
    btnguyen2k/docmsruntime
```

```bs-alert info

`MIME_TYPES_ADD` is supported since runtime `v0.4.0` or Docker image `:20230628`.
```
