# API de Incidentes en Go con PostgreSQL

Este proyecto es una API REST para la gestión de incidentes, desarrollada en **Go** utilizando el framework **Gin** y la librería **GORM** para la conexión con **PostgreSQL**.

> **Nota:** La carpeta de **JavaAPI** era para la primera versión de la API, pero quería aprender a utilizar **Go** y me pasé a este, por lo que la conexión a la base de datos se encuentra en la carpeta llamada **GoAPI**.  
> Al final terminé utilizando **Go** y **Postgres**.

## Requisitos

Antes de ejecutar este proyecto, asegúrate de tener instalados:

- [Go](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Gin](https://github.com/gin-gonic/gin) (`go get github.com/gin-gonic/gin`)
- [GORM](https://gorm.io/) y el driver de PostgreSQL (`go get gorm.io/driver/postgres`)

## Instalación y Configuración

1. **Clonar el repositorio**  
   git clone https://github.com/DufreyM/EjercicioAPI.git
   cd EjercicioAPI/GoAPI

   📂 EjercicioAPI/  
 ├── 📂 GoAPI/  
 │   ├── main.go  
 │   ├── go.mod  
 │   ├── go.sum  
 ├── 📂 JavaAPI/  (Versión en Java incompleta, solo fue para empezar pero quería conocer un nuevo lenguaje.)  


3. **Base de datos**
   Debes cambiar la variable dsn con la información de tu base de datos en postgres.
   Postgres por defecto debe estar activo y configurada en el puerto 5432

4. Editar la variable dsn en el archivo principal main.go. 
   dsn := "host=localhost user=postgres password=TU_PASSWORD dbname=incidentes_db port=5432 sslmode=disable"

5. go run main.go, ejecutar.
**La API se ejecutará en http://localhost:8083/.

## EJEMPLOS DE USO 
Recomendable usar Postman. 

1. POST 
http://localhost:8083/incidentes 
Cuerpo: 
{
  "reportero": "Yo",
  "descripcion": "Gran depresión de los estudiantes universitarios",
  "status": "pendiente"
}

Respuesta: 
{
    "id": 4,
    "reportero": "Yo",
    "descripcion": "Gran depresión de los estudiantes universitarios",
    "status": "pendiente",
    "fecha": "2025-04-01T17:52:57.3416788-06:00"
}

2. GET 
http://localhost:8083/incidentes

Devuelve todos los incidentes: 
[
    {
        "id": 1,
        "reportero": "Juan Pérez",
        "descripcion": "El servidor no responde a las peticiones.",
        "status": "pendiente",
        "fecha": "2025-03-31T18:47:05.451892-06:00"
    },
    {
        "id": 2,
        "reportero": "Mejía",
        "descripcion": "El servidor no responde a las peticiones.",
        "status": "pendiente",
        "fecha": "2025-03-31T18:47:21.439152-06:00"
    }
]

3. DELETE 
http://localhost:8083/incidentes/3 

Respuesta: 
{
    "message": "Incidente eliminado"
}

4. PUT 
http://localhost:8083/incidentes/2

Cuerpo: 
{
  "status": "completo"
}

Respuesta: 
{
    "message": "Estado actualizado"
}

Autor: Leonardo Mejía 
