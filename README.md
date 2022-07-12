# SaveMap
Improved sync.map for more return values

## Installation
1. You can use the below Go command to install SafeMap.

```sh
$ go get -u github.com/xcocx/safemap
```

2. Import it in your code:

```go
import "github.com/xcocx/safemap"
```

## Example
```go
package main

import (
	"fmt"
	"github.com/xcocx/safemap"
)

func main()  {
	m := safemap.New()
	m.Set("key", "value")
	
	fmt.Println(m.Get("key")) // value, true
	fmt.Println(m.GetBool("key")) // false
	fmt.Println(m.GetString("key")) // value
	fmt.Println(m.GetInt("key")) // 0
}
```