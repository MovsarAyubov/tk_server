package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"tk_app_v1/src/api"
	"tk_app_v1/src/database"

	// "tk_app_v1/src/models"

	"github.com/joho/godotenv"
)

func main() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	envFilePath := basepath + "/../.env"
	envFile, err := godotenv.Read(envFilePath)
	if err != nil {
		panic(err)
	}

	database.Connect(database.DatabaseLogin{
		Host:     envFile["PG_HOST"],
		Port:     envFile["PG_PORT"],
		Username: envFile["PG_USER"],
		Password: envFile["PG_PASSWORD"],
		Name:     envFile["PG_DATABASE_NAME"],
	})
	defer database.Connection.Db.Close()

	fmt.Println("success")

	database.CreateTables(database.Connection)
	// database.Connection.Db.Query("DROP TABLE cell")
	// database.Connection.Db.Query("DROP TABLE row")
	// database.Connection.Db.Query("DROP TABLE workers")

	// var cellId int = 1
	// for i := 1; i <= 330; i++ {

	// 	rows = append(rows, models.Row{Id: i, CellId: cellId})
	// 	if i%5 == 0 {
	// 		cellId++
	// 	}

	// }

	// for i := 0; i <= 66; i++ {
	// 	cells = append(cells, i)
	// }

	// for i := 0; i < len(typesOfWork); i++ {
	// 	query := `INSERT INTO type_of_work (name, uom, price, period) VALUES ($1, $2, $3, $4)`
	// 	err = database.Connection.Db.QueryRow(query, typesOfWork[i].Name, typesOfWork[i].Uom, typesOfWork[i].Price, typesOfWork[i].Period).Err()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// for i := 0; i < len(cells); i++ {
	// 	query := `INSERT INTO cell(id) VALUES($1)`
	// 	err = database.Connection.Db.QueryRow(query, cells[i]).Err()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// 	for i := 1; i < len(rows); i++ {
	// 		query := `INSERT INTO row (id, cell_id) VALUES ($1, $2)`
	// 		database.Connection.Db.Exec(query, rows[i].Id, rows[i].CellId)
	// 	}

	// Router
	api.StartRouter(envFile["SERVER_PORT"])

}

// var MyTypeOfWork struct {
// 	id     int
// 	period string
// 	name   string
// 	uom    string
// 	price  int
// }

