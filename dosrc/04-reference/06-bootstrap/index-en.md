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

## Bootstrap Cards

[Bootstrap Cards](https://getbootstrap.com/docs/5.0/components/card/) are embedded in the document using the following syntax:

<pre><code class="hljs language-markdown">
```bs-cards &lt;card-group-settings>
    [[bs-card &lt;card-settings>
    -img:url to card's header image
    -header:card's header on one line
    -title:card's title on one line
    -subtitle:card's subtitle on one line
    -footer:card's footer on one line
    All other lines are
    card's text. **Markdown** is _supported_.
    ]]

    [[bs-card &lt;card-settings>
    next card
    ]]
```
</code></pre>

```bs-alert warning flex
<i class="fas fa-triangle-exclamation me-3 fa-2xl"></i>
Be aware of the **indent**! Card definition (starts with `[bs-card` and ends with `]]`) are **4-space** indented from the card group definition (starts with <code>\```bs-cards</code> and ends with <code>\```</code>).
```

**Card group settings** - instruct how cards in this card group are displayed.
|Setting|Default Value|Description|
|---|---|---|
|`equals`|`false`|If set, all cards in the group have same height.|
|`lightbox=<gallery-name>`| |If set, 'lightbox' mode is enabled to header images of cards, and images added to the gallery specified by `gallery-name`.|
|`cols=<num-columns>`| |Number of cards per row (value between `1-12`). If not specified, cards fill up one row and then span to the next row, and so on.|
|`cols-sm=<num-columns>`| |Number of cards per row (value between `1-12`) for _small_ screen size. If not specified, cards fill up one row and then span to the next row, and so on.|
|`cols-md=<num-columns>`| |Number of cards per row (value between `1-12`) for _medium_ screen size. If not specified, cards fill up one row and then span to the next row, and so on.|
|`cols-lg=<num-columns>`| |Number of cards per row (value between `1-12`) for _large_ screen size. If not specified, cards fill up one row and then span to the next row, and so on.|

```bs-alert info

`cols`, `cols-sm`, `cols-md` and `cols-lg` can be used together to control the number of cards per row for different screen sizes.
```
```bs-alert warning

`lightbox` setting has no effect if no card in the group has header image.
Images from different card groups can belong to a same lightbox gallery (same `gallery-name`).
```

**Card elements and settings**

- The prefix `-img:` specifies the image attached to the header of the card. The image's url must be placed on the same line as the prefix.
- The prefix `-header:` specifies the card's header. The header text must be placed on the same line as the prefix.
- The prefix `-footer:` specifies the card's footer. The footer text must be placed on the same line as the prefix.
- The prefix `-title:` specifies the title of card's body. The title text must be placed on the same line as the prefix.
- The prefix `-subtitle:` specifies the subtitle of card's body. The subtitle text must be placed on the same line as the prefix.
- Other lines are card's text. Markdown is supported.

Examples:

```bs-tabs
    [[bs-tab Markdown
        ```bs-cards cols=3 lightbox=mygallery
            [[bs-card
            -img:https://placehold.co/320x200?text=Image+1
            -title:This is the card title
            -subtitle:The subtitle goes here
            All other lines are the card's body.
            **Markdown** is _supported_.
            This card has title and subtitle but no header and footer.

            Note: This card's image and the next card's image are added
            to a lightbox gallery.
            ]]

            [[bs-card
            -img:https://placehold.co/320x200?text=Image+2
            -header:This is the header
            -footer:Footer has its own section
            This card has **header and footer** but no ~title and subtitle~.

            Note: This card's image and the previous card's image are added
            to a lightbox gallery.
            ]]

            [[bs-card
            -header:<i class="fas fa-heading"></i> **Header** with icon
            -title:<i class="fas fa-1"></i> _Title_ with icon
            -subtitle:<i class="fas fa-2"></i> ~Subtitle~ with icon
            -footer:<i class="fas fa-3"></i> `Footer` with icon
            This is the 3rd card. Be noted that 3 cards have _different_ heights.

            Also, only card's text supports **Markdown**.
            Markdown in `header`, `footer`, `title` and `subtitle` will not work.
            ]]
        ```
    ]]

    [[bs-tab Rendered result
    ```bs-cards cols=3 lightbox=mygallery
        [[bs-card
        -img:https://placehold.co/320x200?text=Image+1
        -title:This is the card title
        -subtitle:The subtitle goes here
        All other lines are the card's body.
        **Markdown** is _supported_.
        This card has title and subtitle but no header and footer.

        Note: This card's image and the next card's image are added
        to a lightbox gallery.
        ]]

        [[bs-card
        -img:https://placehold.co/320x200?text=Image+2
        -header:This is the header
        -footer:Footer has its own section
        This card has **header and footer** but no ~title and subtitle~.

        Note: This card's image and the previous card's image are added
        to a lightbox gallery.
        ]]

        [[bs-card
        -header:<i class="fas fa-heading"></i> **Header** with icon
        -title:<i class="fas fa-1"></i> _Title_ with icon
        -subtitle:<i class="fas fa-2"></i> ~Subtitle~ with icon
        -footer:<i class="fas fa-3"></i> `Footer` with icon
        This is the 3rd card. Be noted that 3 cards have different _heights_.

        Also, only card's text supports **Markdown**.
        Markdown in `header`, `footer`, `title` and `subtitle` will not work.
        ]]
    ```
    ]]
```

