package database

var connetction string

/*
init function akan lansung di jalankan ketika di package database
di gunakan di direktori lain
*/
func init() {
	connetction = "MySql"
}

func Getdatabase() string {
	return connetction
}
