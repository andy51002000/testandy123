
package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Member represents the customer model
type Member struct {
	PK         uint
	Username   string
	CreateTime time.Time
}

// BorrowFee represents the transaction model
type BorrowFee struct {
	MemberFK         uint
	PK               uint
	Type             int
	BorrowFee        float64
	CreateTime       time.Time
	TransactionCount int // Transaction count
}

func main() {
	// Connect to the MySQL database
	db, err := sql.Open("mysql", "root:root@tcp(mysql-container:3306)/database_name?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	// Generate fake data
	generateFakeData(db, 1000, 5000)

	fmt.Println("Fake data generated successfully.")
}

// generateFakeData generates fake customer and transaction data and inserts them into the database
func generateFakeData(db *sql.DB, numCustomers, numTransactions int) {
	// Generate fake customers
	customers := createCustomers(db, numCustomers)

	// Generate fake transactions
	createTransactions(db, customers, numTransactions)
}

// createCustomers generates fake customers and inserts them into the database
func createCustomers(db *sql.DB, numCustomers int) []Member {
	var customers []Member
	for i := 0; i < numCustomers; i++ {
		customer := Member{
            PK: uint(i),
			Username:   fmt.Sprintf("User%d", i+1),
			CreateTime: time.Now().Add(-time.Duration(rand.Intn(365*3)) * 24 * time.Hour),
		}
		_, err := db.Exec("INSERT INTO members (username, create_time) VALUES (?, ?)", customer.Username, customer.CreateTime)
		if err != nil {
			panic(err)
		}
		customers = append(customers, customer)
	}
	return customers
}

// createTransactions generates fake transactions and inserts them into the database
func createTransactions(db *sql.DB, customers []Member, numTransactions int) {
	memberTransHist := make(map[uint]int)
    
    
	i := 0
	for i < numTransactions {
		idxCustomer := rand.Intn(len(customers))
		customer := customers[idxCustomer]
		numCustomTransactions := rand.Intn(11)

		for j := 0; j < numCustomTransactions; j++ {
			newCount := memberTransHist[customer.PK] + 1
            currentTime := time.Now()
			transaction := BorrowFee{
				MemberFK:         customer.PK,
				Type:             1,
				BorrowFee:        rand.Float64() * 1000,
				CreateTime:       currentTime,
				TransactionCount: newCount, // Initial transaction count is 1
			}
			_, err := db.Exec("INSERT INTO borrow_fees (member_fk, type, borrow_fee, create_time, transaction_count) VALUES (?, ?, ?, ?, ?)", transaction.MemberFK, transaction.Type, transaction.BorrowFee, transaction.CreateTime, transaction.TransactionCount)
			if err != nil {
				panic(err)
			}

			memberTransHist[customer.PK] = newCount
            
		}

		i = i + 1 + numCustomTransactions
	}
}

