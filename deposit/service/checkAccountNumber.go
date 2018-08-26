package depositservice

import (
	"log"
	"tn-test/conf"
)

func checkAccountNumber(account string) (bool, error) {

	db := conf.Db()
	conf.Db().Close()
	defer db.Close()

	var count int
	err := db.QueryRow(`select count(0) from account where account_number=$1`, account).Scan(&count)
	if err != nil {
		log.Println("loc : deposit")
		log.Println("fn : checkAccountNumber")
		log.Println("error 1 : ", err.Error())
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil

}
