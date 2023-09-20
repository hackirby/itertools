# itertools [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/hackirby/itertools)

Easily iterate through any slice

## Install

```bash
go get github.com/hackirby/itertools
```

## Example

```go
import (
    "fmt"
    "github.com/hackirby/itertools"
)

func main() {
    // Create a new iterator from any slice
    iterator, _ := itertools.Iter([]any{1, "two", 3, "four", 5})

    // Some infos about the iterator
    fmt.Println("Length:", iterator.Len())
    fmt.Println("Index:", iterator.Index())
    fmt.Println("Current item:", iterator.Current()) // nil, call Next() first

    // use itertools.Cycle() to create an infinite iterator
    fmt.Println("Is Cycle ?", iterator.IsCycle())

    // set the iterator to the 2nd element (0-based)
    iterator.SetIndex(1)
    fmt.Println(iterator.Current()) // "two"

    // reset the iterator
    iterator.Reset()
    fmt.Println(iterator.Current()) // nil, call Next() first

    // iterate over the iterator
    for iterator.HasNext() {
	switch item := iterator.Next().(type) {
	case int:
            fmt.Println("int:", item) // can be used as int in this block
	case string:
            fmt.Println("string:", item) // can be used as string in this block
	}
    }

    // iterate over the iterator in reverse order
    for iterator.HasPrev() {
	fmt.Println(iterator.Prev())
    }
}
```
