package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"tochat/pkg"
)

var vers string

// show help
func help() {
	pkg.Clear()
	fmt.Println(`Usage :
  -1 to quit
  s to start as server
  c to start as client`)
	fmt.Print("\nVersion : ", vers, "\n\n")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Text() != "-1" {
		pkg.Clear() // clear the screen
		fmt.Println("To chat (h for help):")
		fmt.Print("> ")
		scan.Scan()
		if scan.Text() == "h" {
			help() // show help
		} else if scan.Text() == "s" { // enter server mod
			fmt.Print("port> ")
			scan.Scan()
			pkg.Serv(scan.Text())
		} else if scan.Text() == "c" { // enter client mod
			fmt.Print("ip> ")
			scan.Scan()
			ip := scan.Text()
			fmt.Print("port> ")
			scan.Scan()
			pkg.Connect(ip, scan.Text())
		} else if scan.Text() == "-1" { // quit
			os.Exit(0)
		} else {
			fmt.Printf("%s: Not a good command...", scan.Text())
			time.Sleep(1 * time.Second)
		}
	}
}
