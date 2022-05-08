package MySql

import (
	"database/sql"
)

type MySqlProvider struct {
	host       string
	port       string
	user       string
	password   string
	datasource string
	connection *sql.DB
	Result     *sql.Rows
}

//TODO: find a way to implement all of Provider Interface methods

func (this MySqlProvider) Query(params map[string]string) {
	var query string = params["query"]
	if this.connection.Stats().InUse > 0 {
		rows, rr := this.connection.Query(query)
		if rr == nil {
			this.Result = rows
		}
	}
}

func (this MySqlProvider) GetResult() *sql.Rows {
	return this.Result
}

func (this *MySqlProvider) createConnectionString() string {
	return this.GetUser() + ":" + this.password + "@tcp(" +
		this.host + ":" + this.GetPort() + ")/" + this.GetDatasource()
}
func (this MySqlProvider) Connect() bool {
	db, err := sql.Open("mysql", this.createConnectionString())
	//db, err := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/test")
	this.connection = db
	if err != nil {
		return false
	} else {
		return true
	}
}

//Setters
func (this *MySqlProvider) SetUser(value string) {
	this.user = value
}
func (this *MySqlProvider) SetPassword(value string) {
	this.password = value
}

func (this *MySqlProvider) SetHost(value string) {
	this.host = value
}
func (this *MySqlProvider) SetPort(value string) {
	this.port = value
}
func (this *MySqlProvider) SetDatasource(value string) {
	this.datasource = value
}

//Getters
func (this *MySqlProvider) GetUser() string {
	return this.user
}
func (this *MySqlProvider) GetPassword() string {
	return this.password
}

func (this *MySqlProvider) GetHost() string {
	return this.host
}

func (this *MySqlProvider) GetPort() string {
	return this.port
}

func (this *MySqlProvider) GetDatasource() string {
	return this.datasource
}
