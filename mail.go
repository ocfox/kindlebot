package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

type EmailAccount struct {
	MailAddress string
	Password    string
	Server      string
}

func TestBuild() []byte {
	var headers = make(map[string]string)
	headers["Content-Transfer-Encoding"] = "7bit"
	headers["Content-Type"] = "text/plain; charset=US-ASCII; format=flowed"

	var data []byte
	for k, v := range headers {
		data = append(data, []byte(fmt.Sprintf("%s: %s\r\n", k, v))...)
	}
	data = append(data, []byte("Hello, this is a test email message.\r\n")...)

	return data
}

func BuildAttachment(attachment Attachment) []byte {
	var headers = make(map[string]string)
	headers["Content-Transfer-Encoding"] = "base64"
	headers["Content-Type"] = attachment.MIME + "; name=\"" + attachment.Filename + "\""
	headers["Content-Disposition"] = "attachment; filename=\"" + attachment.Filename + "\"; " + "size=" + fmt.Sprint(len(attachment.Data))

	var data []byte

	for k, v := range headers {
		data = append(data, []byte(fmt.Sprintf("%s: %s\r\n", k, v))...)
	}
	data = append(data, []byte("\r\n")...)

	data = append(data, attachment.Data...)
	return data
}

func BuildMessage(send EmailAccount, recipient string, attachment Attachment) io.Reader {
	var headers = make(map[string]string)
	headers["MIME-Version"] = "1.0"
	headers["Date"] = time.Now().Format(time.RFC1123Z)
	headers["From"] = send.MailAddress
	headers["To"] = recipient
	headers["Subject"] = "[kindle-bot] " + attachment.Filename
	headers["Content-Type"] = "multipart/mixed; boundary=\"=_kuroneko\""

	var data []byte
	for k, v := range headers {
		data = append(data, []byte(fmt.Sprintf("%s: %s\r\n", k, v))...)
	}
	data = append(data, []byte("--=_kuroneko\r\n")...)
	data = append(data, BuildAttachment(attachment)...)
	data = append(data, []byte("\r\n--=_kuroneko--\r\n")...)

	msg := bytes.NewReader(data)

	return msg
}

func SendMail(mail EmailAccount, recipient string, attachment Attachment) {

	auth := sasl.NewPlainClient("", mail.MailAddress, mail.Password)
	msg := BuildMessage(mail, recipient, attachment)
	to := []string{recipient}
	err := smtp.SendMailTLS(mail.Server, auth, mail.MailAddress, to, msg)
	println("sent mail")
	if err != nil {
		log.Fatal(err)
	}
}
