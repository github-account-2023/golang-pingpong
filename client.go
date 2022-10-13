package main

import (
	"bufio"
	"log"
	"net"
)

func tcpClient(host string, port string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", host+":"+port)
	if err != nil {
		log.Println(err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Println("TCP connection failed:", err.Error())
		return
	}
	_, err = conn.Write([]byte("Hello?"))
	if err != nil {
		log.Println(err)
		return
	}
	response := make([]byte, 1024)
	length, err := conn.Read(response)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("TCP ->", string(response[:length]))
	err = conn.Close()
	if err != nil {
		log.Println(err)
		return
	}
}

func udpClient(host string, port string) {
	buf := make([]byte, 2048)
	conn, err := net.Dial("udp", host+":"+port)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = conn.Write([]byte("udp"))
	if err != nil {
		return
	}
	if err != nil {
		log.Println(err)
		return
	}
	length, err := bufio.NewReader(conn).Read(buf)
	if err == nil {
		log.Println("UDP ->", string(buf[:length]))
	} else {
		log.Println(err)
	}
	err = conn.Close()
	if err != nil {
		return
	}
}
