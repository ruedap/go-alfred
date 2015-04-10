# go-alfred [![Build Status](http://img.shields.io/travis/ruedap/go-alfred.svg?style=flat-square)](https://travis-ci.org/ruedap/go-alfred) [![Coverage Status](http://img.shields.io/coveralls/ruedap/go-alfred/master.svg?style=flat-square)](https://coveralls.io/r/ruedap/go-alfred)

Alfred workflow utility library in Golang


## Installation

This package can be installed with the go get command:

```
go get github.com/ruedap/go-alfred
```

## Usage

``` go
package main

import (
	"fmt"

	"github.com/ruedap/go-alfred"
)

func main() {
	resp := alfred.NewResponse()
	item := alfred.ResponseItem{
		Valid:    true,
		UID:      "uid-foo",
		Title:    "title-foo",
		Subtitle: "Subtitle foo.",
		Arg:      "arg-foo",
		Icon:     "icon-foo.png",
	}
	resp.AddItem(&item)

	xml, err := resp.ToXML()
	if err != nil {
		title := fmt.Sprintf("Error: %v", err.Error())
		subtitle := "Foo Workflow Error"
        arg := title
		errXML := alfred.ErrorXML(title, subtitle, arg)
		fmt.Println(errXML)
		// <?xml version="1.0" encoding="UTF-8"?>
		// <items><item valid="false" arg="Error: xxx" uid="error"><title>Error: xxx</title><subtitle>Foo Workflow Error</subtitle><icon>/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertStopIcon.icns</icon></item></items>
		return
	}

	fmt.Println(xml)
	// <?xml version="1.0" encoding="UTF-8"?>
	// <items><item valid="true" arg="arg-foo" uid="uid-foo"><title>title-foo</title><subtitle>Subtitle foo.</subtitle><icon>icon-foo.png</icon></item></items>
}
```

## Example

* [alfred-emma-css-workflow](https://github.com/ruedap/alfred-emma-css-workflow)


## License

Released under the [MIT license](http://ruedap.mit-license.org/2015).


## Author

[ruedap](https://github.com/ruedap)
