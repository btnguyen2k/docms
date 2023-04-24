```bs-alert warning flex
<i class="fas fa-triangle-exclamation fa-2xl me-3"> </i>
Markdown syntaxes detailed in this document apply to the frontend templates included in the offical `DO CMS runtime`. If you use custom frontend templates from 3rd parties, please check their documentations.
```

Since Markdown supports HTML, you can embed [Bootstrap components](https://getbootstrap.com/docs/5.0/components/) in the document using HTML/CSS. However, it would be more convenient if there are shorter Markdown syntaxes that are easier to remember and to use for that purpose.

`DO CMS runtime` supports following Bootstrap components:

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

If the alert box _does not need a title_, leave **the first line** blank:
```bs-tabs
    [[bs-tab Markdown
        ```bs-alert info

        This alert box does not have a title, and it has a different style, too.
        By the way, [Markdown](https://www.markdownguide.org/) can be used in the _alert body_.
        ```
    ]]
    [[bs-tab Rendered result
    ```bs-alert info

    This alert box does not have a title, and it has a different style, too.
    By the way, [Markdown](https://www.markdownguide.org/) can be used in the _alert body_.
    ```
    ]]
```

All Bootstrap alert styles are available: `primary`, `secondary`, `success`, `danger`, `warning`, `info`, `light`, `dark`.

If you want an alert box with an icon on one side, add `flex` attribute to the alert box, also reserve the title section for the icon:
```bs-tabs
    [[bs-tab Markdown
        ```bs-alert warning flex
        <i class="fas fa-triangle-exclamation fa-xl me-2"></i>
        This `warning` alert has a nice [FontAwesome icon](https://fontawesome.com/icons/triangle-exclamation?f=classic&s=solid&sz=xl) on the left.<br/>
        Remember to reserve the title section for the icon itself.
        ```
    ]]
    [[bs-tab Rendered result
    ```bs-alert warning flex
    <i class="fas fa-warning fa-xl me-2"></i>
    This `warning` alert has a nice [FontAwesome icon](https://fontawesome.com/icons/triangle-exclamation?f=classic&s=solid&sz=xl) on the left.<br/>
    Remember to reserve the title section for the icon itself.
    ```
    ]]
```

Of course, you can also use [Bootstrap icons](https://icons.getbootstrap.com/):
```bs-tabs
    [[bs-tab Markdown
        ```bs-alert primary flex
        <i class="bi bi-info-circle me-2" style="font-size: 2rem;"></i>
        [GitHub Flavored Markdown](https://github.github.com/gfm/#what-is-github-flavored-markdown-), often shortened as GFM, is the dialect of Markdown that is currently supported for user content on **GitHub.com** and **GitHub Enterprise**.
        ```
    ]]
    [[bs-tab Rendered result
    ```bs-alert primary flex
    <i class="bi bi-info-circle me-2" style="font-size: 2rem;"></i>
    [GitHub Flavored Markdown](https://github.github.com/gfm/#what-is-github-flavored-markdown-), often shortened as GFM, is the dialect of Markdown that is currently supported for user content on **GitHub.com** and **GitHub Enterprise**.
    ```
    ]]
```

## Bootstrap Tabs

A [Bootstrap tabs](https://getbootstrap.com/docs/5.0/components/navs-tabs/) block is embedded in the document using the following syntax:

```markdown
    ```bs-tabs
    [[bs-tab Tab1
    Content of the first tab,
    can span multiple lines.
    ]]

    [[bs-tab Second tab
    And the content, can use _Markdown_ too.
    ]]
    ```
```
The Markdown snippet above produces the following result:
```bs-tabs
[[bs-tab Tab1
Content of the first tab,
can span multiple lines.
]]

[[bs-tab Second tab
And the content, can use _Markdown_ too.
]]
```

Tabs can also be placed vertically. The following Markdown snippet:
```markdown
    ```bs-tabs vertical
    [[bs-tab Tab1
    Content 1 in vertically aligned tabs.
    ]]

    [[bs-tab Tab2
    Content 2 in vertically aligned tabs.
    ]]
    ```
```
produces the following result:
```bs-tabs vertical
[[bs-tab Tab1
Content 1 in vertically aligned tabs.
]]

[[bs-tab Tab2
Content 2 in vertically aligned tabs.
]]
```

See more:
- [Supported Markdown syntaxes](../markdown/)
