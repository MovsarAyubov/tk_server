package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Connection = DatabaseConnection{}

type DatabaseConnection struct {
	Db *sqlx.DB
}

type DatabaseLogin struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func Connect(login DatabaseLogin) {
	// Database connect
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		login.Host, login.Port, login.Username, login.Password, login.Name)
	fmt.Println(psqlInfo)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)

	}
	Connection.Db = db
	fmt.Println("Successfully connected to database!")
}

func CreateTables(connection DatabaseConnection) {
	createWorkerTable(connection)
	createTypeOfWorkTable(connection)
	createCellTable(connection)
	createRowTable(connection)
	createDoneWorkTable(connection)
}

func createWorkerTable(connection DatabaseConnection) {
	_, err := connection.Db.Exec(`
			CREATE TABLE IF NOT EXISTS workers (
				id SERIAL PRIMARY KEY,
				name VARCHAR(20) NOT NULL,
				firstname VARCHAR(20) NOT NULL,
				patronomic VARCHAR(20) NOT NULL
			)
	`)
	if err != nil {
		panic(err)
	}
}

func createTypeOfWorkTable(connection DatabaseConnection) {
	_, err := connection.Db.Exec(`
		CREATE TABLE IF NOT EXISTS type_of_work (
			id SERIAL PRIMARY KEY,
			name VARCHAR(80) NOT NULL,
			uom VARCHAR(100) NOT NULL,
			price FLOAT NOT NULL,
			period VARCHAR(30) NOT NULL
		)
	`)
	if err != nil {
		panic(err)
	}
}

func createCellTable(connection DatabaseConnection) {
	_, err := connection.Db.Exec(`
		CREATE TABLE IF NOT EXISTS cell (
			id SERIAL PRIMARY KEY
		)
	`)
	if err != nil {
		panic(err)
	}
}

func createRowTable(connection DatabaseConnection) {
	_, err := connection.Db.Exec(`
		CREATE TABLE IF NOT EXISTS row (
			id SERIAL PRIMARY KEY,
			cell_id INT  NOT NULL,
			FOREIGN KEY (cell_id) REFERENCES cell(id)
		)
	`)
	if err != nil {
		panic(err)
	}
}

func createDoneWorkTable(connection DatabaseConnection) {
	_, err := connection.Db.Exec(`
		CREATE TABLE IF NOT EXISTS done_work (
			id SERIAL PRIMARY KEY,
			date VARCHAR(40) NOT NULL,
			worker_id INT  NOT NULL,
			type_of_work_id INT  NOT NULL,
			cell_id INT ,
			row_id INT ,
			count NUMERIC NOT NULL,
			income NUMERIC NOT NULL,
			FOREIGN KEY (worker_id) REFERENCES workers(id),
			FOREIGN KEY (type_of_work_id) REFERENCES type_of_work(id),
			FOREIGN KEY (cell_id) REFERENCES cell(id)
		)
	`)
	if err != nil {
		panic(err)
	}
}
