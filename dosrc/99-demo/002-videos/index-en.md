**Embed videos with <code>```video</code>**

This video is from external source:
```video https://placehold.co/320x200/100090/ffffff.mp4?font=roboto&text=Hello,+world ratio=16x9
```

This video is attached to the document, with _max width_ and aligned to the _center_:
```video embedded-video-320x200-1.mp4 ratio=21x9 width=440 center
```

**Videos can also be placed inside Bootstrap cards**
```bs-cards equals cols-md=2
    [[bs-card
    -footer: This video is from external source
    ```video https://placehold.co/320x200/900010/ffffff.mp4?font=lora&text=Hello,+world ratio=16x9
    ```
    ]]
    [[bs-card
    -footer: This video is attached to the document
    ```video embedded-video-320x200-1.mp4 ratio=21x9
    ```
    ]]
```

**Videos can also be Youtube videos**

```bs-cards
    [[bs-card
    -subtitle: Video with ratio 16x9, with custom start time
    ```video https://www.youtube.com/watch?v=YnopHCL1Jk8&t=45 center ratio=16x9 width=480
    ```
    ]]
```

```bs-cards
    [[bs-card
    -subtitle: Video with ratio 21x9
    ```video https://www.youtube.com/embed/lUbiDYhGpHQ center ratio=21x9 width=480
    ```
    ]]
```
