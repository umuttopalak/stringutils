# stringutils

stringutils is a collection of small utility tools for `string` type.

## Installation
```bash
go get -u github.com/umuttopalak/stringutils
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/umuttopalak/stringutils-demo"
)

func main(){
	reversed, err := stringutils.Reverse("vigo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reversed) // ogiv
}
```

## Contributor(s)

* [Umut Topalak](https://github.com/umuttopalak) - Creator, maintainer


## Contribute

All PR’s are welcome!

1. `fork` (https://github.com/umuttopalak/stringutils/fork)
1. Create your `branch` (`git checkout -b my-feature`)
1. `commit` yours (`git commit -am 'add some functionality'`)
1. `push` your `branch` (`git push origin my-feature`)
1. Than create a new **Pull Request**!

