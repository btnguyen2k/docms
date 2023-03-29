DO CMS supports <a href="https://github.github.com/gfm/" target="_blank">GitHub Flavored Markdown</a>, or often shortened as GFM, with <a href="https://getbootstrap.com/docs/5.0/content/tables/" target="_blank">Bootstrap</a> style.

Some notable Markdown features supported by Do CMS are as the following:

## Tables

|col-1|col-2|
|---|---|
|row1 - content1|row1 - content2|
|row2 - content1|row2 - content2|
|row3 - content1|row3 - content2|

## Block quotes

> Block quotes feature is supported, stylist.

## Task list items

- [x] Task 1 is done
  - [ ] Nested task is not done
  - [x] Another nested task, which is done
- [ ] Another task

## Strikethrough

This text is ~~strikethrough~~.

> Even ~~inside a block quote~~.

## Code highlighting

Inline code spans with single backstick (\`) , for example `javascript let i = 0`, as well as double backsticks (\``), for example ``for (i := 0; i < 10; i++)``, are supported.

Code blocks with three backsticks (```) are highlighted:

```cpp
// most common programming are supported
#include <iostream>

int main() {
    std::cout << "Hello World!";
    return 0;
}
```

## Mathematical and Chemical formulas

DO CMS supports Mathematical and Chemical formulas using <a href="https://katex.org/docs/support_table.html" target="_blank">KaTeX</a> and <a href="https://mhchem.github.io/MathJax-mhchem/" target="_blank">mhchem</a>.

Inline Mathematical formulas are enclosed between two single-dollar signs (`$`): $x^2 + y^2 = z^2$

Block Mathematical formulas are enclosed between two double-dollar signs (`$$`):
$$
\begin{equation}
\begin{split}
(a - b)^2 &= (a - b)(a - b) \\
&= a(a - b) - b(a - b)      \\
&= a^2 -ab -ba + b^2        \\
&= a^2 + 2ab + b^2          \nonumber
\end{split}
\end{equation}
$$

Similar syntax for inline Chemical formulas: $\ce{CO2 + C -> 2 CO}$

and block Chemical formulas:
$$
C_p[\ce{H2O(l)}] = \pu{75.3 J // mol K}
$$

> Note: The two double-dollar signs (`$$`) must be on their own lines in order for block formulas to work.
> That means, the following will not work: `$$a+b=c$$`
