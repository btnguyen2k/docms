Working with DO CMS is similar to a developer's daily tasks: documents are written in [Markdown](https://www.markdownguide.org/) and then pushed to a Git repository to trigger CI/CD pipelines. The CI/CD pipeline firstly invokes `DOCLI` tool to pre-process the website's content, and then containerizes the pre-processed content together with the `DO CMS runtime` as a Docker image. Finally, the package can be deployed as a website to end users.

The flow is summarized as the diagram:

```mermaid
flowchart LR
    Author([fa:fa-user<br>Content author])==1:push document==>Git{Git repo}
    subgraph CICD[3:CI/CD]
        DOCLI(A><code>DOCLI</code>: Pre-process content)-->P(B>Package pre-processed data together with<br><code>DO CMS Runtime</code> as Docker image)
    end
    Git<==2:pull content==>CICD
    CICD==4:push image==>CR[(Container Registry)]
    CICD==5:deploy==>Env{{Dev/UAT/Prod<br>env}}
    Env==6:pull image==>CR
    Users([fa:fa-users<br>Users])<==visit website==>Env
```

See more:
- [DOCLI tool](../../components/cli/)
- [DO CMS runtime](../../components/runtime/)
- [Supported Markdown syntax](../../reference/markdown/)
