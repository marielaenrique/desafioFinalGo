package main

import (
	"database/sql"
	"desafioFinalGo/cmd/server/handler"
	"desafioFinalGo/internal/odontologo"
	"desafioFinalGo/internal/paciente"
	"desafioFinalGo/internal/turno"
	"desafioFinalGo/pkg/middleware"
	"desafioFinalGo/pkg/store"
	"log"
	"os"

	"desafioFinalGo/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Desafío Go Backend 3
// @version 1.0
// @description Esta API maneja Odontólogos, Pacientes y Turnos
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

	repositoryTurnos := turno.NewRepository(storage)
	serviceTurnos := turno.NewService(repositoryTurnos)
	turnoHandler := handler.NewTurnoHandler(serviceTurnos)

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))

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

	turnos := r.Group("/turnos")
	{

		turnos.GET(":id", turnoHandler.GetByID())
		turnos.POST("", turnoHandler.Post())
		turnos.PUT(":id", turnoHandler.Put())
		turnos.PATCH(":id", turnoHandler.Patch())
		turnos.DELETE(":id", turnoHandler.Delete())
	}

	r.Run(":8080")
}
