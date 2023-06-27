**Markdown hỗ trợ hiển thị hình ảnh**

![An external image](https://placehold.co/320x200/100090/ffffff?font=roboto&text=Hello,+world "Hình này từ bên ngoài")

![An attached image](embedded-img-320x200-1.jpg "Hình này đính kèm với bài viết")

**Hình ảnh có thể được gắn vào trong Bootstrap card**

```bs-cards lightbox=mygallery cols=2 equals
    [[bs-card
    -img:https://placehold.co/320x200/009010/ffffff?font=oswald&text=Hello,+world
    -footer: Hình này từ bên ngoài
    ]]
    [[bs-card
    -img:embedded-img-320x200-2.jpg
    -footer: Hình này đính kèm với bài viết
    ]]
```

**Cũng có thể bỏ vào bên trong ô của bảng**
|   |   |
|:-:|:-:|
|Hình này từ bên ngoài<br/>![An external image](https://placehold.co/320x200/901000/ffffff?font=playfair-display&text=Hello,+world "Hình này từ bên ngoài")|Hình này đính kèm với bài viết<br/>![An attached image](embedded-img-320x200-3.jpg "Hình này đính kèm với bài viết")|
