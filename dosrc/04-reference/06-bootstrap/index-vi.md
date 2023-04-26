```bs-alert warning flex
<i class="fas fa-triangle-exclamation fa-2xl me-3"> </i>
Các cú pháp Markdown mô tả trong tài liệu này áp dụng cho các giao diện đồ hoạ đi kèm với phiên bản `DO CMS runtime` gốc. Nếu bạn sử dụng một bộ giao diện đồ hoạ của bên thứ 3, vui lòng kiểm tra tài liệu đi kèm.
```

Vì Markdown cũng hỗ trợ HTML, bạn có thể nhúng các [giao diện Bootstrap](https://getbootstrap.com/docs/5.0/components/) vào tài liệu thông qua HTML/CSS. Tuy nhiên, sẽ tiện hơn nếu có cú pháp Markdown ngắn gọn hơn, dễ nhớ và dễ sử dụng hơn để làm chuyện đó.

`DO CMS runtime` hỗ trợ các giao diện Bootstrap sau:

## Bootstrap Alert

Nhúng [Bootstrap Alert](https://getbootstrap.com/docs/5.0/components/alerts/) vào tài liệu thông qua cú pháp sau:
```bs-tabs
    [[bs-tab Markdown
        ```bs-alert primary
        Dòng đầu tiên dành riêng cho phần tiêu đề
        Các dòng tiếp theo là nội dung
        của alert.
        ```
    ]]
    [[bs-tab Hiển thị
    ```bs-alert primary
    Dòng đầu tiên dành riêng cho phần tiêu đề
    Các dòng tiếp theo là nội dung
    của alert.
    ```
    ]]
```

Nếu bạn muốn _không có tiêu đề_, hãy để trống **dòng đầu tiên**:
```bs-tabs
    [[bs-tab Markdown
        ```bs-alert info

        Hộp alert này không có phần tiêu đề, và nó cũng sử dụng tông màu khác luôn.
        Mà nhân tiện, [Markdown](https://www.markdownguide.org/) cũng có thể được _sử dụng_ trong phần **nội dung** nhé.
        ```
    ]]
    [[bs-tab Hiển thị
    ```bs-alert info

    Hộp alert này không có phần tiêu đề, và nó cũng sử dụng tông màu khác luôn.
    Mà nhân tiện, [Markdown](https://www.markdownguide.org/) cũng có thể được _sử dụng_ trong phần **nội dung** nhé.
    ```
    ]]
```

Các tông màu của hộp Bootstrap alert có thể sử dụng: `primary`, `secondary`, `success`, `danger`, `warning`, `info`, `light`, `dark`.

Nếu bạn muốn một hộp alert đi kèm với 1 biểu tượng ở 1 bên, hãy thêm từ khoá `flex`, và lưu ý để biểu tượng ở dòng tiêu đề:
```bs-tabs
    [[bs-tab Markdown
        ```bs-alert warning flex
        <i class="fas fa-triangle-exclamation fa-xl me-2"></i>
        Hộp alert này sử dụng tông màu `warning` và có một [icon FontAwesome](https://fontawesome.com/icons/triangle-exclamation?f=classic&s=solid&sz=xl) ở bên trái.<br/>
        Nhớ đừng quên để riêng cái biểu tượng ở dòng dành cho tiêu đề (tức là dòng đầu tiên).
        ```
    ]]
    [[bs-tab Hiển thị
    ```bs-alert warning flex
    <i class="fas fa-triangle-exclamation fa-xl me-2"></i>
    Hộp alert này sử dụng tông màu `warning` và có một [icon FontAwesome](https://fontawesome.com/icons/triangle-exclamation?f=classic&s=solid&sz=xl) ở bên trái.<br/>
    Nhớ đừng quên để riêng cái biểu tượng ở dòng dành cho tiêu đề (tức là dòng đầu tiên).
    ```
    ]]
```

Và dĩ nhiên là bạn cũng có thể dùng bộ [icon Bootstrap](https://icons.getbootstrap.com/):
```bs-tabs
    [[bs-tab Markdown
        ```bs-alert primary flex
        <i class="bi bi-info-circle me-2" style="font-size: 2rem;"></i>
        [GitHub Flavored Markdown](https://github.github.com/gfm/#what-is-github-flavored-markdown-), hay còn được viết tắt GFM, là 1 phương ngữ của Markdown được hỗ trợ để viết nội dung trên **GitHub.com** và **GitHub Enterprise**.
        ```
    ]]
    [[bs-tab Hiển thị
    ```bs-alert primary flex
    <i class="bi bi-info-circle me-2" style="font-size: 2rem;"></i>
    [GitHub Flavored Markdown](https://github.github.com/gfm/#what-is-github-flavored-markdown-), hay còn được viết tắt GFM, là 1 phương ngữ của Markdown được hỗ trợ để viết nội dung trên **GitHub.com** và **GitHub Enterprise**.
    ```
    ]]
```

## Bootstrap Card

Các [Bootstrap Card](https://getbootstrap.com/docs/5.0/components/card/) có thể được nhúng vào tài liệu bằng cú pháp sau:

<pre><code class="hljs language-markdown">
```bs-cards &lt;thông-số-tuỳ-biến-nhóm-card>
    [[bs-card &lt;thông-số-tuỳ-biến-card>
    -img:đường dẫn tới ảnh nằm trên phần header của card
    -header:header của card, phải nằm trọn trên dòng này
    -title:tiêu đề phần nội dung của card, phải nằm trọn trên dòng này
    -subtitle:tiêu đề phụ phần nội dung của card, phải nằm trọn trên dòng này
    -footer:footer của card, phải nằm trọn trên dòng này
    Các dòng còn lại là phần nội dung của card.
    Có thể sử dụng **Markdown** trong phần _nội dung_.
    ]]

    [[bs-card &lt;thông-số-tuỳ-biến-card>
    card tiếp theo
    ]]
```
</code></pre>

```bs-alert warning flex
<i class="fas fa-triangle-exclamation me-3 fa-2xl"></i>
Lưu ý **canh lề**! Phần mô tả của card (bắt đầu từ `[bs-card` đến `]]`) được canh lề bằng **4 dấu cách** so với phần mô tả của nhóm-card khoản (bắt đầu từ <code>\```bs-cards</code> đến <code>\```</code>).
```

**Thông số tuỳ biến nhóm-card** - tuỳ biến cách các card được hiển thị trong nhóm-card.
|Thông số|Mặc định|Mô tả|
|---|---|---|
|`equals`|`false`|Nếu được thiết lập, các card trong nhóm-card sẽ có cùng chiều cao.|
|`lightbox=<gallery-name>`| |Nếu được thiết lập, chế độ 'lightbox' sẽ được gắn vào các ảnh trên phần header của các card, và các ảnh này sẽ được nhóm chung trong một trưng bày có tên được chỉ định bởi `gallery-name`.|
|`cols=<num-columns>`| |Số lượng card trên 1 dòng (số trong khoảng `1-12`). Nếu không được thiết lập, các card sẽ dàn đều hết trên 1 dòng, tràn qua dòng tiếp theo, và cứ thế tiếp tục.|
|`cols-sm=<num-columns>`| |Số lượng card trên 1 dòng (số trong khoảng `1-12`) cho màn hình kích cỡ _nhỏ_. Nếu không được thiết lập, các card sẽ dàn đều hết trên 1 dòng, tràn qua dòng tiếp theo, và cứ thế tiếp tục.|
|`cols-md=<num-columns>`| |Số lượng card trên 1 dòng (số trong khoảng `1-12`) cho màn hình kích cỡ _trung bình_. Nếu không được thiết lập, các card sẽ dàn đều hết trên 1 dòng, tràn qua dòng tiếp theo, và cứ thế tiếp tục.|
|`cols-lg=<num-columns>`| |Số lượng card trên 1 dòng (số trong khoảng `1-12`) cho màn hình kích cỡ _lớn_. Nếu không được thiết lập, các card sẽ dàn đều hết trên 1 dòng, tràn qua dòng tiếp theo, và cứ thế tiếp tục.|

```bs-alert info

`cols`, `cols-sm`, `cols-md` và `cols-lg` có thể dùng chung với nhau để thiết lập số lượng card trên dòng cho các kích thước màn hình khác nhau.
```
```bs-alert warning

Thông số `lightbox` sẽ không có tác dụng nếu trong nhóm-card không có card nào chứa ảnh trên phần header.
Ảnh ở các nhóm-card khác nhau có thể nằm chung trong 1 trung bày (chung giá trị `gallery-name`).
```

**Các thành phần của card và thông số tuỳ biến**

- Tiền tố `-img:` chỉ định ảnh được gắn vào header của card. Đường dẫn ảnh phải được đặt trên cùng một dòng với tiền tố.
- Tiền tố `-header:` chỉ định nội dung header của card. Nội dung header phải được đặt trên cùng một dòng với tiền tố.
- Tiền tố `-footer:` chỉ định nội dung footer của card. Nội dung footer phải được đặt trên cùng một dòng với tiền tố.
- Tiền tố `-title:` chỉ định title cho phần nội dung của card. Title phải được đặt trên cùng một dòng với tiền tố.
- Tiền tố `-subtitle:` chỉ định title phụ cho phần nội dung của thẻ. Title phụ phải được đặt trên cùng một dòng với tiền tố.
- Các dòng khác là nội dung chính của card, có thể sử dụng Markdown.

Ví dụ:

```bs-tabs
    [[bs-tab Markdown
        ```bs-cards cols=3 lightbox=mygallery
            [[bs-card
            -img:https://placehold.co/320x200?text=Image+1
            -title:Title của card
            -subtitle:Title phụ của card
            Các dòng khác là nội dung chính của card.
            Có thể _sử dụng_ **Markdown**.
            Card này có title và title phụ, nhưng không có phần header và footer.

            Lưu ý: Ảnh trên header của thẻ này và thẻ kế tiếp được bật chế độ lightbox.
            ]]

            [[bs-card
            -img:https://placehold.co/320x200?text=Image+2
            -header:Header của card
            -footer:Footer của card nằm riêng
            Thẻ này có **header và footer** nhưng không có ~title và title phụ~.

            Lưu ý: Ảnh trên header của thẻ này và thẻ trước đó được bật chế độ lightbox.
            ]]

            [[bs-card
            -header:<i class="fas fa-heading"></i> **Header** có icon
            -title:<i class="fas fa-1"></i> _Title_ có icon
            -subtitle:<i class="fas fa-2"></i> ~Subtitle~ có icon
            -footer:<i class="fas fa-3"></i> `Footer` có icon
            Đây là card thứ 3. Lưu ý rằng _chiều cao_ của 3 card khác nhau.

            Chỉ có thể sử dụng **Markdown** trong phần nội dung chính của card.
            Cú pháp Markdown trong phần `header`, `footer`, `title` và `title phụ` sẽ không hoạt động.
            ]]
        ```
    ]]

    [[bs-tab Hiển thị
    ```bs-cards cols=3 lightbox=mygallery
        [[bs-card
        -img:https://placehold.co/320x200?text=Image+1
        -title:Title của card
        -subtitle:Title phụ của card
        Các dòng khác là nội dung chính của card.
        Có thể _sử dụng_ **Markdown**.
        Card này có title và title phụ, nhưng không có phần header và footer.

        Lưu ý: Ảnh trên header của thẻ này và thẻ kế tiếp được bật chế độ lightbox.
        ]]

        [[bs-card
        -img:https://placehold.co/320x200?text=Image+2
        -header:Header của card
        -footer:Footer của card nằm riêng
        Thẻ này có **header và footer** nhưng không có ~title và title phụ~.

        Lưu ý: Ảnh trên header của thẻ này và thẻ trước đó được bật chế độ lightbox.
        ]]

        [[bs-card
        -header:<i class="fas fa-heading"></i> **Header** có icon
        -title:<i class="fas fa-1"></i> _Title_ có icon
        -subtitle:<i class="fas fa-2"></i> ~Subtitle~ có icon
        -footer:<i class="fas fa-3"></i> `Footer` có icon
        Đây là card thứ 3. Lưu ý rằng _chiều cao_ của 3 card khác nhau.

        Chỉ có thể sử dụng **Markdown** trong phần nội dung chính của card.
        Cú pháp Markdown trong phần `header`, `footer`, `title` và `title phụ` sẽ không hoạt động.
        ]]
    ```
    ]]
```

Thông số tuỳ biến:

|Thông số|Mặc định|Mô tả|
|---|---|---|
|`no-muted`|`false`|Tiêu đề phụ và footer của card sẽ được gắn CSS `text-muted`. Nếu thiết lập thông số `no-muted`, CSS `text-muted` sẽ không được gắn vào.|
|`text=<style>`| |Tuỳ biến màu chữ. Giá trị hợp lệ: xem [màu chữ của Bootstrap](https://getbootstrap.com/docs/5.0/utilities/colors/#colors) (không tính phần tiền tố `text-`!).|
|`bg=<style>`| |Tuỳ biến màu nền. Giá trị hợp lệ: xem [màu nền của Bootstrap](https://getbootstrap.com/docs/5.0/utilities/background/#background-color) (không tính phần tiền tố `bg-`!).|
|`border=<style>`| |Tuỳ biến màu viền. Giá trị hợp lệ: xem [màu viền của Bootstrap](https://getbootstrap.com/docs/5.0/utilities/borders/#border-color) (không tính phần tiền tố `border-`!).|

```bs-alert info

Thông số `text`, `bg` và `border` có thể được dùng chung với nhau.
```

Ví dụ:

```bs-tabs
    [[bs-tab Markdown
        ```bs-cards cols-sm=1 cols-md=2 cols-lg=3 equals
            [[bs-card border=primary text=info
            -footer:Footer không "no-muted"
            Card này _tuỳ biến_ màu viền và màu chữ.
            ]]

            [[bs-card bg=success text=light no-muted
            -footer:Footer với "no-muted"
            Card này _tuỳ biến_ màu nền và màu chữ.

            Xem phần footer khi có thông số `no-muted`.
            ]]

            [[bs-card bg=secondary text=light no-muted
            -header:Header
            Card này _tuỳ biến_ màu nền và màu chữ.

            Header của card không bị ảnh hưởng bởi thông số `no-muted`.
            ]]

            [[bs-card bg=dark text=light
            -title:Title
            -subtitle:Title phụ không "no-muted"
            Card này _tuỳ biến_ màu nền và màu chữ.
            ]]

            [[bs-card bg=light text=dark no-muted
            -title:Title
            -subtitle:Title phụ với "no-muted"
            Card này _tuỳ biến_ màu nền và màu chữ.

            Xem phần title phụ khi có thông số `no-muted`.
            ]]

            [[bs-card border=danger bg=warning text=dark no-muted
            Tuỳ biến màu _viền_, _nền_ và _chữ_.

            Và lưu ý rằng toàn bộ các card trong nhóm-card này có cùng chiều cao.

            Và số lượng card mỗi dòng sẽ thay đổi tuỳ theo kích thước màn hình.
            ]]
        ```
    ]]

    [[bs-tab Render result
    ```bs-cards cols-sm=1 cols-md=2 cols-lg=3 equals
        [[bs-card border=primary text=info
        -footer:Footer không "no-muted"
        Card này _tuỳ biến_ màu viền và màu chữ.
        ]]

        [[bs-card bg=success text=light no-muted
        -footer:Footer với "no-muted"
        Card này _tuỳ biến_ màu nền và màu chữ.

        Xem phần footer khi có thông số `no-muted`.
        ]]

        [[bs-card bg=secondary text=light no-muted
        -header:Header
        Card này _tuỳ biến_ màu nền và màu chữ.

        Header của card không bị ảnh hưởng bởi thông số `no-muted`.
        ]]

        [[bs-card bg=dark text=light
        -title:Title
        -subtitle:Title phụ không "no-muted"
        Card này _tuỳ biến_ màu nền và màu chữ.
        ]]

        [[bs-card bg=light text=dark no-muted
        -title:Title
        -subtitle:Title phụ với "no-muted"
        Card này _tuỳ biến_ màu nền và màu chữ.

        Xem phần title phụ khi có thông số `no-muted`.
        ]]

        [[bs-card border=danger bg=warning text=dark no-muted
        Tuỳ biến màu _viền_, _nền_ và _chữ_.

        Và lưu ý rằng toàn bộ các card trong nhóm-card này có cùng chiều cao.

        Và số lượng card mỗi dòng sẽ thay đổi tuỳ theo kích thước màn hình.
        ]]
    ```
    ]]
```

## Bootstrap Tab

Một khối [Bootstrap tab](https://getbootstrap.com/docs/5.0/components/navs-tabs/) được nhúng vào tài liệu bằng cú pháp sau:

```markdown
    ```bs-tabs
    [[bs-tab Tên Tab1
    Nội dung của tab đầu tiên,
    có thể nằm trên nhiều dòng
    ]]

    [[bs-tab Tab thứ 2
    Và nội dung có thể sử dụng _Markdown_.
    ]]
    ```
```
Đoạn Markdown ở trên sẽ tạo ra kết quả như sau:
```bs-tabs
[[bs-tab Tên Tab1
Nội dung của tab đầu tiên,
có thể nằm trên nhiều dòng
]]

[[bs-tab Tab thứ 2
Và nội dung có thể sử dụng _Markdown_.
]]
```

Mặc định các tab được sắp xếp theo hàng ngang, nhưng chúng cũng có thể nằm theo hàng dọc. Ví dụ đoạn Markdown như sau:
```markdown
    ```bs-tabs vertical
    [[bs-tab Tab1
    Nội dung của tab 1.
    ]]

    [[bs-tab Tab2
    Nội dung của tab 2.
    ]]
    ```
```
sẽ tạo ra kết quả như sau:
```bs-tabs vertical
[[bs-tab Tab1
Nội dung của tab 1.
]]

[[bs-tab Tab2
Nội dung của tab 2.
]]
```

Xem thêm:
- [Cú pháp Markdown được hỗ trợ](../markdown/)
