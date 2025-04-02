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
   ```sh
   git clone https://github.com/DufreyM/EjercicioAPI.git

2. **Base de datos**
   Debes cambiar la variable dsn con la información de tu base de datos en postgres.
   Debe estar activa y configurada en el puerto 5432

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