// var typesOfWork = []models.TypeOfWorkModel{
// 	{
// 		Id:     1,
// 		Name:   "Вынос матов",
// 		Uom:    "Ряд",
// 		Period: "Выброска",
// 		Price:  200},
// 	{
// 		Id:     2,
// 		Name:   "Вынос отдельно кубиков",
// 		Uom:    "ряд",
// 		Period: "Выброска",
// 		Price:  100,
// 	},
// 	{
// 		Id:     3,
// 		Name:   "Вынос растений без отключения поливов",
// 		Uom:    "Ряд",
// 		Period: "Выброска",
// 		Price:  340,
// 	},
// 	{
// 		Id:     4,
// 		Name:   "Вынос растений при отключении поливов",
// 		Uom:    "Ряд",
// 		Period: "Выброска",
// 		Price:  290,
// 	},
// 	{
// 		Id:     5,
// 		Name:   "Мойка лотков без керхера",
// 		Uom:    "Ряд",
// 		Period: "Выброска",
// 		Price:  250,
// 	},
// 	{
// 		Id:     6,
// 		Name:   "Мойка лотков после керхера",
// 		Uom:    "Ряд",
// 		Period: "Выброска",
// 		Price:  110,
// 	},
// 	{
// 		Id:     7,
// 		Name:   "Мойка форточек в производственном отделении",
// 		Uom:    "Ряд",
// 		Period: "Выброска",
// 		Price:  315,
// 	},
// 	{
// 		Id:     8,
// 		Name:   "Мытье полов",
// 		Uom:    "Ряд",
// 		Period: "Выброска",
// 		Price:  110,
// 	},
// 	{
// 		Id:     9,
// 		Name:   "Подметание и уборка с поднятием напольного покрытия",
// 		Uom:    "Ряд",
// 		Period: "Выброска",
// 		Price:  160,
// 	},
// 	{
// 		Id:     10,
// 		Name:   "Подметание при выброске",
// 		Uom:    "Ряд",
// 		Period: "Выброска",
// 		Price:  150,
// 	},
// 	{
// 		Id:     11,
// 		Name:   "Подметание с подъемом труборельс",
// 		Uom:    "Ряд",
// 		Period: "Выброска",
// 		Price:  150,
// 	},
// 	{
// 		Id:     12,
// 		Name:   "Подъем труборельс с обработкой ножек",
// 		Uom:    "Ряд",
// 		Period: "Выброска",
// 		Price:  200,
// 	},
// 	{
// 		Id:     13,
// 		Name:   "Разрез растений",
// 		Uom:    "Ряд",
// 		Period: "Выброска",
// 		Price:  110,
// 	},
// 	{
// 		Id:     14,
// 		Name:   "Сбор стебледержателей",
// 		Uom:    "Ряд",
// 		Period: "Выброска",
// 		Price:  40,
// 	},
// 	{
// 		Id:     15,
// 		Name:   "Сбор шпагата",
// 		Uom:    "Ряд",
// 		Period: "Выброска",
// 		Price:  70,
// 	},
// 	{
// 		Id:     16,
// 		Name:   "Снятие акваштекира с капельниц",
// 		Uom:    "Ряд",
// 		Period: "Выброска",
// 		Price:  70,
// 	},
// 	{
// 		Id:     17,
// 		Name:   "Снятие капельниц",
// 		Uom:    "Ряд",
// 		Period: "Выброска",
// 		Price:  30,
// 	},
// 	{
// 		Id:     18,
// 		Name:   "Срезка растений",
// 		Uom:    "Ряд",
// 		Period: "Выброска",
// 		Price:  60,
// 	},
// 	{
// 		Id:     19,
// 		Name:   "Намотка шпагата",
// 		Uom:    "Шт",
// 		Period: "Общее производство",
// 		Price:  2,
// 	},
// 	{
// 		Id:     20,
// 		Name:   "Подготовка крючков",
// 		Uom:    "Шт",
// 		Period: "Общее производство",
// 		Price:  1.6,
// 	},
// 	{
// 		Id:     21,
// 		Name:   "Подметание",
// 		Uom:    "Ряд",
// 		Period: "Общее производство",
// 		Price:  110,
// 	},
// 	{
// 		Id:     22,
// 		Name:   "Почасовые работы",
// 		Uom:    "Час",
// 		Period: "Общее производство",
// 		Price:  125,
// 	},
// 	{
// 		Id:     23,
// 		Name:   "Почасовые работы(грузчики)",
// 		Uom:    "Час",
// 		Period: "Общее производство",
// 		Price:  170,
// 	},
// 	{
// 		Id:     24,
// 		Name:   "Почасовые работы(звеньевые)",
// 		Uom:    "Час",
// 		Period: "Общее производство",
// 		Price:  170,
// 	},
// 	{
// 		Id:     25,
// 		Name:   "Сбор коробок",
// 		Uom:    "Шт",
// 		Period: "Общее производство",
// 		Price:  1,
// 	},
// 	{
// 		Id:     26,
// 		Name:   "Установка капельниц",
// 		Uom:    "Ряд",
// 		Period: "Общее производство",
// 		Price:  50,
// 	},
// 	{
// 		Id:     27,
// 		Name:   "Подрутка",
// 		Uom:    "Ряд",
// 		Period: "Плодоношение",
// 		Price:  180,
// 	},
// 	{
// 		Id:     28,
// 		Name:   "Припускание",
// 		Uom:    "Ряд",
// 		Period: "Плодоношение",
// 		Price:  60,
// 	},
// 	{
// 		Id:     29,
// 		Name:   "Прищипка",
// 		Uom:    "Ряд",
// 		Period: "Плодоношение",
// 		Price:  60,
// 	},
// 	{
// 		Id:     30,
// 		Name:   "Сбор урожая Бьёрн F1(второй сорт)",
// 		Uom:    "Кг",
// 		Period: "Плодоношение",
// 		Price:  2,
// 	},
// 	{
// 		Id:     31,
// 		Name:   "Сбор урожая Бьёрн F1(нестандарт)",
// 		Uom:    "Кг",
// 		Period: "Плодоношение",
// 		Price:  1.5,
// 	},
// 	{
// 		Id:     32,
// 		Name:   "Сбор урожая Бьёрн F1(первый сорт)",
// 		Uom:    "Кг",
// 		Period: "Плодоношение",
// 		Price:  6,
// 	},
// 	{
// 		Id:     33,
// 		Name:   "Удаление листа",
// 		Uom:    "Ряд",
// 		Period: "Плодоношение",
// 		Price:  180,
// 	},
// 	{
// 		Id:     34,
// 		Name:   "Дренаж",
// 		Uom:    "Ряд",
// 		Period: "Посадка",
// 		Price:  50,
// 	},
// 	{
// 		Id:     35,
// 		Name:   "Завоз матов",
// 		Uom:    "Ряд",
// 		Period: "Посадка",
// 		Price:  120,
// 	},
// 	{
// 		Id:     36,
// 		Name:   "Обработка и установка капельниц",
// 		Uom:    "Ряд",
// 		Period: "Посадка",
// 		Price:  140,
// 	},
// 	{
// 		Id:     37,
// 		Name:   "Перенос рассады",
// 		Uom:    "Час",
// 		Period: "Посадка",
// 		Price:  125,
// 	},
// 	{
// 		Id:     38,
// 		Name:   "Перенос рассады(187.5)",
// 		Uom:    "Час",
// 		Period: "Посадка",
// 		Price:  187.5,
// 	},
// 	{
// 		Id:     39,
// 		Name:   "Подвязка растений и подкрутка",
// 		Uom:    "Ряд",
// 		Period: "Посадка",
// 		Price:  300,
// 	},
// 	{
// 		Id:     40,
// 		Name:   "Подкрутка в период посадки",
// 		Uom:    "Ряд",
// 		Period: "Посадка",
// 		Price:  200,
// 	},
// 	{
// 		Id:     41,
// 		Name:   "Развеска шпагата",
// 		Uom:    "Ряд",
// 		Period: "Посадка",
// 		Price:  380,
// 	},
// 	{
// 		Id:     42,
// 		Name:   "Расстановка растений",
// 		Uom:    "Ряд",
// 		Period: "Посадка",
// 		Price:  240,
// 	},
// 	{
// 		Id:     43,
// 		Name:   "Снятие пленки с матов, установка стебледержателей и капельниц",
// 		Uom:    "Ряд",
// 		Period: "Посадка",
// 		Price:  120,
// 	},
// 	{
// 		Id:     44,
// 		Name:   "Удаление листа",
// 		Uom:    "Ряд",
// 		Period: "Посадка",
// 		Price:  150,
// 	},
// 	{
// 		Id:     45,
// 		Name:   "Удаление пасынков",
// 		Uom:    "Ряд",
// 		Period: "Посадка",
// 		Price:  150,
// 	},
// 	{
// 		Id:     16,
// 		Name:   "Установка стебледержателей",
// 		Uom:    "Ряд",
// 		Period: "Посадка",
// 		Price:  70,
// 	},
// }

// var cells = []int{}

// var rows = []models.Row{}
