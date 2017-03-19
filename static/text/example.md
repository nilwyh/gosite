Installation
------------

Blackfriday is compatible with Go 1. If you are using an older
release of Go, consider using v1.1 of blackfriday, which was based
on the last stable release of Go prior to Go 1. You can find it as a
tagged commit on github.

With Go 1 and git installed:

    go get github.com/russross/blackfriday

will download, compile, and install the package into your `$GOPATH`
directory hierarchy. Alternatively, you can achieve the same if you
import it into a project:

    import "github.com/russross/blackfriday"

and `go get` without parameters.

Usage
-----

For basic usage, it is as simple as getting your input into a byte
slice and calling:

    output := blackfriday.MarkdownBasic(input)

This renders it with no extensions enabled. To get a more useful
feature set, use this instead:

    output := blackfriday.MarkdownCommon(input)

### Sanitize untrusted content

Blackfriday itself does nothing to protect against malicious content. If you are
dealing with user-supplied markdown, we recommend running blackfriday's output
through HTML sanitizer such as
[Bluemonday](https://github.com/microcosm-cc/bluemonday).

Here's an example of simple usage of blackfriday together with bluemonday:

``` go
import (
    "github.com/microcosm-cc/bluemonday"
    "github.com/russross/blackfriday"
)

// ...
unsafe := blackfriday.MarkdownCommon(input)
html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
```