A quick recap of the directory structure that stores website content:

- Top level is the `root directory` that stores all website content.
- The next level is `topic directory` that stores documents belonging to a topic.
- The last level is `document directory` that stores assets of a document (sucA quick recap of the directory structure that store website content:

- Top level is the `root directory` that stores all website content.
- The next level is `topic directory` that stores documents belonging to a topic.
- The last level is `document directory` that stores assets of a document (such as Markdown and image files).

Each directory has a _metadata file_ right underneath it:
- The metadata file of the `root directory` contains information such as name, description or contact details of the website; the metadata file of the `topic directory` contains information such as title or description of the topic; and the metadata file of the `document directory` contains information such as title or summary of the document.
- The metadata file is in either YAML (named `meta.yaml` or `meta.yml`) or JSON format (named `meta.json`). If multiple files exist, `meta.yaml` has the highest priority, next is `meta.yml` and `meta.json` is last.

An example of the directory structure is as the following:

```
dodata/                      <-- root directory
├── 01-intro/                <---- topic directory, id of this topic is "intro"
│   ├── 01-whatitis/         <------ document directory, this document belong to topic "intro"
│   │   ├── index-en.md      <-------- content file (in language  "en") of document "whatitis"
│   │   ├── index-vi.md      <-------- content file (in language  "vi") of document "whatitis"
│   │   └── meta.yaml        <-------- metadata file of document "whatitis"
│   ├── 02-howitwork/        <------ document directory
│   │   ...
│   │   └── meta.yaml        <-------- metadata file of document "howitwork"
│   ├── 03-markdown/         <------ document directory
│   │   ...
│   │   └── meta.yaml        <-------- metadata file of document "markdown"
│   └── meta.yaml            <------ metadata file of topic "intro"
├── 02-quickstart/           <---- topic directory
│   ...
│   └── meta.yaml            <------ metadata file of topic "quickstart"
├── 03-components/           <---- topic directory
│   ...
└── meta.yaml                <---- metadata of website, located right underneath the root directory
```

See also:
- [Site metadata file](../sitemetadata/)
- [Topic metadata file](../topicmetadata/)
- [Docment metadata file](../documentmetadata/)
