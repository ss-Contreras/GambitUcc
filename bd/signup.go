package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ss-Contreras/GambitUcc/models"
	"github.com/ss-Contreras/GambitUcc/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza registro")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateADD) VALUES ('" + sig.UserEmail + " ',' " + sig.UserUUID + "','" + tools.FechaMYSQL() + "')"
	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp > Ejecución exitosa ")
	return nil
}
