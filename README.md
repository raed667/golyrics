# GoLyrics
Get the lyrics of any song with Go lang
=== 


## Installation

```bash
go get -u github.com/RaedsLab/golyrics
```

## Example 

```go
package main

import (
	"fmt"

	"github.com/raedslab/golyrics"
)

func main() {
	fmt.Printf(golyrics.Fetch("Queen","I Want To Break Free"))
}
```



