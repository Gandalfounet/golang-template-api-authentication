// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"net/smtp"
)

// smtpServer data to smtp server
type smtpServer struct {
 host string
 port string
}
// serverName URI to smtp server
func (s *smtpServer) Address() string {
 return s.host + ":" + s.port
}

func Send(msg string) {
    // Sender data.
    from := "dummygandalf34@gmail.com"
    password := "Lescarottessontcuites31320!"
    // Receiver email address.
    to := []string{
        "dummygandalf34@gmail.com",
        "tariq.riahi@gmail.com",
    }
    // smtp server configuration.
    smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
    // Message.
    message := []byte(msg)
    // Authentication.
    auth := smtp.PlainAuth("", from, password, smtpServer.host)
    // Sending email.
    err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Email Sent!")
}


