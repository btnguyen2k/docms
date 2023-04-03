The simplest way to get started with DO CMS locally is a Docker environment and a text editor (better if it supports Markdown).

- Firstly, create a directory structure for your website content. Details of the directory structure will come shortly.
- Secondly, start writing content for your website using your favourite text editor. DO CMS supports <a href="https://github.github.com/gfm/" target="_blank">GitHub Flavored Markdown</a> as well as Mathematics and Chemical formulas.
- Lastly, spin up a `DO CMS runtime` instance as Docker container with the directory storing website content mounted to the container.

## Directory to store website content

An example of the directory structure is as the following:

```plain
dodata/
├── 01-intro/
│   ├── 01-whatitis/
│   │   ├── index-en.md
│   │   ├── index-vi.md
│   │   └── meta.yaml
│   ├── 02-howitwork/
│   │   ├── docms-howitwork.svg
│   │   ├── index-en.md
│   │   ├── index-vi.md
│   │   └── meta.yaml
│   ├── 03-markdown/
│   │   ├── index-en.md
│   │   ├── index-vi.md
│   │   └── meta.yaml
│   └── meta.yaml
├── 02-quickstart/
│   ...
│   └── meta.yaml
├── 03-components/
│   ...
└── meta.yaml
```

- Top level is the `root directory` that stores all website content.
- The next level is `topic directory` that stores documents belonging to a topic.
  - The `topic directory` follows the pattern `(\d+)-(\w+)`, where the first part determines the sorting order of topics, and the last part is the topic id. The two parts are separated by a dash (`-`).
  - Topic id must be unique and should be lower-cased.
- The last level is `document directory` that stores assets of a document (such as Markdown and image files).
  - The `document directory` follows the pattern `(\d+)-(\w+)`, where the first part determines the sorting order of documents, and the last part is the document id. The two parts are separated by a dash (`-`).
  - Document id must be unique within a topic and should be lower-cased.
- `DO CMS runtime`, by default, loads website content from the directory `dodata`. Hence it's convenient to name the root directory `dodata` too.
- Each `root directory`, `topic directory` and `document directory` has a `meta.yaml` file:
  - `meta.yaml` of `root directory` contains *metadata* of the website (such as name, description or contact information).
  - `meta.yaml` of `topic directory` contains *metadata* of the topic (such as title or description).
  - `meta.yaml` of `document directory` contains *metadata* of the document (such as title or summary).

> Metadata can also be stored in `meta.json` file. If both files `meta.yaml` and `meta.json` exist, the YAML file has priority.

## Render website content with DO CMS runtime

Once your website content is ready, it's time to view it with `DO CMS runtime`. The simplest way to do that is to spin up a Docker container version of `DO CMS runtime` with the `root directory` mounted to the container.

Here is a sample command to spin up such `DO CMS runtime` instance:

```shell
docker run -p 8000:8000 -v ./dodata:/app/dodata btnguyen2k/docmsrutime:latest
```

Then open http://localhost:8000 in a browser.