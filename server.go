package main

import (
	"log"
	"net"
	"os"
	"strconv"
)

func sendUDPResponse(conn *net.UDPConn, remoteAddr *net.UDPAddr) {
	log.Println("Got UDP from", remoteAddr)
	_, err := conn.WriteToUDP([]byte(remoteAddr.String()), remoteAddr)
	if err != nil {
		log.Println(err)
	}
}

func udpServer(host string, port string) {
	udpPort, err := strconv.Atoi(port)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: udpPort,
		IP:   net.ParseIP(host),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("UDP server is listening on " + host + ":" + port)
	for {
		_, remoteAddr, err := ser.ReadFromUDP(p)
		if err != nil {
			log.Println(err)
			continue
		}
		go sendUDPResponse(ser, remoteAddr)
	}
}

func sendTCPResponse(conn net.Conn) {
	remoteAddr := conn.RemoteAddr()
	log.Println("Got TCP from", remoteAddr.String())
	_, err := conn.Write([]byte(remoteAddr.String()))
	if err != nil {
		log.Println(err)
		return
	}
}

func tcpServer(host string, port string) {
	l, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(l)
	log.Println("TCP server is listening on " + host + ":" + port)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go sendTCPResponse(conn)
	}
}
