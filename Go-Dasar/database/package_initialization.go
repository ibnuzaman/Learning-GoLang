package database

var connection string

func init() {
	connection = "Terkoneksi"
}

func GetDatabase() string {
	return connection
}
