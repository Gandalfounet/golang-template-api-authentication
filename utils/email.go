// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
    "os"
    "log"
    "github.com/joho/godotenv"
	"net/smtp"
    "html/template"
    "bytes"
    //"path/filepath"
    "strconv"
    "strings"
    "time"
    "github.com/vanng822/go-premailer/premailer"
    "github.com/jaytaylor/html2text"
)

var (
    debug     bool
    templates *template.Template
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
type ContentLoginToken struct {
    Email  string
    Message  string
    Name   string
    Date   string
    URL    string
    Token  string
    Expiry time.Time
}
func Send(msg ContentLoginToken, templateName string) {
    fmt.Println(msg)
    errTmp := parseTemplates()
    if errTmp != nil {
        log.Fatal("Error loading templates")
    }

    err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Sender data.
    from := os.Getenv("email_smtp_user")
    password := os.Getenv("email_smtp_password")
    // Receiver email address.
    to := []string{
        "tariq.riahi@gmail.com",
    }
    // smtp server configuration.
    smtpServer := smtpServer{host: os.Getenv("email_smtp_host"), port: os.Getenv("email_smtp_port")}
    // Message.

    buf := new(bytes.Buffer)
    
    contentMsg := msg

    if err := templates.ExecuteTemplate(buf, templateName, contentMsg); err != nil {
        fmt.Println(err)
        return
    }
    prem, _ := premailer.NewPremailerFromString(buf.String(), premailer.NewOptions())
    html, err := prem.Transform()
    if err != nil {
        fmt.Println(err)
        return
    }
    html2 := html

    text, err := html2text.FromString(html2, html2text.Options{PrettyTables: true})
    
    if err != nil {
        fmt.Println(err)
        return
    }

    messageHtml := []byte(text)
    
    // Authentication.
    auth := smtp.PlainAuth("", from, password, smtpServer.host)
    // Sending email.
    err = smtp.SendMail(smtpServer.Address(), auth, from, to, messageHtml)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Email Sent!")
}

func parse(msg string) (string, error) {
    buf := new(bytes.Buffer)
    if err := templates.ExecuteTemplate(buf, "message", msg); err != nil {
        return "", err
    }
    prem, _ := premailer.NewPremailerFromString(buf.String(), premailer.NewOptions())
    html, err := prem.Transform()
    if err != nil {
        return "", err
    }
    html2 := html

    text, err := html2text.FromString(html2, html2text.Options{PrettyTables: true})
    
    if err != nil {
        return "", err
    }
    
    return text, nil
}

func parseTemplates() error {
    templates = template.New("").Funcs(fMap)

    //return filepath.Walk("./modules/User/Shared/templates", walkPath)
    templates.ParseFiles("./modules/User/Shared/templates/email/footer.html")
    templates.ParseFiles("./modules/User/Shared/templates/email/header.html")
    templates.ParseFiles("./modules/User/Shared/templates/email/styles.html")
    templates.ParseFiles("./modules/User/Shared/templates/email/auth/loginToken.html")
    templates.ParseFiles("./modules/User/Shared/templates/email/auth/resetPassword.html")
    templates.ParseFiles("./modules/User/Shared/templates/email/auth/updatePassword.html")
    templates.ParseFiles("./modules/User/Shared/templates/email/auth/validatePassword.html")
    templates.ParseFiles("./modules/User/Shared/templates/email/auth/message.html")

    return nil
}

func walkPath(path string, info os.FileInfo, err error) error {
    if strings.Contains(path, ".html") {
        _, err = templates.ParseFiles(path)
        return err
    }
    return err
}

var fMap = template.FuncMap{
    "formatAsDate":     formatAsDate,
    "formatAsDuration": formatAsDuration,
}

func formatAsDate(t time.Time) string {
    year, month, day := t.Date()
    return fmt.Sprintf("%d.%d.%d", day, month, year)
}

func formatAsDuration(t time.Time) string {
    dur := t.Sub(time.Now())
    hours := int(dur.Hours())
    mins := int(dur.Minutes())

    v := ""
    if hours != 0 {
        v += strconv.Itoa(hours) + " hours and "
    }
    v += strconv.Itoa(mins) + " minutes"
    return v
}



