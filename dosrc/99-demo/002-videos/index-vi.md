**Nhúng video với cú pháp <code>```video</code>**

Video này từ bên ngoài:
```video https://placehold.co/320x200/100090/ffffff.mp4?font=roboto&text=Hello,+world ratio=16x9
```

Video này đính kèm theo tài liệu, _giới hạn độ rộng_ và _canh giữa_:
```video embedded-video-320x200-1.mp4 ratio=21x9 width=440 center
```

**Video cũng có thể được đặt bên trong Bootstrap card**
```bs-cards equals cols-md=2
    [[bs-card
    -footer: Video này từ bên ngoài
    ```video https://placehold.co/320x200/900010/ffffff.mp4?font=lora&text=Hello,+world ratio=16x9
    ```
    ]]
    [[bs-card
    -footer: Video này đính kèm theo tài liệu
    ```video embedded-video-320x200-1.mp4 ratio=21x9
    ```
    ]]
```

**Hỗ trợ nhúng video từ Youtube**

```bs-cards
    [[bs-card
    -subtitle: Video hiển thị với tỉ lệ 16x9, đặt lại thời gian bắt đầu
    ```video https://www.youtube.com/watch?v=YnopHCL1Jk8&t=45 center ratio=16x9 width=480
    ```
    ]]
```

```bs-cards
    [[bs-card
    -subtitle: Video hiển thị với tỉ lệ 21x9
    ```video https://www.youtube.com/embed/lUbiDYhGpHQ center ratio=21x9 width=480
    ```
    ]]
```
