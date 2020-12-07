package replicalog

type log struct {
	action int
	str    string
	next   *log
	before *log
}

var lognum int
var now *log

//Getlognum lognum records the number of action done
func Getlognum() int {
	return lognum
}

//Initlog create now=nil
func Initlog() {
	now = nil
	lognum = 0
}

//Recordlog set a new record in log
func Recordlog(action int, str string) {
	log1 := new(log)
	log1.before = now
	if now != nil {
		now.next = log1
	}
	log1.action = action
	log1.str = str
	lognum++
}
