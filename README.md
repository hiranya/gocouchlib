# gocouchlib

CouchDB Client for Go


## Documentation
GoDoc for this project can be found at <http://godoc.org/github.com/hiranya/gocouchlib>

## Installation
To install the latest version of gocouchlib
```bash
go get github.com/hiranya/gocouchlib
```

## Examples

### Example-1
```go
package main

import (
	"fmt"
	"github.com/hiranya/gocouchlib"
	"net/url"
)

func main() {
	s := &gocouchlib.Server{"http://localhost:5984",
		url.UserPassword("testuser", "password"),
	}

	db := &gocouchlib.Database{"gocouch", s}

	isExist, _ := db.Exists()
	fmt.Println("DB Exists:", isExist)
}

```

## Development
Want to contribute? Great! Please use gitflow workflow to submit your features or patches <http://nvie.com/posts/a-successful-git-branching-model/>


## License
Copyright (C) 2015  Hiranya Samarasekera <hiranyas@gmail.com>

This software is licensed under GNU GPL v3.0 http://www.gnu.org/licenses/
