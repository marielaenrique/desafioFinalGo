package main

import (
	"database/sql"
	"desafioFinalGo/cmd/server/handler"
	"desafioFinalGo/internal/odontologo"
	"desafioFinalGo/internal/paciente"
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
		odontologos.POST("", odontologoHandler.Post())
		odontologos.PUT(":id", odontologoHandler.Put())
		odontologos.PATCH(":id", odontologoHandler.Patch())
		odontologos.DELETE(":id", odontologoHandler.Delete())
	}

	pacientes := r.Group("/pacientes")
	{
		pacientes.GET(":id", pacienteHandler.GetByID())
		pacientes.POST("", pacienteHandler.Post())
		pacientes.PUT(":id", pacienteHandler.Put())
		pacientes.PATCH(":id", pacienteHandler.Patch())
		pacientes.DELETE(":id", pacienteHandler.Delete())

	}

	r.Run(":8080")

}
