# Desafío Opcional - Sesión 5

## Objetivo
Levantar un stack multi-contenedor con Docker Compose, usando redes dedicadas y volúmenes persistentes.

## Servicios
- **mongo-desafio5**
  - Imagen: mongo:7
  - Puerto: 27018:27017
  - Volumen: mongo-datos:/data/db
  - Usuario: curso/curso123
- **redis-desafio5**
  - Imagen: redis:8-alpine
  - Puerto: 6379:6379
  - Volumen: redis-datos:/data

## Red
- Nombre: red-desafio5
- Tipo: bridge
- Contenedores conectados: mongo-desafio5, redis-desafio5

## Comandos principales
docker-compose up -d
docker ps
docker-compose logs mongo
docker-compose logs redis
docker inspect mongo-desafio5
docker inspect redis-desafio5
docker network inspect desafio-sesion5_red-desafio5
docker volume ls
docker-compose down -v