package main

import (
	"database/sql"
	"desafioFinalGo/cmd/server/handler"
	"desafioFinalGo/internal/odontologo"
	"desafioFinalGo/pkg/store"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	bd, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/turnos-odontologia")
	if err != nil {
		log.Fatal(err)
	}

	storage := store.NewSqlStore(bd)
	repository := odontologo.NewRepository(storage)
	service := odontologo.NewService(repository)
	odontologoHandler := handler.NewOdontologoHandler(service)

	r := gin.Default()

	odontologos := r.Group("/odontologos")
	{
		odontologos.GET(":id", odontologoHandler.GetByID())
		odontologos.POST("", odontologoHandler.Post())
		odontologos.PUT(":id", odontologoHandler.Put())
	}

	r.Run(":8080")
}
