# Windows color support for [CoLog](https://github.com/comail/colog)

Uses [go-colorable](https://github.com/mattn/go-colorable) to display unix terminal colors on Windows.
Behaviour for other systems is not altered.

![wincolog demo](http://i.imgur.com/Jelu7by.png)

```go
package main

import (
	"log"

	"comail.io/go/colog"
	"comail.io/go/wincolog"
)

func main() {
	colog.Register()
	colog.SetOutput(wincolog.Stdout())
	log.Println("trace: Trace should be plain")
	log.Println("debug: Debug should be blue")
	log.Println("info: Info should be green")
	log.Println("warn: Warning should be yellow")
	log.Println("error: Error should be red")
	log.Println("alert: Alert should be white on re")
}
```
