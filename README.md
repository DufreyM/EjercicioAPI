## Ingreso
Utilicé Java con SpringBoot por lo que es necesario tener Java instalado y bien configurada la JVM. Además de tener el puerto 8083 que esta configurado libre. Utilicé Spring Initializer

## Ejemplo de uso 

Este es un POST EN 
http://localhost:8083/incidents 
{
    "reporter": "Juan Perez",
    "description": "La impresora no funciona correctamente.",
    "status": "pendiente"
}

PUT para actualizar el status del 1 
http://localhost:8083/incidents/1 
{
    "status": "Completado"
}
 
GET 
http://localhost:8083/incidents 
No tiene body. 

DELETE 
http://localhost:8083/incidents/2 
No tiene body porque solo borra el incidente por su id. 