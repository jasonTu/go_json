package main

import (
	"log"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


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
