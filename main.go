package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 4 {
		mode := args[1]
		host := args[2]
		port := args[3]
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		if mode == "server" {
			go tcpServer(host, port)
			udpServer(host, port)
		} else if mode == "client" {
			tcpClient(host, port)
			udpClient(host, port)
		} else {
			log.Println("Unsupported mode! Please use server or client")
			os.Exit(3)
		}
	} else {
		log.Println("Usage:", args[0], "mode ip port")
		//host := "127.0.0.1"
		//port := "8964"
		//go tcpServer(host, port)
		//go udpServer(host, port)
		//tcpClient(host, port)
		//udpClient(host, port)
	}
}
