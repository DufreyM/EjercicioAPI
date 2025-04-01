package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//estructura data
/*type incidente struct {
	ID string `json:"id"`
	Reportero string `json:"reportero"`
	Descripcion string `json:"descripcion"`
	Status string `json:"status"`
	Fecha time.Time `json:"fecha"`
}*/

var db *gorm.DB
var err error

// Modelo para GORM
type Incidente struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Reportero   string    `json:"reportero"`
	Descripcion string    `json:"descripcion"`
	Status      string    `json:"status"`
	Fecha       time.Time `json:"fecha"`
}

// Función para conectar a la base de datos
func initDB() {
	// Cambia estos valores según tu configuración
	dsn := "host=localhost user=postgres password=12345 dbname=incidentes_db port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	fmt.Println("Conexión a la base de datos exitosa")

	// Migración: Crear la tabla si no existe
	if err := db.AutoMigrate(&Incidente{}); err != nil {
		log.Fatal("Error al migrar la base de datos:", err)
	}
	fmt.Println("Tabla 'incidentes' creada o actualizada")
}

//Album data 
/*var incidentes = []incidente{
	{ID: "1", Reportero: "Majo", Descripcion: "Accidente automovilisto en Quiché", Status: "pendiente"},
	{ID: "2", Reportero: "Mejía", Descripcion: "Gran depresión de los estudiantes universitarios", Status: "completo"}, 
}*/

//Get para los incidentes
/*func getincidentes(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, incidentes)
}*/

// Obtener todos los incidentes
func getIncidentes(c *gin.Context) {
	var incidentes []Incidente
	db.Find(&incidentes)
	c.JSON(http.StatusOK, incidentes)
}

//Post para los incidentes
func postIncidentes(c *gin.Context) {
	var newIncidente Incidente

	if err := c.BindJSON(&newIncidente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if newIncidente.Reportero == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El reportero es obligatorio"})
		return 
	}

	if len(newIncidente.Descripcion) < 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La descripción debe tener al menos 10 caracteres"})
		return 
	}

	newIncidente.Fecha = time.Now()

	db.Create(&newIncidente)
	c.JSON(http.StatusCreated, newIncidente)
}

//Función para buscar algo según su id. 
/*func getIncidentesbyID(c *gin.Context){
	id := c.Param("id")

	for _, a:= range incidentes{
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Incidente no encontrado"})
}*/

// Obtener un incidente por ID
func getIncidenteByID(c *gin.Context) {
	id := c.Param("id")
	var incidente Incidente

	if err := db.First(&incidente, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Incidente no encontrado"})
		return
	}

	c.JSON(http.StatusOK, incidente)
}

//Funcion para borrar un incidente
/*func deleteIncidentebyID(c *gin.Context){
	id := c.Param("id")

	for i, a:= range incidentes{
		if a.ID == id {
			incidentes = append(incidentes[:1], incidentes[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Incidente eliminado"})
			return 
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Incidente no encontrado"})
}*/

// Eliminar un incidente por ID
func deleteIncidenteByID(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Incidente{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Incidente no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Incidente eliminado"})
}

/*func updateIncidentbyID(c *gin.Context){
	id := c.Param("id")
	var updateIncidente Incidente 

	if err := c.BindJSON(&updateIncidente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return 
	}

	for i, a := range incidentes {
		if a.ID == id {
			incidentes[i].Status = updateIncidente.Status
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Estado Actualizado"})
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Incidente no encontrado"})
}*/

// Actualizar el estado de un incidente por ID
func updateIncidenteByID(c *gin.Context) {
	id := c.Param("id")
	var incidente Incidente

	if err := db.First(&incidente, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Incidente no encontrado"})
		return
	}

	var updateData struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	incidente.Status = updateData.Status
	db.Save(&incidente)

	c.JSON(http.StatusOK, gin.H{"message": "Estado actualizado"})
}


func main(){
	initDB()
	router:=gin.Default()
	router.GET("/incidentes", getIncidentes)
	router.POST("/incidentes", postIncidentes)
	router.GET("/incidentes/:id", getIncidenteByID)
	router.DELETE("/incidentes/:id", deleteIncidenteByID)
	router.PUT("/incidentes/:id", updateIncidenteByID)
	router.Run("localhost:8083")
}