Card settings:

|Setting|Default Value|Description|
|---|---|---|
|`no-muted`|`false`|Card's subtitle and footer text is assigned CSS class `text-muted`. If `no-muted` is set, CSS class `text-muted` is removed.|
|`text=<style>`| |Style for card's text. Accept any valid value of [Bootstrap's text colors](https://getbootstrap.com/docs/5.0/utilities/colors/#colors) (without the prefix `text-`!).|
|`bg=<style>`| |Style for card's background. Accept any valid value of [Bootstrap's background colors](https://getbootstrap.com/docs/5.0/utilities/background/#background-color) (without the prefix `bg-`!).|
|`border=<style>`| |Style for card's border. Accept any valid value of [Bootstrap's border colors](https://getbootstrap.com/docs/5.0/utilities/borders/#border-color) (without the prefix `border-`!).|

```bs-alert info

`text`, `bg` and `border` can be mixed together.
```

Examples:

```bs-tabs
    [[bs-tab Markdown
        ```bs-cards cols-sm=1 cols-md=2 cols-lg=3 equals
            [[bs-card border=primary text=info
            -footer:Footer without "no-muted"
            This card has _custom style_ for border and text.
            ]]

            [[bs-card bg=success text=light no-muted
            -footer:Footer with "no-muted"
            This card has _custom style_ for background and text.

            See footer's style with `no-muted` setting.
            ]]

            [[bs-card bg=secondary text=light no-muted
            -header:Header
            This card has _custom style_ for background and text.

            Card's header is not affected by `no-muted` setting.
            ]]

            [[bs-card bg=dark text=light
            -title:Title
            -subtitle:Subtitle without "no-muted"
            This card has _custom style_ for background and text.
            ]]

            [[bs-card bg=light text=dark no-muted
            -title:Title
            -subtitle:Subtitle with "no-muted"
            This card has _custom style_ for background and text.

            See subtitle's style with `no-muted` setting.
            ]]

            [[bs-card border=danger bg=warning text=dark no-muted
            Custom _border_, _background_ and _text_ together.

            Also be noted that all cards have same height.

            And different number of cards per row, depends on the screen size.
            ]]
        ```
    ]]

    [[bs-tab Rendered result
    ```bs-cards cols-sm=1 cols-md=2 cols-lg=3 equals
        [[bs-card border=primary text=info
        -footer:Footer without "no-muted"
        This card has _custom style_ for border and text.
        ]]

        [[bs-card bg=success text=light no-muted
        -footer:Footer with "no-muted"
        This card has _custom style_ for background and text.

        See footer's style with `no-muted` setting.
        ]]

        [[bs-card bg=secondary text=light no-muted
        -header:Header
        This card has _custom style_ for background and text.

        Card's header is not affected by `no-muted` setting.
        ]]

        [[bs-card bg=dark text=light
        -title:Title
        -subtitle:Subtitle without "no-muted"
        This card has _custom style_ for background and text.
        ]]

        [[bs-card bg=light text=dark no-muted
        -title:Title
        -subtitle:Subtitle with "no-muted"
        This card has _custom style_ for background and text.

        See subtitle's style with `no-muted` setting.
        ]]

        [[bs-card border=danger bg=warning text=dark no-muted
        Custom _border_, _background_ and _text_ together.

        Also be noted that all cards have same height.

        And different number of cards per row, depends on the screen size.
        ]]
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
