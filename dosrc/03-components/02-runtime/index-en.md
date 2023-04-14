`DO CMS runtime` is responsible for rendering website content. It also offers other features such as i18n and fulltext search.

For your convenience, `DO CMS runtime` has already been containerized as a [Docker image](https://hub.docker.com/r/btnguyen2k/docmsruntime).

The following command spins up an instance of `DO CMS runtime` from Docker image:

```shell
docker run --name docms \
    -p 8000:8000 \
    -v /path/to/my/dodata:/app/dodata \
    btnguyen2k/docmsruntime
```

## Environment Variables

**HTTP_LISTEN_ADDR** and **HTTP_LISTEN_PORT** - specify listen address & port number. Default value for HTTP_LISTEN_ADDR is `0.0.0.0` and HTTP_LISTEN_PORT is `8000`. The following command makes `DO CMS runtime` to listen on `127.0.0.1:3000`:
```shell
docker run --name docms \
    -p 8000:3000 \
    -v /path/to/my/dodata:/app/dodata \
    -e HTTP_LISTEN_ADDR=127.0.0.1 -e HTTP_LISTEN_PORT=3000 \
    btnguyen2k/docmsruntime
```

**DOCMS_DATA_DIR** - specify the root directory where website content is located. Default value is `./dodata`. The following command makes `DO CMS runtime` to load website content from directory `/mydata`:
```shell
docker run --name docms \
    -p 8000:8000 \
    -v /path/to/my/dodata:/mydata \
    -e DOCMS_DATA_DIR=/mydata \
    btnguyen2k/docmsruntime
```
