**Markdown support image natively**

![An external image](https://placehold.co/320x200/100090/ffffff?font=roboto&text=Hello,+world "An external image")

![An attached image](embedded-img-320x200-1.jpg "An attached image")

**Images can also be attached to Bootstrap cards**

```bs-cards lightbox=mygallery cols-md=2 equals
    [[bs-card
    -img:https://placehold.co/320x200/009010/ffffff?font=oswald&text=Hello,+world
    -footer: This image is from external source
    ]]
    [[bs-card
    -img:embedded-img-320x200-2.jpg
    -footer: This image is attached to the document
    ]]
```

**Can also be placed inside table cells**
|   |   |
|:-:|:-:|
|This image is from external source<br/>![An external image](https://placehold.co/320x200/901000/ffffff?font=playfair-display&text=Hello,+world "An external image")|This image is attached to the document<br/>![An attached image](embedded-img-320x200-3.jpg "An attached image")|
