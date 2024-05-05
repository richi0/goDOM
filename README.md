# goDOM [![GoDoc](https://pkg.go.dev/badge/goDOM.svg)](https://pkg.go.dev/goDOM)

goDOM parses HTML documents and makes data extraction easy.

## What does goDOM do?

goDOM provides methodes to work with html documents in a similar way as the Javascript Document interface.

After parsing a document many of the well known DOM methods can be used to find and filter nodes inside the document.

## How do I use goDOM?

### Install

```
go get -u github.com/...
```

### Example 1

Get all the urls from a html document.

```go
import (
	"fmt"
	"goDOM"
	"net/http"
)

func main() {
	// fetch data from a website
	url := "https://en.wikipedia.org/wiki/Go_(programming_language)"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// parse website content
	dom, err := goDOM.New(resp.Body)
	links := dom.GetElementsByTagName("a")

	// print urls
	for _, link := range links {
		fmt.Println(link.Attributes()["href"])
	}

	// Print output:
	// /wiki/Main_Page
	// /wiki/Wikipedia:Contents
	// /wiki/Portal:Current_events
	// /wiki/Special:Random
	// /wiki/Wikipedia:About
	// ...
}

```

### Example 2

Print the title of a html document.

```go
import (
	"fmt"
	"goDOM"
	"net/http"
)

func main() {
	// fetch data from a website
	url := "https://en.wikipedia.org/wiki/Go_(programming_language)"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// parse website content
	dom, err := goDOM.New(resp.Body)
	titleElement := dom.GetElementsByTagName("title")[0]

	// print title
	fmt.Println(titleElement.Text(false))

	// Print output:
	// Go (programming language) - Wikipedia
}
```

### Example 3

Print the text of a specific node in html document.

```go
import (
	"fmt"
	"goDOM"
	"net/http"
)

func main() {
	// fetch data from a website
	url := "https://en.wikipedia.org/wiki/Go_(programming_language)"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// parse website content
	dom, err := goDOM.New(resp.Body)
	historyElement := dom.GetElementById("History")
	paragraph := historyElement.Parent().NextElementSibling()

	// print paragraph
	fmt.Println(paragraph.Text(true))

	// Print output:
	// Go was designed at Google in 2007 to improve programming productivity in...
}
```

## Documentation

Find the full documentation of the package here: https://pkg.go.dev/github.com/richi0/goDOM
