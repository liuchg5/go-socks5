package main

import (
     "fmt"
     //"encoding/binary"
     //"io"
     "log"
     //"net"
     "os"
     //"testing"
     //"time"
     "socks5"
	"flag"
)
var pass = flag.String("pass", "", "password, username:foo")
var addr = flag.String("addr", "", "listen addr")

func main() {
	flag.Parse()
     fmt.Println("begin...")
     // Create a socks server
     creds := socks5.StaticCredentials{
          "foo": *pass, 
     }
     cator := socks5.UserPassAuthenticator{Credentials: creds}
     conf := &socks5.Config{
          AuthMethods: []socks5.Authenticator{cator},
          Logger:      log.New(os.Stdout, "", log.LstdFlags),
     }
     serv, err := socks5.New(conf)
     if err != nil {
          panic("err: " + err.Error())
     }

     // Start listening
     //go func() {
          if err := serv.ListenAndServe("tcp", *addr); err != nil {
               panic("err: " + err.Error())
          }
     //}()
    
     fmt.Println("end!!!")
}

