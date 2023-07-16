# api-rest-golang

## Sumario

Proyecto de un ejemplo de API RESTful escrito en el lenguaje de programación Go (Golang).
La API REST expone unos endpoints para poder hacer operaciones CRUD a una interfaz o API de repositorio a implementar, en este ejemplo la API de repositorio esta implementada en MySQL pero pudiera ser otra implementación. Los recursos que se pueden consumir son de peliculas.

## Endpoints

1. /movies?page=\[n° de página\] \[GET\] \[Obtener una lista de 5 peliculas, se puede navegar de 5 en 5 peliculas a través del query param "page"\]
2. /movies \[POST\] \[Subir una pelicula en formato "multipart/form-data"\]
3. /movies/\[id\] \[GET\] \[Obtener una pelicula a través de su "id"\]
4. /movies/\[id\] \[DELETE\] \[Eliminar una pelicula a través de su "id"\]
5. /movies/\[id\] \[PUT\] \[Actualizar una pelicula a través de su "id", tiene que ser en formato "multipart/form-data"\]

## Estado actual del proyecto

El proyecto lo considero terminado, hace su funcionalidad exitosamente.
Haré otra branch en donde implementaré la API de repositorio en MongoDB para poder probar que a través de la arquitectura hexagonal se pueden hacer aplicaciones escalables y flexibles a cambios.