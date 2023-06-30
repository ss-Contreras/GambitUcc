package bd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ss-Contreras/GambitUcc/models"
	secretmgo "github.com/ss-Contreras/GambitUcc/secretm.go"
)

var SecretModel models.SecretRDSJon
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretmgo.GetSecret(os.Getenv("SecreName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("Mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err := Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Conexion a la base de datos funcionando")
	return nil
}

func ConnStr(claves models.SecretRDSJon) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}
