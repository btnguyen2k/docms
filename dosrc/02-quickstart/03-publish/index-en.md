It's time to publish your website to its users. The quickest and easiest way to do so is bundling the directory `dodata` together with a `DO CMS runtime` instance in a Docker image. Then deploy the image to your favourite container environment.

Since `DO CMS runtime` has its own Docker image, the following Dockerfile snippet create a new image with the directory `dodata` bundled with the `DO CMS runtime`:

```dockerfile
FROM btnguyen2k/docmsruntime:latest as docmsruntime
LABEL maintainer="Your name"
COPY ./dodata /app/dodata
WORKDIR /app
EXPOSE 8000
CMD ["/app/main"]
```

If you wish to pre-process website content before bundling with `DO CMS runtime`:

```dockerfile
FROM golang:1.17-alpine AS builder_docs
RUN mkdir /build
COPY . /build
RUN cd /build \
    && go install github.com/btnguyen2k/docms/docli \
    && docli build --purge --src dosrc --out dodata

FROM btnguyen2k/docmsruntime:latest as docmsruntime
LABEL maintainer="Your name"
COPY --from=builder_docs /build/dodata /app/dodata
WORKDIR /app
EXPOSE 8000
CMD ["/app/main"]
```

Those steps can be done manually, or automatically via CI/CD pipelines. A working GitHub Actions flow that pre-processes website content, bundles with `DO CMS runtime` and publishes to Azure Container App can be found <a href="https://github.com/btnguyen2k/docms/blob/main/.github/workflows/dodocs.yml" target="_blank">here</a>.