package pkg

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handle(conn net.Conn) {
	scan := bufio.NewReader(conn)
	for {
		msg, err := scan.ReadString('\n') // recv
		msg = encrypt(msg)
		if err != nil {
			print(err)
			break
		} else if msg == "Server is now down...\n" {
			fmt.Printf("\rServer> %s\n", msg)
			conn.Close()
			os.Exit(0)
		}
		fmt.Print("\r" + msg)
		fmt.Print("> ")
	}
}

func Connect(ip string, port string) {
	fmt.Print("[+]Connecting to ", ip+"...")
	conn, err := net.Dial("tcp", ip+":"+port) //connect to the server
	defer func() {
		fmt.Println("connection closed...")
		conn.Close()
	}()
	if err != nil {
		println("Server not found...")
		os.Exit(3)
	}
	Clear()
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("\nnickname> ")
	s.Scan()
	nick := s.Text()
	fmt.Println()
	for {
		go handle(conn)
		scan := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		msg, err := scan.ReadString('\n')
		if err != nil {
			print(err)
			break
		} else if strings.Contains(msg, "-1") {
			fmt.Println("Connection closed...")
			conn.Close()
			break
		}
		msg = nick + "> " + msg
		msg = encrypt(msg)
		_, err = fmt.Fprintf(conn, "%s", msg)
		if err != nil {
			fmt.Println("Connection lost...")
			conn.Close()
			os.Exit(0)
		}
	}
}
