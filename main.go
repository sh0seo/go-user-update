package main

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("DCCafe User RFID Update")

	// input
	email := "bhlee@digicaps.com"
	// input
	cardNumber := "0934853485092384509324850928340598342095830458340958"

	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PW")

	conn := "user=dcadmin password=%s host=%s dbname=dcadmin sslmode=disable"

	if len(email) == 0 {
		fmt.Printf("[Error] user email\n")
		return
	}

	if len(cardNumber) == 0 {
		fmt.Printf("[Error] Card Number empty\n")
		return
	}

	if len(dbHost) == 0 {
		fmt.Printf("[Error] Not find DB HOST\n")
		return
	}

	if len(dbPassword) == 0 {
		fmt.Printf("[Error] Not find DB Password\n")
		return
	}

	db, err := sql.Open("postgres", fmt.Sprintf(conn, dbPassword, dbHost))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var status int
	err = db.QueryRow("SELECT 1").Scan(&status)
	if err != nil {
		panic(err)
	}

	if status != 1 {
		panic(fmt.Sprintf("SELECT 1 is not %d\n", status))
	}

	// get User
	var user Users
	err = db.QueryRow(`SELECT email, rfid, company, index, name, leave, regdate, updatedate FROM users WHERE email = $1`, email).Scan(&user.email, &user.rfid, &user.company, &user.index, &user.name, &user.leave, &user.regdate, &user.updatedate)
	if err != nil {
		panic(err)
	}

	fmt.Printf("find user %s\n", user.name)

	rfid := hash(cardNumber)

	status = 0
	err = db.QueryRow(`SELECT 1 FROM users WHERE rfid = $1`, rfid).Scan(&status)
	if err == nil {
		fmt.Printf("[Error] Exist Card Number:%s\n", cardNumber)
		return
	}

	fmt.Printf("Get new RFID %s\n", rfid)

	_, err = db.Exec(`UPDATE users SET rfid = $1, updatedate = now() WHERE index = $2`, rfid, user.index)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Update User: %s, CardNumber: %s\n", user.name, cardNumber)
}

func hash(cardNumber string) string {
	sum := sha256.Sum256([]byte(cardNumber))
	return fmt.Sprintf("%x", sum)
}

// Users DB
type Users struct {
	email      string // PK
	rfid       string // PK
	company    string
	index      int
	name       string
	leave      bool
	regdate    time.Time
	updatedate time.Time
}
