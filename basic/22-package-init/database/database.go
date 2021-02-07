package database

var connection string

func init() {
	connection = "MySQL Connected"
}

func GetDatabase() string {
	return connection
}
