```bs-alert warning flex
<i class="fas fa-triangle-exclamation fa-2xl me-2"> </i>
Các cú pháp Markdown mô tả trong tài liệu này áp dụng cho các giao diện đồ hoạ đi kèm với phiên bản `DO CMS runtime` gốc. Nếu bạn sử dụng một bộ giao diện đồ hoạ của bên thứ 3, vui lòng kiểm tra tài liệu đi kèm.
```

## GitHub Flavored Markdown

DO CMS hỗ trợ [GitHub Flavored Markdown](https://github.github.com/gfm/), hay thường được viết tắt là GFM. Vui lòng xem tài liệu của GFM để biết chi tiết.

## Phong cách Bootstrap

Giao diện đồ hoạ đi kèm với `DO CMS runtime` đã bao gồm sẵn [Bootstrap](https://getbootstrap.com) (phiên bản v5.x). Bạn có thể sử dụng các phong cách và tính năng của Bootstrap trong tài liệu với cú pháp sau:

```bs-tabs
    [[bs-tab Markdown
    Đây là &lt;button type="button" class="btn btn-primary">nút nhấn&lt;/button> với phong cách Bootstrap.
    ]]
    [[bs-tab Hiển thị
    Đây là <button type="button" class="btn btn-primary">nút nhấn</button> với phong cách Bootstrap.
    ]]
```

## Biểu tượng

Bộ biểu tượng miễn phí của [FontAwesome](https://fontawesome.com/search?m=free) (phiên bản v6.x) và [Bootstrap icons](https://icons.getbootstrap.com/) cũng được kèm sẵn để sử dụng.

```bs-tabs
    [[bs-tab Markdown
    Đây là ví dụ icon của Bootstrap &lt;i class="bi bi-globe">&lt;/i>
    và đây là icon của FontAwesome &lt;i class="fas bi-book">&lt;/i>
    ]]
    [[bs-tab Hiển thị
    Đây là ví dụ icon của Bootstrap <i class="bi bi-globe"></i>
    và đây là icon của FontAwesome <i class="fas bi-book"></i>
    ]]
```

## GitHub Gist

[GitHub Gist](https://gist.github.com) có thể được nhúng vào văn bản với cú pháp sau:

```
    ```gh-gist github-username/github-gist-id
    ```
```
Ví dụ:
```bs-tabs
    [[bs-tab Markdown
        ```gh-gist btnguyen2k/d7577db0981cb157ae95b67b9f6dd733
        ```
    ]]
    [[bs-tab Hiển thị
    ```gh-gist btnguyen2k/d7577db0981cb157ae95b67b9f6dd733
    ```
    ]]
```

## Công thức Toán học và Hoá học

Các công thức Toán học và Hoá học cũng được hỗ trợ theo cú pháp [KaTeX](https://katex.org/docs/support_table.html) và [mhchem](https://mhchem.github.io/MathJax-mhchem/).

Công thức Toán học nhúng được đặt nằm giữa 2 dấu đô-la-đơn (`$`):
```bs-tabs
    [[bs-tab Markdown
        Đây là 1 công thức Toán nhúng: $x^2 + y^2 = z^2$
    ]]
    [[bs-tab Hiển thị
    Đây là 1 công thức Toán nhúng: $x^2 + y^2 = z^2$
    ]]
```

Cú pháp tương tự với các công thức Hoá học:
```bs-tabs
    [[bs-tab Markdown
        Tổng hợp nước cách nào nhỉ? $\ce{2H2 + O2 -> H2O}$
    ]]
    [[bs-tab Hiển thị
    Tổng hợp nước cách nào nhỉ? $\ce{2H2 + O2 -> H2O}$
    ]]
```

Công thức Toán dạng khối:
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
    [[bs-tab Hiển thị
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

Và công thức Hoá dạng khối:
```bs-tabs
    [[bs-tab Markdown
        ```katex
        \ce{Zn^2+  <=>[+ 2OH-][+ 2H+]  $\underset{\text{amphoteres Hydroxid}}{\ce{Zn(OH)2 v}}$  <=>[+ 2OH-][+ 2H+]  $\underset{\text{Hydroxozikat}}{\ce{[Zn(OH)4]^2-}}$}
        ```
    ]]
    [[bs-tab Hiển thị
    ```katex
    \ce{Zn^2+  <=>[+ 2OH-][+ 2H+]  $\underset{\text{amphoteres Hydroxid}}{\ce{Zn(OH)2 v}}$  <=>[+ 2OH-][+ 2H+]  $\underset{\text{Hydroxozikat}}{\ce{[Zn(OH)4]^2-}}$}
    ```
    ]]
```

## Code highlighting

Hầu hết các ngôn ngữ lập trình phổ biến đều được hỗ trợ:


```bs-tabs
    [[bs-tab Markdown
        Ứng dụng đầu tiên của tôi:
        ```cpp
        #include <iostream>

        int main() {
            std::cout << "Hello World";
            return 0;
        }
        ```

        Trang web đầu tiên:
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

        Và bộ tài liệu:
        ```markdown
        # Welcome
        _Hello_ **world**!
        ```
    ]]
    [[bs-tab Hiển thị
    Ứng dụng đầu tiên của tôi:
    ```cpp
    #include <iostream>

    int main() {
        std::cout << "Hello World";
        return 0;
    }
    ```

    Trang web đầu tiên:
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

    Và bộ tài liệu:
    ```markdown
    # Welcome
    _Hello_ **world**!
    ```
    ]]
```

Xem thêm:
- Hỗ trợ [giao diện Bootstrap](../bootstrap/).