package pkg

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
)

// list of clients connected
var client []net.Conn
var path, _ = os.UserHomeDir()
var f, _ = os.OpenFile(path+"/.tochat", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0600)

// remove element from array
func remove(slice []net.Conn, s int) []net.Conn {
	return append(slice[:s], slice[s+1:]...)
}

// show local ip
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func handlec(conn net.Conn) {
	// defer close conn
	defer func() {
		//fmt.Println(conn, "closed...")
		conn.Close()
	}()
	var e net.Conn
	scan := bufio.NewReader(conn)
	for {
		cmsg, err := scan.ReadString('\n')
		cmsg = encrypt(cmsg)
		if err != nil {
			print(err)
			break
		}
		if strings.Contains(cmsg, "-1") { // remove client from list
			//fmt.Println(conn, "closed...")
			for i, e := range client {
				if e == conn {
					remove(client, i)
				}
			}
			conn.Close()
		} else {
			fmt.Fprintf(f, "%s", cmsg)
			cmsg = encrypt(cmsg)
			for _, e = range client { // send the message to all other client
				if e == conn {

				} else {
					e.Write([]byte(cmsg))
				}
			}
		}
	}
}

// show all messages sent
func show_msg() {
	s := bufio.NewScanner(os.Stdin)
	var m []byte
	for {
		Clear()
		m, _ = ioutil.ReadFile(string(path) + "/.tochat")
		fmt.Println(string(m))
		fmt.Print("(h for help)> ")
		s.Scan()
		if s.Text() == "-1" {
			break
		} else if s.Text() == "h" {
			Clear()
			fmt.Print(`Usage:
  -1 to quit to command
	anything else to skip
	`)
			fmt.Print("Press Enter to continue...")
			fmt.Scanln()
		}
	}
}

// show clients and allow to kick them from the server
func sclose(l net.Listener) {
	scan := bufio.NewScanner(os.Stdin)
	km := encrypt("Server is now down...\n")
	var n string
	for {
		for i, e := range client { // list all clients
			fmt.Println(i+1, e.RemoteAddr())
		}
		fmt.Print("to kick> ")
		scan.Scan()
		n = scan.Text()
		n, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("Number not on the list...")
			break
		}
		if n == -1 { // quit the menu
			break
		} else if n > len(client) { // verify is the client number is ok
			fmt.Println("Number not on the list...")
		} else if n < 1 {
			fmt.Println("Number not on the list...")
		} else {
			for i, e := range client {
				if i+1 == n {
					e.Write([]byte(km)) // send the kick message
					client = remove(client, i)
					e.Close()
				}
			}
		}
	}
}

func accept(l net.Listener) {
	for { //infinite loop handle connection
		conn, err := l.Accept()
		if err != nil {
			return
		}
		client = append(client, conn)
		go handlec(conn) //handle client
	}
}

// start server
func Serv(port string) {
	l, err := net.Listen("tcp", ":"+port) // create listener
	if err != nil {
		print(err)
		os.Exit(3)
	}
	defer func() { // defer to close the server
		l.Close()
		//fmt.Println("Connection closed...")
		os.Exit(0)
	}()
	Clear()
	scan := bufio.NewScanner(os.Stdin)
	go accept(l)
	for {
		Clear()
		fmt.Print("Command (h for help)> ")
		scan.Scan()
		if scan.Text() == "h" { // show the help
			Clear()
			fmt.Print(`Usage:
	l to list and kill client
	m to see messages
	ip to see which is the server ip and port
	-1 to stop the server or stop the command`)
			fmt.Print("\nPress Enter to continue...")
			fmt.Scanln()
		} else if scan.Text() == "l" { // show clients
			Clear()
			sclose(l)
		} else if scan.Text() == "m" { // show messages
			show_msg()
		} else if scan.Text() == "ip" { //show server ip and port
			Clear()
			fmt.Println("[+]Your ip is", GetLocalIP())
			fmt.Println("[+]Listening on port", port+"...")
			fmt.Print("Press Enter to continue...")
			fmt.Scanln()
		} else if scan.Text() == "-1" { // close the server
			f.Close()
			os.Remove(path + "/.tochat")
			for _, e := range client {
				e.Write([]byte("Server is now down...\n"))
				l.Close()
				os.Exit(0)
			}
			break
		}
	}

}
