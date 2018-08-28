package depositservice

import (
	"crypto/tls"
	"log"
	"strconv"
	"tn-test/conf"

	"github.com/akbarkn/aknenv"
	gomail "gopkg.in/gomail.v2"
)

func sendMail(request map[string]interface{}) error {

	db := conf.Db()
	conf.Db().Close()
	defer db.Close()

	//get target
	var target string
	err := db.QueryRow(`select email from account where account_number=$1`, request["account"]).Scan(&target)
	if err != nil {
		return err
	}

	log.Println("target :", target)

	m := gomail.NewMessage()
	m.SetHeader("From", aknenv.GetEnv("MAIL_USER"))
	m.SetHeader("To", target)
	m.SetHeader("Subject", "Transaction Detail")
	m.SetBody("text/html", "<p>HELLO TEST</p>")

	port, err := strconv.Atoi(aknenv.GetEnv("MAIL_SMTP_PORT"))
	if err != nil {
		log.Println(err.Error())
		return err
	}
	d := gomail.NewDialer(aknenv.GetEnv("MAIL_SMTP"), port, aknenv.GetEnv("MAIL_USER"), aknenv.GetEnv("MAIL_PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		// panic(err)
		log.Println(err.Error())
		return err
	}

	log.Println("d :", d)

	return nil

}
