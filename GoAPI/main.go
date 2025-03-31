package main 

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

//estructura data 
type incidente struct {
	ID string `json:"id"`
	Reportero string `json:"reportero"`
	Descripcion string `json:"descripcion"`
	Status string `json:"status"`
	Fecha time.Time `json:"fecha"`
}

//Album data 
var incidentes = []incidente{
	{ID: "1", Reportero: "Majo", Descripcion: "Accidente automovilisto en Quiché", Status: "pendiente"}, 
}

//Get para los incidentes
func getincidentes(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, incidentes)
}

//Post para los incidentes
func postIncidentes(c *gin.Context) {
	var newIncidente incidente

	if err := c.BindJSON(&newIncidente); err != nil {
		return
	}

	newIncidente.Fecha = time.Now()

	incidentes = append(incidentes, newIncidente)
	c.IndentedJSON(http.StatusCreated, newIncidente)
}

//Función para buscar algo según su id. 
func getIncidentesbyID(c *gin.Context){
	id := c.Param("id")

	for _, a:= range incidentes{
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Incidente no encontrado"})
}

//Funcion para borrar un incidente
func deleteIncidentebyID(c *gin.Context){
	id := c.Param("id")

	for i, a:= range incidentes{
		if a.ID == id {
			incidentes = append(incidentes[:1], incidentes[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Incidente eliminado"})
			return 
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Incidente no encontrado"})
}

func updateIncidentbyID(c *gin.Context){
	id := c.Param("id")
	var updateIncidente incidente 

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
}


func main(){
	router:=gin.Default()
	router.GET("/incidentes", getincidentes)
	router.POST("/incidentes", postIncidentes)
	router.GET("/incidentes/:id", getIncidentesbyID)
	router.DELETE("/incidentes/:id", deleteIncidentebyID)
	router.PUT("/incidentes/:id", updateIncidentbyID)
	router.Run("localhost:8083")
}