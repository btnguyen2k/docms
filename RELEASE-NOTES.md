# DOCMS Release Notes

## 2023-07-x

Runtime `v0.4.x`:
- Template `blogy`: update list of special pages to `['about', 'contact', 'disclaimer', 'privacy', 'privacy-policy']`.
- Bug fixes & enhancements.

## 2023-06-29

Runtime `v0.4.0.1`:
- Bug fixes.

## 2023-06-28

Runtime `v0.4.0`:
- New backend API to return feeds.
- Extend media MINE types to support attached audios, videos, documents, and more.
- Bug fixes & enhancements.
- Support videos:
  - With tag <code>```video</code>
  - Videos from Youtube

## 2023-05-11

Runtime `v0.3.1`:
- Add Google Analytics support.
- Add site mode `blog`.
- New FE template `Blogy`.
- Add diagram support using [mermaid](https://mermaid.js.org/).
- Bug fixes & enhancements.

CLI `v0.3.1.3`:
- Bug fixes & enhancements.

## 2023-05-01

CLI `v0.3.1.1`:
- Command `touch` to update document's timestamp.
- Minor enhancement.

CLI `v0.3.1`:
- In-sync with runtime `v0.3.1` (site mode/author, page special purpose/style and author).

## 2023-04-25

Runtime `v0.3.0`:
- Support tags at site metadata level
- Support tags at document metadata level
- Render site's tags with syntax <code>[[do-tag</code>
- Support [Bootstrap components](https://getbootstrap.com/docs/5.0/components/):
  - `Cards` with syntax <code>```bs-cards</code>, support lightbox.

CLI `v0.3.0.1`:
- Pre-process tags, to be in-sync with the runtime.

## 2023-04-18

Runtime `v0.2.0`:
- Support [KaTeX](https://katex.org/) block with syntax <code>```katex_</code>
- Links to external websites will open new windows/tabs.
- Embed GitHub Gist with syntax <code>```gh-gist</code>
- Support [Bootstrap components](https://getbootstrap.com/docs/5.0/components/):
  - `Alerts` with syntax <code>```bs-alert</code>
  - `Tabs` with syntax <code>```bs-tabs</code>

## 2023-04-13

CLI `v0.1.1`:
- Command `new` - helper to create assets with default metadata.
  - `new site`: Create new site content metadata
  - `new topic`: Create new topic metadata
  - `new document`: Create new document metadata

Runtime `v0.1.1`:
- Bug fixes.

## 2023-04-12

CLI `v0.1.0`:
- Pre-process website content with command `docli build`.
- Build fulltext index while pre-processing website content.

Runtime `v0.1.0`:
- Render website content, supporting 3 themes: `bootstrap` (default), `coderdocs` and `prettydocs`.
- Fulltext search.
- Multi-languages support.
