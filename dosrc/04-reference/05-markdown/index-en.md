```bs-alert warning flex
<i class="fas fa-triangle-exclamation fa-2xl me-3"> </i>
Markdown syntaxes detailed in this document apply to the frontend templates included in the offical `DO CMS runtime`. If you use custom frontend templates from 3rd parties, please check their documentations.
```

## GitHub Flavored Markdown

DO CMS supports [GitHub Flavored Markdown](https://github.github.com/gfm/), or often shortened as GFM. Please check GFM's documentation for syntax details.

## Bootstrap style

Frontend templates of `DO CMS runtime` import [Bootstrap](https://getbootstrap.com) by default (currently v5.x). You are free to use any Bootstrap style or feature. For example:

```bs-tabs
    [[bs-tab Markdown
    This is a Bootstrap &lt;button type="button" class="btn btn-primary">button&lt;/button>
    ]]
    [[bs-tab Rendered result
    This is a Bootstrap <button type="button" class="btn btn-primary">button</button>
    ]]
```

## Icons

Free version of [FontAwesome icons](https://fontawesome.com/search?m=free) (currently v6.x) and [Bootstrap icons](https://icons.getbootstrap.com/) are also available for use.

```bs-tabs
    [[bs-tab Markdown
    This is a Bootstrap icon &lt;i class="bi bi-globe">&lt;/i>
    and this is a FontAwesome icon &lt;i class="fas bi-book">&lt;/i>
    ]]
    [[bs-tab Rendered result
    This is a Bootstrap icon <i class="bi bi-globe"></i><br/>
    and this is a FontAwesome icon <i class="fas bi-book"></i>
    ]]
```

## Tags from site metadata file

Assuming the [site metadata file](../sitemetadata/) contains the following tags section:
```yaml
tags:
  build: ${build_datetime}
  demo:
    tags1:
      key1: value 1
      key2: value 2
    tags2: [1, "2", true]
```

Tags defined in the file can be embedded to the document using the syntax `[[do-tag tag-key]]`:
```bs-tabs
    [[bs-tab Markdown
        This document was updated on [[do-tag build]].

        Fetching tag value from a hierarchy structure: <code>[[do-tag demo.tag1.key1]]</code><br/>
        And also from an array (0-based): **[[do-tag demo.tag2.2]]**

        If tag does not exist: [[do-tag demo.tag3]]<br/>
        Or out of array index: [[do-tag demo.tag2.-1]] / [[do-tag demo.tag2.5]]<br/>
    ]]
    [[bs-tab Rendered result
    This document was updated on [[do-tag build]].

    Fetching tag value from a hierarchy structure: <code>[[do-tag demo.tag1.key1]]</code><br/>
    And also from an array (0-based): **[[do-tag demo.tag2.2]]**

    If tag does not exist: [[do-tag demo.tag3]]<br/>
    Or out of array index: [[do-tag demo.tag2.-1]] / [[do-tag demo.tag2.5]]<br/>
    ]]
```

## GitHub Gist

[GitHub Gist](https://gist.github.com) can be embedded to the document using the following syntax:

```
    ```gh-gist github-username/github-gist-id
    ```
```
For example:
```bs-tabs
    [[bs-tab Markdown
        ```gh-gist btnguyen2k/d7577db0981cb157ae95b67b9f6dd733
        ```
    ]]
    [[bs-tab Rendered result
    ```gh-gist btnguyen2k/d7577db0981cb157ae95b67b9f6dd733
    ```
    ]]
```

## Mathematical and Chemical formulas

Mathematical and Chemical formulas are supported following [KaTeX](https://katex.org/docs/support_table.html) and [mhchem](https://mhchem.github.io/MathJax-mhchem/) syntax.

Inline Mathematical formulas are enclosed between two single-dollar signs (`$`):
```bs-tabs
    [[bs-tab Markdown
        This is an inline Maths formula: $x^2 + y^2 = z^2$
    ]]
    [[bs-tab Rendered result
    This is an inline Maths formula: $x^2 + y^2 = z^2$
    ]]
```

The same syntax for inline Chemical formulas:
```bs-tabs
    [[bs-tab Markdown
        Can we make water? Probably: $\ce{2H2 + O2 -> H2O}$
    ]]
    [[bs-tab Rendered result
    Can we make water? Probably: $\ce{2H2 + O2 -> H2O}$
    ]]
```

Block Mathematical formulas:
```bs-tabs
    [[bs-tab Markdown
        ```katex
        \begin{equation}
        \begin{split}
        (a - b)^2 &= (a - b)(a - b) \\
        &= a(a - b) - b(a - b)      \\
        &= a^2 -ab -ba + b^2        \\
        &= a^2 + 2ab + b^2          \nonumber
        \end{split}
        \end{equation}
        ```
    ]]
    [[bs-tab Rendered result
    ```katex
    \begin{equation}
    \begin{split}
    (a - b)^2 &= (a - b)(a - b) \\
    &= a(a - b) - b(a - b)      \\
    &= a^2 -ab -ba + b^2        \\
    &= a^2 + 2ab + b^2          \nonumber
    \end{split}
    \end{equation}
    ```
    ]]
```

Block Chemical formulas:
```bs-tabs
    [[bs-tab Markdown
        ```katex
        \ce{Zn^2+  <=>[+ 2OH-][+ 2H+]  $\underset{\text{amphoteres Hydroxid}}{\ce{Zn(OH)2 v}}$  <=>[+ 2OH-][+ 2H+]  $\underset{\text{Hydroxozikat}}{\ce{[Zn(OH)4]^2-}}$}
        ```
    ]]
    [[bs-tab Rendered result
    ```katex
    \ce{Zn^2+  <=>[+ 2OH-][+ 2H+]  $\underset{\text{amphoteres Hydroxid}}{\ce{Zn(OH)2 v}}$  <=>[+ 2OH-][+ 2H+]  $\underset{\text{Hydroxozikat}}{\ce{[Zn(OH)4]^2-}}$}
    ```
    ]]
```

## Code highlighting

Popular programming languages are supported:


```bs-tabs
    [[bs-tab Markdown
        My first program:
        ```cpp
        #include <iostream>

        int main() {
            std::cout << "Hello World";
            return 0;
        }
        ```

        My first site:
        ```html
        <!DOCTYPE html>
        <html>
        <head>
            <title>Welcome</title>
            <style>body {color: blue;}</style>
            <script type="application/javascript">
            function $init() {return true;}
            </script>
        </head>
        <body>
            <h1>Hello World!</h1>
        </body>
        </html>
        ```

        And my first document:
        ```markdown
        # Welcome
        _Hello_ **world**!
        ```
    ]]
    [[bs-tab Rendered result
    My first program:
    ```cpp
    #include <iostream>

    int main() {
        std::cout << "Hello World";
        return 0;
    }
    ```

    My first site:
    ```html
    <!DOCTYPE html>
    <html>
    <head>
        <title>Welcome</title>
        <style>body {color: blue;}</style>
        <script type="application/javascript">
        function $init() {return true;}
        </script>
    </head>
    <body>
        <h1>Hello World!</h1>
    </body>
    </html>
    ```

    And my first document:
    ```markdown
    # Welcome
    _Hello_ **world**!
    ```
    ]]
```

See more:
- [Bootstrap components](../bootstrap/) support.
