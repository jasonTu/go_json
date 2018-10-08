package main

import (
	"fmt"
)

type AuditLogBase interface {
	virtualfunc() string
}

type AuditLog struct {
	Id        int
	LogTime   int
}

type AuditLogExt struct {
	Id        int
	LogTime   int
	Ip        string
}

func (log AuditLog) virtualfunc() string {
	return "AuditLog"
}

func (log AuditLogExt) virtualfunc() string {
	return "AuditLogExt"
}

func getAuditLogs(logType int) (logs []AuditLogBase) {
	var alogExt AuditLogExt
	var alog AuditLog

	for i := 0; i < 5; i++ {
		if logType == 1 {
			alogExt.Id = i
			alogExt.LogTime = 1500000000 + i
			alogExt.Ip = "192.168.1.1"
			logs = append(logs, alogExt)
		} else {
			alog.Id = i
			alog.LogTime = 1500000000 + i
			logs = append(logs, alog)
		}
	}
	return
}

func main() {
	var logs []AuditLogBase
	logs = getAuditLogs(1)
	fmt.Println("logs are: ", logs)
}

/*
Result；
logs are:  [
{0 1500000000 192.168.1.1}
{1 1500000001 192.168.1.1}
{2 1500000002 192.168.1.1}
{3 1500000003 192.168.1.1}
{4 1500000004 192.168.1.1}
]
 */
