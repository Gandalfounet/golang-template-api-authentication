// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
    "os"
    "github.com/joho/godotenv"
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
    err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Sender data.
    from := os.Getenv("email_smtp_user")
    password := os.Getenv("email_smtp_password")
    // Receiver email address.
    to := []string{
        "dummygandalf34@gmail.com",
        "tariq.riahi@gmail.com",
    }
    // smtp server configuration.
    smtpServer := smtpServer{host: os.Getenv("email_smtp_host"), port: os.Getenv("email_smtp_port")}
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


