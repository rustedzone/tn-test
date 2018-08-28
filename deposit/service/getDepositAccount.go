package depositservice

import (
	"log"
	"time"
	"tn-test/conf"

	"github.com/akbarkn/aknstruct"
)

func GetDepositAccount(account string) ([]map[string]interface{}, error) {

	db := conf.Db()
	conf.Db().Close()
	defer db.Close()

	var result []map[string]interface{}

	// errResult := make([]map[string]interface{}, 0)

	rows, err := db.Query(`with sumcash as (
		select sum(cash) from tx_deposit where account_number=$1
	)
	select a.seqno_deposit,a.account_number,a.deposit_date,a.cash,b.* from tx_deposit a,sumcash b where account_number=$1`, account)
	if err != nil {
		log.Println("loc : deposit")
		log.Println("fn : GetDepositAccount")
		log.Println("error 1 : ", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var d Deposit
		rows.Scan(
			&d.DepositID,
			&d.AccountNumber,
			&d.Date,
			&d.Deposit,
			&d.Sum,
		)
		date, err := time.Parse("2006-01-02T15:04:05Z", d.Date)
		if err != nil {
			log.Println("loc : deposit")
			log.Println("fn : GetDepositAccount")
			log.Println("error 2 : ", err.Error())
			return nil, err
		}
		d.Date = date.Format("02-01-2006")
		maps := aknstruct.Map(d)
		result = append(result, maps)
	}

	log.Println("result :", result)

	return result, nil

}
