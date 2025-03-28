# Football Matches API

Este proyecto es un backend en Go que gestiona información sobre partidos de fútbol, utilizando PostgreSQL como base de datos.
Cómo ejecutar el proyecto

Para correr la API, simplemente usa Docker Compose:
```bash
docker compose up --build
```
Esto levantará tanto el backend como la base de datos sin necesidad de configuraciones adicionales.

## Documentación de la API

La API está documentada con Swagger. Una vez que el backend está corriendo, puedes acceder a la documentación interactiva en:
URL local:
http://localhost:8080/swagger/index.html

Endpoints principales

    GET /matches → Obtiene todos los partidos

    GET /matches/{id} → Obtiene un partido por ID

    POST /matches → Crea un nuevo partido

    PUT /matches/{id} → Actualiza un partido

    PATCH /matches/{id} → Actualiza parcialmente un partido

    DELETE /matches/{id} → Elimina un partido

Tecnologías utilizadas

    Go (Gorilla Mux para routing)

    PostgreSQL (almacenamiento de datos)

    Docker & Docker Compose (contenedores)

    Swagger (documentación de API)

Endpoints PATCH

Estos endpoints permiten actualizar atributos específicos de un partido.

    PATCH /matches/{id}/goals → Registra un nuevo gol en el partido

    PATCH /matches/{id}/yellowcards → Registra una tarjeta amarilla

    PATCH /matches/{id}/redcards → Registra una tarjeta roja

    PATCH /matches/{id}/extratime → Establece el tiempo extra del partido
    
## Funcionamiento de las 5 principales rutas:
Metodo de GET
![imagen](https://github.com/user-attachments/assets/ce538478-f1ba-4ddd-9c59-6dfd03f008ca)

Crear partido: 
![imagen](https://github.com/user-attachments/assets/5226cf5b-aebc-4b8d-abef-7670b7673924)
Muestra de que funciono: 
![imagen](https://github.com/user-attachments/assets/16a4969b-1d58-498c-b6ea-10978a1450d8)


Buscar partido por ID
![imagen](https://github.com/user-attachments/assets/43eb5091-4fe3-4681-a53b-b29546576699)

Actualizar partido: el mismo que cree: 
![imagen](https://github.com/user-attachments/assets/f067666e-2fb2-4e5c-884a-8323b75821ff)

Prueba de que funciona: 
![imagen](https://github.com/user-attachments/assets/ecb96a3b-a37b-4f9b-ba95-1022e325f41e)


Eliminar partido por ID
![imagen](https://github.com/user-attachments/assets/0cb0ad87-e61f-4ed6-8e70-8bd7dba17be7)

## Funcionamiento rutas adicionales: 
Registro de gol:
![imagen](https://github.com/user-attachments/assets/e19d26c7-b8cc-4222-9021-838676379e3d)

Registro de tarjeta amarilla:
![imagen](https://github.com/user-attachments/assets/c276fcb0-a762-4c8e-8443-ed15ee5daaa7)

Registro de tarjeta roja: 
![imagen](https://github.com/user-attachments/assets/687a2712-e18c-4dd5-ba6b-3844f953faff)

Registro de tiempo extra:
![imagen](https://github.com/user-attachments/assets/b46b3476-3433-4ca1-a1ea-e76eb985966b)

Documentacion del API:
![imagen](https://github.com/user-attachments/assets/fce2ce79-e9d0-421c-8743-12da1c271ed5)

