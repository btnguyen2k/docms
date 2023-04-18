```bs-alert warning flex
<i class="fas fa-triangle-exclamation fa-2xl me-2"> </i>
Các cú pháp Markdown mô tả trong tài liệu này áp dụng cho các giao diện đồ hoạ đi kèm với phiên bản `DO CMS runtime` gốc. Nếu bạn sử dụng một bộ giao diện đồ hoạ của bên thứ 3, vui lòng kiểm tra tài liệu đi kèm.
```

Vì Markdown cũng hỗ trợ HTML, bạn có thể nhúng các [giao diện Bootstrap](https://getbootstrap.com/docs/5.0/components/) vào tài liệu thông qua HTML/CSS. Tuy nhiên, sẽ tiện hơn nếu có cú pháp Markdown ngắn gọn hơn, dễ nhớ và dễ sử dụng hơn để làm chuyện đó.

`DO CMS runtime` hỗ trợ các giao diện Bootstrap sau:

## Bootstrap Alerts

Nhúng [Bootstrap Alerts](https://getbootstrap.com/docs/5.0/components/alerts/) vào tài liệu thông qua cú pháp sau:
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
    [[bs-tab Rendered result
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
    [[bs-tab Rendered result
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
    [[bs-tab Rendered result
    ```bs-alert primary flex
    <i class="bi bi-info-circle me-2" style="font-size: 2rem;"></i>
    [GitHub Flavored Markdown](https://github.github.com/gfm/#what-is-github-flavored-markdown-), hay còn được viết tắt GFM, là 1 phương ngữ của Markdown được hỗ trợ để viết nội dung trên **GitHub.com** và **GitHub Enterprise**.
    ```
    ]]
```

## Bootstrap Tabs

Một nhóm [Bootstrap tabs](https://getbootstrap.com/docs/5.0/components/navs-tabs/) được nhúng vào tài liệu bằng cú pháp sau:

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
