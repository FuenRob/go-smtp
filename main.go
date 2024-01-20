package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func main() {
	host := "smtp.gmail.com"
	port := "587"
	from := "usuario@gmail.com"
	password := "CONTRASEÑA_EMAIL"

	toList := []string{"receptor@gmail.com"}
	msg := "From: " + from + "\r\n" +
		"To: " + toList[0] + "\r\n" +
		"Subject: Correo de prueba\r\n" +
		"\r\n" +
		"¡¡¡Hola Gophers!!!"
	body := []byte(msg)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Correo enviado exitosamente a todos los destinatarios")
}
