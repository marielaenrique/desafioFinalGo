package main

import (
	"database/sql"
	"desafioFinalGo/cmd/server/handler"
	"desafioFinalGo/internal/odontologo"
	"desafioFinalGo/internal/paciente"
	"desafioFinalGo/pkg/middleware"
	"desafioFinalGo/pkg/store"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error when loading file .env")
	}

	bd, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/turnos-odontologia")
	if err != nil {
		log.Fatal(err)
	}

	storage := store.NewSqlStore(bd)

	repositoryOdontologos := odontologo.NewRepository(storage)
	serviceOdontologos := odontologo.NewService(repositoryOdontologos)
	odontologoHandler := handler.NewOdontologoHandler(serviceOdontologos)

	repositoryPacientes := paciente.NewRepository(storage)
	servicePacientes := paciente.NewService(repositoryPacientes)
	pacienteHandler := handler.NewPacienteHandler(servicePacientes)

	r := gin.Default()

	odontologos := r.Group("/odontologos")
	{
		odontologos.GET(":id", odontologoHandler.GetByID())
		odontologos.POST("", middleware.Authentication(), odontologoHandler.Post())
		odontologos.PUT(":id", middleware.Authentication(), odontologoHandler.Put())
		odontologos.PATCH(":id", middleware.Authentication(), odontologoHandler.Patch())
		odontologos.DELETE(":id", middleware.Authentication(), odontologoHandler.Delete())
	}

	pacientes := r.Group("/pacientes")
	{
		pacientes.GET(":id", pacienteHandler.GetByID())
		pacientes.POST("", middleware.Authentication(), pacienteHandler.Post())
		pacientes.PUT(":id", middleware.Authentication(), pacienteHandler.Put())
		pacientes.PATCH(":id", middleware.Authentication(), pacienteHandler.Patch())
		pacientes.DELETE(":id", middleware.Authentication(), pacienteHandler.Delete())

	}

	r.Run(":8080")

}
