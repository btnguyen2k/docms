You may find that some features of DO CMS do not work (for example fulltext search) when website content is rendered directly from the `root directory`. In order to fully enable all the features, website content needs to be pre-processed before being rendered by `DO CMS runtime`.

Website content is pre-processed via `DOCLI` tool.  Install it with the following command:

```shell
go install github.com/btnguyen2k/docms/docli@latest
```

Let's put the website content in `dosrc` directory and run the following command to pre-process:

```shell
docli build --src ./dosrc --out ./dodata
```

`DOCLI` tool performs the following tasks:
- Check the website content in the source directory (`dosrc`) for errors.
- Copy website content from the source directory (`dosrc`) to the output directory (`dodata`).
- Build fulltext index and store the index in the output directory (`dodata`).

After pre-processing, rendering website content from `dodata` directory unlocks all DO CMS features.

> `DOCLI` is installed to directory `$GOPATH/bin`. You may need to add the directory to system's PATH.

See more:
- [DOCLI tool](../../components/cli/)
- [DO CMS runtime](../../components/runtime/)
