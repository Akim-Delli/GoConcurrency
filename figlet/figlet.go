package main

import (
	"fmt"
	"github.com/getwe/figlet4go"
)

func main() {
	ascii := figlet4go.NewAsciiRender()
	// most simple Usage
	renderStr, _ := ascii.Render("FeedTheBeast")
	fmt.Println(renderStr)
}
