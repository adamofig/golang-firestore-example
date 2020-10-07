package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Menu struct {
	Nombre string `json:"nombre"`
	Tamano string `json:"tamano"`
}

func main() {

	var db *sql.DB
	var err error
	db, err = initTcpConnectionPool()
	if err != nil {
		fmt.Println(err)
	}

	//queryExample(db)
	str2 := "INSERT INTO ventas (cantidad, idmenu, idcliente, horaEntrega, fechaCompra, fechaEntrega) VALUES ('1', '3', '13', '12:18:48', '2018-12-03 12:18:48', '2018-12-03');INSERT INTO ventas (cantidad, idmenu, idcliente, horaEntrega, fechaCompra, fechaEntrega) VALUES ('1', '3', '13', '08:20:59', '2018-12-30 08:20:59', '2018-12-30');"

	db.Query(str2)
	// insertData(db)
}

func insertData(db *sql.DB) {

	var year int = 2021      //Quiero información para  2019 2020
	var month time.Month = 5 //
	var day int = 25         // aleatorio entre 1 y 30 excepto por febrero
	var hour int = 9         // aleatorio entre 9:00 y 21:00
	var min int = 30         // aleatorio entre 0 y 59

	datetime := time.Date(year, month, day, hour, min, 0, 0, time.UTC)

	fmt.Println(datetime.Format("2006-02-01T03:04:05"))
	//arr1 := []int{900, 1000, 1100, 1250, 1300, 1370, 1400, 1420, 1425, 1450, 1500, 1510, 1520, 1540, 1560, 1580, 1590, 1600, 1610, 1600, 1630, 1640, 1650, 1680, 1700, 1720, 1750, 1790, 1815, 1830}
	// arr2 := []int{0, 0, 0, 0, 0, 0, 1100, 1220, 1305, 1450, 1500, 1570, 1600, 1640, 1690, 1710, 1750, 1750, 1815, 1880, 1900, 1940, 1990, 2000, 2100, 2220, 2290, 2390, 2500, 2620}
	//arr3 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 700, 610, 650, 680, 715, 700, 780, 800, 840, 890, 910, 920, 950, 930, 990, 1100}
	arr4 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	sum1 := 0
	// sum2 := 0
	// sum3 := 0

	for i, _ := range arr4 {
		sum1 += arr4[i]
		// sum2 += arr2[i]
		// sum3 += arr3[i]

		// fmt.Println(idclient, idmenu)
		var multiInsert string = ""

		for j := 0; j < arr4[i]; j++ {

			numYear, numMonth := getMonthYear(i + 1)
			dateRandom := randate(numYear, numMonth)
			// log.Println(numYear, numMonth, dateRandom)

			rand.Seed(time.Now().UnixNano())
			idclient := rand.Intn(16) + 1
			idmenu := rand.Intn(30) + 1

			queryInsert := getInsertSqlString(idmenu, idclient, dateRandom, dateRandom, dateRandom)

			// log.Println(queryInsert)
			multiInsert = multiInsert + queryInsert
			// println(j)
		}
		println("::::", multiInsert)

		_, err := db.Query(multiInsert)

		if err != nil {
			log.Println("ocurrio un err", err)
		}
	}
}

//se le pasa un indice del 1 al 30 y obtiene un año y mes
func getMonthYear(index int) (int, int) {

	countYear := (index - 1) / 12
	if countYear >= 1 {
		monthCount := countYear * 12
		months := index - monthCount
		// println(" se tiene que aumentar", countYear, months)
		return 2018 + countYear, months
	}
	return 2018, index

}

func randate(year int, month int) time.Time {

	min := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	//println("aqui tengo una fecha", delta)

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

// Obtene una consulta de insersión de sql con los parametros que se le mandan para la tabla menu
func getInsertSqlString(idMenu int, idCliente int, horaEntrega time.Time, fechaCompra time.Time, fechaEntrega time.Time) string {
	// println("INSERT INTO `ventas` (`cantidad`, `idmenu`, `idcliente`, `horaEntrega`, `fechaCompra`, `fechaEntrega`) VALUES ('5', '1', '1', '16:24:15', '2019-12-22 00:00:00', '2019-12-22');")
	queryInsert := fmt.Sprintf("INSERT INTO ventas (cantidad, idmenu, idcliente, horaEntrega, fechaCompra, fechaEntrega) VALUES ('1', '%v', '%v', '%v', '%v', '%v');", idMenu, idCliente, horaEntrega.Format("15:04:05"), fechaEntrega.Format("2006-01-02 15:04:05"), fechaEntrega.Format("2006-01-02"))
	// println(queryInsert)
	return queryInsert
}

// Objetivo, insertar 10,000 Datos a la base de ventas.
// Cantidad no importa
// idMenu aleatorio entre 1 y 30 , idCliente aleatorio entre 1 y 500
// hora de entrega aleatorio entre las 9 am y 9pm con intervalos de 15 min.
// fecha de compra.
// compra desde telefono, sucursal, web_app, app_movil.
// fecha de entrega. regular ecepto por dias festivos.
// var db *sql.DB

// INSERT INTO `pasteleria_db`.`ventas` (`cantidad`, `idmenu`, `idcliente`, `horaEntrega`, `fechaCompra`, `fechaEntrega`) VALUES ('1', '1', '1', '16:24:15', '2019-12-22 00:00:00', '2019-12-22');

// db.Query("")

func queryExample(db *sql.DB) {

	rows, err := db.Query("SELECT * FROM menu limit 3;")
	if err != nil {
	}

	fmt.Println(rows)

	_ = db
	_ = err
	print("hola como estan")

	/* Bucle para recorrer todos los registros */
	// Parece que es necesario usar scan, y respetar el orden de los registros.
	// Muy tedioso entender como mapear. y que pasa si luego la base cambia, rompería el código.
	for rows.Next() {
		//var menu Menu
		var id string
		var nombre string
		var tamano string
		var sabor string
		var precio int

		/* Leemos el registro */
		rows.Scan(&id, &nombre, &tamano, &sabor, &precio)
		/* Escribimos el título por pantalla */
		fmt.Println(id, nombre, tamano, sabor, &precio)
	}

}

func initTcpConnectionPool() (*sql.DB, error) {
	// [START cloud_sql_mysql_databasesql_create_tcp]
	var (
		dbUser    = "root"
		dbPwd     = "root"
		dbTcpHost = "35.193.207.95"
		dbName    = "pasteleria_db"
	)
	var dbURI string
	dbURI = fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPwd, dbTcpHost, dbName)
	fmt.Println(dbURI)
	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}
	// [START_EXCLUDE]
	configureConnectionPool(dbPool)
	// [END_EXCLUDE]
	return dbPool, nil
	// [END cloud_sql_mysql_databasesql_create_tcp]
}

// configureConnectionPool sets database connection pool properties.
// For more information, see https://golang.org/pkg/database/sql
func configureConnectionPool(dbPool *sql.DB) {
	// [START cloud_sql_mysql_databasesql_limit]
	// Set maximum number of connections in idle connection pool.
	dbPool.SetMaxIdleConns(5)
	// Set maximum number of open connections to the database.
	dbPool.SetMaxOpenConns(7)
	// [END cloud_sql_mysql_databasesql_limit]
	// [START cloud_sql_mysql_databasesql_lifetime]
	// Set Maximum time (in seconds) that a connection can remain open.
	dbPool.SetConnMaxLifetime(1800)
	// [END cloud_sql_mysql_databasesql_lifetime]
}
