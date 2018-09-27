package main

import "fmt"
import "os"
import "net"
import "bufio"
import "strconv"

func main() {
	var ip string
	var port int
	if len(os.Args) > 1 {
		ip = os.Args[1]
		port, _ = strconv.Atoi(os.Args[2])
		proto := "tcp"
		version := "4"
		err := nc(ip, port, proto, version)
		_ = err
	} else {
		fmt.Fprintf(os.Stderr, "usage ./%s [ip] [port]", os.Args[0])
		os.Exit(1)
	}

}

func nc(ip string, port int, proto string, version string) error {
	if version != "4" {
		return fmt.Errorf("only implemented version tcp/ip version 4")
	} else if proto != "tcp" {
		return fmt.Errorf("only implemented tcp protocol")
	} else {
		conn, _ := net.Dial("tcp4", ip+":"+strconv.Itoa(port))
		for {
			in, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			fmt.Fprintf(conn, in+"\n")
			out, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Fprintf(os.Stdout, "%s", out)

		}

	}

	return nil

}
