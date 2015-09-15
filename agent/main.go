package main

import (
	"log"
	"net"
	"flag"
)


var listenAddr = flag.String("listenAddr", "", "listen for connect")
var remoteAddr = flag.String("remoteAddr", "", "peer ip or proxy ip")


func main() {
	flag.Parse()
	// Listen on TCP port 2000 on all interfaces.
	l, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		localConn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		
		remoteConn, err := net.Dial("tcp", *remoteAddr)
		if err != nil {
			log.Println("remote dial failed!!!", err)
			localConn.Close()
			continue
		}
		go worker(localConn, remoteConn)
	}
}

func worker(localConn net.Conn, remoteConn net.Conn) {
	go rd2wt(localConn, remoteConn)
	go rd2wt(remoteConn, localConn)
}

func rd2wt(rd net.Conn, wt net.Conn) {
	b := make([]byte, 1)
	var err error = nil
	for {
		_, err = rd.Read(b)
		if err != nil {
			log.Println("Read failed! err=", err)
			rd.Close()
			wt.Close()
			return
		}
		_, err = wt.Write(b)
		if err != nil {
			log.Println("Write failed! err=", err)
			rd.Close()
			wt.Close()
			return
		}
	}
}



