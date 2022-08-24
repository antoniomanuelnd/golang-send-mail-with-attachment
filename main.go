package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"strings"
)

type Mail struct {
	Sender  string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
}

func main() {

	addr := ""
	from := ""
	to := []string{
		"",
	}
	cc := []string{
		"",
	}
	bcc := []string{
		"",
	}

	subject := "Test mail"
	body := "This is a test mail"

	request := Mail{
		Sender:  from,
		To:      to,
		Cc:      cc,
		Bcc:     bcc,
		Subject: subject,
		Body:    body,
	}

	data := BuildMail(request, "test.txt")

	//If auth is required
	//auth := smtp.PlainAuth("", user, password, host)
	//err := smtp.SendMail(addr, auth, from, to, data)

	//If auth is not required
	err := smtp.SendMail(addr, nil, from, to, data)
	if err != nil {
		panic(err)
	}
	fmt.Println("Mail sent")

}

func BuildMail(mail Mail, filename string) []byte {
	var buf bytes.Buffer

	//Senders, receivers and subject
	buf.WriteString(fmt.Sprintf("From: %s\r\n", mail.Sender))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";")))
	buf.WriteString(fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";")))
	buf.WriteString(fmt.Sprintf("Bcc: %s\r\n", strings.Join(mail.Bcc, ";")))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", mail.Subject))

	boundary := "my-boundary"
	buf.WriteString("Mime-Version: 1.0;\r\n")
	buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))

	buf.WriteString(fmt.Sprintf("\n--%s\r\n", boundary))
	buf.WriteString("Content-Type: text/plain; charset=\"UTF-8\";\r\n")
	buf.WriteString(fmt.Sprintf("\r\n%s", mail.Body))

	buf.WriteString(fmt.Sprintf("\n--%s\r\n", boundary))
	buf.WriteString("Content-Type: text/plain; charset=\"UTF-8\";\r\n")
	buf.WriteString("Content-Transfer-Encoding: base64;\r\n")
	buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\r\n", filename))
	buf.WriteString(fmt.Sprintf("Content-ID: <%s>\r\n\r\n", filename))

	data := readFile(filename)

	b := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(b, data)
	buf.Write(b)
	buf.WriteString(fmt.Sprintf("\n--%s--\r\n", boundary))
	buf.WriteString("--")

	return buf.Bytes()
}

func readFile(fileName string) []byte {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
