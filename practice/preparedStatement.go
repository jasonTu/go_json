package main

import (
	"log"
	"fmt"
	"strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var G_DB *sql.DB

func init() {
	var err error
	G_DB, err = sql.Open("mysql", "root:Trend#1..@/auditlog")
	if err != nil {
	    log.Fatal(err)
	}
}

func PrintLogTime(btime, etime, admin string) {
	var log_time string
	// Get data from database
	sql := fmt.Sprintf(
		"select `log_time` from `fakecompany` where `log_time` BETWEEN %s AND %s AND `admin` IN ('%s')",
		btime, etime, admin,
	)
	log.Println(sql)
	rows, err := G_DB.Query(sql)
	if err != nil {
	    log.Fatal(err)
	}
	// Parse the data
	for rows.Next() {
	    if err := rows.Scan(&log_time); err != nil {
	        log.Fatal(err)
	    }
	    fmt.Printf("log_time %s\n", log_time)
	}
	if err := rows.Err(); err != nil {
	    log.Fatal(err)
	}
}

func PrintLogTimeV2(btime, etime, admin string) {
	var log_time string
	// Get data from database
	rows, err := G_DB.Query(
		"select `log_time` from `fakecompany` where `log_time` BETWEEN ? AND ? AND `admin` IN (?)",
		btime, etime, admin,
	)
	if err != nil {
	    log.Fatal(err)
	}
	// Parse the data
	for rows.Next() {
	    if err := rows.Scan(&log_time); err != nil {
	        log.Fatal(err)
	    }
	    fmt.Printf("log_time %s\n", log_time)
	}
	if err := rows.Err(); err != nil {
	    log.Fatal(err)
	}
}

func TestOneCase(admin, admin2 string) {
	var log_time string
	var admins []interface{}
	// admins := make([]interface{}, 5)
	// admins[0] = admin
	// admins[1] = admin2
	// admins[2] = "log_time"
	// // admins[3] = "asc"
	// admins[3] = "1"
	// admins[4] = "3"
	// stmt := "select `log_time` from `fakecompany` where `admin` IN (?, ?) order by ? asc limit ?, ?"
	stmt := "select `log_time` from `fakecompany`limit 1"
	rows, err := G_DB.Query(stmt, admins...)
	if err != nil {
	    log.Fatal(err)
	}
	for rows.Next() {
	    if err := rows.Scan(&log_time); err != nil {
	        log.Fatal(err)
	    }
	    fmt.Printf("log_time %s\n", log_time)
	}
	if err := rows.Err(); err != nil {
	    log.Fatal(err)
	}
	admins = append(admins, "testa")
	log.Println(admins)
}

func main() {
	// btime, etime, admin := "1547715675", "1547725675", "SYS_ACCOUNT"
	// PrintLogTimeV2(btime, etime, admin)
	// btime, etime, admin = "1547715675", "1547725675", "SYS_ACCOUNT') limit 1  -- .com"
	// PrintLogTimeV2(btime, etime, admin)
	// btime, etime, admin = "1547715675", "1547725675", "SYS_ACCOUNT') or 1=1  -- .com"
	// PrintLogTimeV2(btime, etime, admin)
	admin := "jason"
	admin2 := "tony"
	TestOneCase(admin, admin2)
	accounts := "\"aaa\""
	accounts = strings.Replace(accounts, "\"", "", -1)
	fmt.Println(accounts)
	acts := strings.Split(accounts, ",")
	fmt.Println(acts)
}


/*
func main() {
	db, err := sql.Open("mysql", "root:Trend#1..@/auditlog")
	// rows, err := db.Query(
	// 	"select `log_time` from `company99` where ? BETWEEN ? AND ?",
	// 	"log_time", "0", "1948733687",
	// )
	// rows, err := db.Query(
	// 	"select `log_time` from `company99` where `id` BETWEEN ? AND ?",
	// 	"88", "93",
	// )
	// rows, err := db.Query(
	// 	"select `log_time` from `tmskynet.com` where `log_time` BETWEEN ? AND ? AND `admin` IN (?, ?)",
	// 	"1547732687", "1547733687", "support@tmskynet.com", "jak@hahq') limit 1  -- .com",
	// )
	// rows, err := db.Query(
	// 	"select `log_time` from `tmskynet.com` where `log_time` BETWEEN 1547732687 AND 1547733687 AND `admin` IN ('support@tmskynet.com', \"'jak@hahq') limit 1  -- .com\")",
	// )
	// rows, err := db.Query(
	// 	"select `log_time` from `tmskynet.com` where `log_time` BETWEEN ? AND ? AND (`admin`=? OR `admin`=?)",
	// 	"1547732687", "1547733687", "'support@tmskynet.com'", "'support2@tmskynet.com'",
	// )

	// Sql injection sample
	// rows, err := db.Query(
	// 	"select `log_time` from `company99` where `log_time` BETWEEN 1549296000 AND 1549348494 AND `admin` IN ('xyz') or 1=1 -- )",
	// )
	// Use prepared statement to prevent sql injection
	// rows, err := db.Query(
	// 	"select `log_time` from `company99` where `log_time` BETWEEN ? AND ? AND `admin` IN (?, ?, ?)",
	// 	"1549296000", "1549348494", "audituser", "admin@jason.com", "'xyz') or 1=1 -- ",
	// )
	// Sql injection with partial correct statement
	// rows, err := db.Query(
	// 	"select `log_time` from `company99` where `id` BETWEEN ? AND ? AND `admin` IN (?, ?, ?)",
	// 	"88", "89", "audituser", "admin@jason.com", "'xyz') or 1=1 -- ",
	// )
	// Just use one placeholder in IN lookup statement will no take effect
	rows, err := db.Query(
		"select `log_time` from `company99` where `id` BETWEEN ? AND ? AND `admin` IN (?)",
		"88", "89", "audituser, admin@jason.com, 'xyz') or 1=1 -- ",
	)
	if err != nil {
	    log.Fatal(err)
	}
	for rows.Next() {
	    var log_time string
	    if err := rows.Scan(&log_time); err != nil {
	        log.Fatal(err)
	    }
	    // fmt.Printf("log_time %s, account %s\n", log_time, account)
	    fmt.Printf("log_time %s\n", log_time)
	}
	if err := rows.Err(); err != nil {
	    log.Fatal(err)
	}
}
*/
