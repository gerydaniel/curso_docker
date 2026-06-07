# Desafío Opcional - Sesión 3

## Imagen base

- python:3.12-alpine

## Dockerfile

- WORKDIR: /app
- COPY: . /app
- RUN: pip install requests
- CMD: ["python", "-c", "print('Hola desde el contenedor del desafío 3')"]

## Construcción de la imagen

- Comando: docker build -t desafio3:1.0 .
- Capas creadas: (captura de docker history)

## Ejecución

- Comando: docker run --name contenedor-desafio3 desafio3:1.0
- Salida: "Hola desde el contenedor del desafío 3"

## Inspección

- docker inspect, docker logs, docker top → capturas adjuntas.

## Limpieza

- docker rm -f contenedor-desafio3
- docker rmi desafio3:1.0