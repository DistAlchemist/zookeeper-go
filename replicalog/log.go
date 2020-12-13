package replicalog

import "fmt"

//Log structure of a piece of log
type Log struct {
	Action int
	Info   int
	Str    string
	next   *Log
	before *Log
}

//Lognum the number of logs record now
var Lognum int

//Lognow the number of logs of this station
var Lognow int

//Record pointer to all logs
var Record *[]Log

//Getlognum lognum records the number of action done
func Getlognum() int {
	return Lognum
}

//Initlog create now=nil
func Initlog() {
	Lognow = 0
	Lognum = 0
	var newlog Log
	newlog.before = nil
	Record = new([]Log)
	*Record = append(*Record, newlog) //add record number 0
}

//Recordlog set a new record in log
func Recordlog(action int, info int, str string) {
	fmt.Printf("log No.%d :%d %d %s recorded\n", Lognow, action, info, str)
	log1 := new(Log)
	log1.before = &(*Record)[Lognow]
	if (*Record) != nil {
		(*Record)[Lognow].next = log1
	}
	log1.Action = action
	log1.Info = info
	log1.Str = str
	*Record = append(*Record, *log1)
	Lognum++
	Lognow++
}

//Getlog by number
func Getlog(num int) Log {
	return (*Record)[num]
}
