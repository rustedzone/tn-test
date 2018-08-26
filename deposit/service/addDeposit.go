package depositservice

import (
	"log"
	account "tn-test/account/service"
	"tn-test/conf"
)

func AddDeposit(req map[string]interface{}) error {

	db := conf.Db()
	conf.Db().Close()
	defer db.Close()

	//checkAccountNumber
	isExist, err := checkAccountNumber(req["account"].(string))
	if err != nil {
		log.Println("loc : deposit")
		log.Println("fn : addDeposit")
		log.Println("error 1 : ", err.Error())
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Println("loc : deposit")
		log.Println("fn : addDeposit")
		log.Println("error 2 : ", err.Error())
		tx.Rollback()
		return err
	}

	//insert if does not exist
	if !isExist {
		err = account.TxAddAccount(req, tx)
		if err != nil {
			log.Println("loc : deposit")
			log.Println("fn : addDeposit")
			log.Println("error 3 : ", err.Error())
			tx.Rollback()
			return err
		}
	}

	//insert to transaction
	var newID int64
	err = tx.QueryRow(`insert into tx_deposit(account_number,deposit_date,cash) values($1,$2,$3) returning seqno_deposit`, req["account"], req["date"], req["deposit"]).Scan(&newID)
	if err != nil {
		log.Println("loc : deposit")
		log.Println("fn : addDeposit")
		log.Println("error 4 : ", err.Error())
		tx.Rollback()
		return err
	}

	log.Println("new transaction ID :", newID)

	// everything is okay, commit
	err = tx.Commit()
	if err != nil {
		log.Println("loc : deposit")
		log.Println("fn : addDeposit")
		log.Println("error on commit : ", err.Error())
		tx.Rollback()
		return err
	}

	return nil

}
