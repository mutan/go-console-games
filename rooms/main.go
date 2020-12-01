package main

import "fmt"

func main() {
	printHeading()
}

func printHeading() {
	fmt.Print(`
 ____
|  _ \    ___     ___    _ __ ___    ___
| |_) |  / _ \   / _ \  | '_ ' _ \  / __|
|  _ <  | (_) | | (_) | | | | | | | \__ \
|_| \_\  \___/   \___/  |_| |_| |_| |___/
`)

	fmt.Println(`
                Rooms v1.0                
  Horror game by Akim Gubanov (c) 2020`)

	return
}
