package db_util

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetDbConnection(host string, port string, user string, passwd string, dbName string) *sql.DB {
	//  "cool:Cx111111@@tcp(10.30.1.89:3306)/coolcollege"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, passwd, host, port, dbName)
	fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func GetCoolPdbConnection(host string, port string) *sql.DB {
	dbUser := "cool"
	dbPassWd := "Cx111111@"
	dbName := "coolcollege"
	return GetDbConnection(host, port, dbUser, dbPassWd, dbName)
}

func TestPdb() {
	host := "10.30.1.89"
	port := "3306"
	db := GetCoolPdbConnection(host, port)
	querySql := "select count(*) from enterprise"
	rows, err := db.Query(querySql)
	if err != nil {
		panic(err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var count int
			err := rows.Scan(&count)
			if err != nil {
				panic(err)
			} else {
				fmt.Printf("count: %d", count)
			}
		}
	}
}

func TestDb() {
	host := "10.30.1.89"
	post := "3306"
	dbUser := "cool"
	dbPassWd := "Cx111111@"
	dbName := "coolcollege"
	db := GetDbConnection(host, post, dbUser, dbPassWd, dbName)
	querySql := "select id, original_name from enterprise limit 10"
	rows, err := db.Query(querySql)
	if err != nil {
		panic(err)
	} else {
		defer rows.Close()
		var enterpriseList [10]Enterprise
		i := 0
		for rows.Next() {
			// var id int
			// var name string
			var enterprise Enterprise
			err := rows.Scan(&enterprise.id, &enterprise.name)
			if err != nil {
				panic(err)
			} else {
				fmt.Printf("id: %d, name: %s\n", enterprise.id, enterprise.name)
				enterpriseList[i] = enterprise
				i++
			}
		}
		fmt.Println(len(enterpriseList))
		for i := 0; i < len(enterpriseList); i++ {
			fmt.Println(enterpriseList[i])
		}
	}
}

type Enterprise struct {
	id   int
	name string
}
