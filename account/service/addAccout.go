package accountservice

import (
	"database/sql"
	"errors"
	"log"
)

func TxAddAccount(req map[string]interface{}, tx *sql.Tx) error {

	//check e-mail
	if req["email"] == nil || len(req["email"].(string)) == 0 {
		return errors.New("for new account, e-mail field is required")
	}

	stmt, err := tx.Prepare(`insert into account(account_number,email) values($1,$2)`)
	if err != nil {
		log.Println("loc : deposit")
		log.Println("fn : TxAddAccount")
		log.Println("error 1 : ", err.Error())
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req["account"], req["email"])
	if err != nil {
		log.Println("loc : deposit")
		log.Println("fn : TxAddAccount")
		log.Println("error 2 : ", err.Error())
		tx.Rollback()
		return err
	}

	return nil
}
