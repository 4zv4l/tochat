package main

import (
	"./pkg"
	"bufio"
	"fmt"
	"os"
	"time"
)

func help() {
	pkg.Clear()
	fmt.Println(`Usage :
  -1 to quit
  s to start as server
  c to start as client`)
	fmt.Print("Press Enter to continue...")
	fmt.Scanln()
}

func main() {
	for {
		pkg.Clear()
		scan := bufio.NewScanner(os.Stdin)
		fmt.Println("To chat (h for help):")
		fmt.Print("> ")
		scan.Scan()
		if scan.Text() == "h" {
			help()
		} else if scan.Text() == "s" {
			fmt.Print("port> ")
			scan.Scan()
			pkg.Serv(scan.Text())
		} else if scan.Text() == "c" {
			fmt.Print("ip> ")
			scan.Scan()
			ip := scan.Text()
			fmt.Print("port> ")
			scan.Scan()
			pkg.Connect(ip, scan.Text())
		} else if scan.Text() == "-1" {
			os.Exit(0)
		} else {
			print("Not a good parameter...")
			time.Sleep(2 * time.Second)
		}
	}
}
