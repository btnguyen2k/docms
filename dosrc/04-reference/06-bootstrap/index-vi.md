```bs-alert warning flex
<i class="fas fa-triangle-exclamation fa-2xl me-2"> </i>
Markdown syntaxes detailed in this document apply to the frontend templates included in the offical `DO CMS runtime`. If you use custom frontend templates from 3rd parties, please check their documentations.
```

Since Markdown supports HTML, you can embed [Bootstrap components](https://getbootstrap.com/docs/5.0/components/) in the document using HTML/CSS. However, it would be more convenient if there are shorter Markdown syntaxes that are easier to remember and to use for that purpose.

The following Bootstrap components are supported:

## Bootstrap Alerts

[Bootstrap Alerts](https://getbootstrap.com/docs/5.0/components/alerts/) can be embedded in the document using the following syntax:
```bs-tabs
    [[bs-tab Markdown
        ```bs-alert primary
        First line is reserved for the title
        Following lines are for
        the body
        ```
    ]]
    [[bs-tab Rendered result
    ```bs-alert primary
    First line is reserved for the title
    Following lines are for
    the body
    ```
    ]]
```

If the alert box _does not need a title_, leave **the first line** blank.

```bs-tabs
    [[bs-tab Markdown
        ```bs-alert info

        This alert box does not have a title, and it has a different style, too.
        ```
    ]]
    [[bs-tab Rendered result
    ```bs-alert info

    This alert box does not have a title, and it has a different style, too.
    ```
    ]]
```

All Bootstrap alert styles are available: `primary`, `secondary`, `success`, `danger`, `warning`, `info`, `light`, `dark`.
