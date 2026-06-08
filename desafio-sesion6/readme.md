# Desafío Opcional - Sesión 6

## Objetivo
Levantar un stack de microservicios usando Docker Compose, con volúmenes persistentes y redes dedicadas, mostrando la comunicación entre servicios.

## Servicios
- **api-sesion6**
  - Node.js
  - Puerto: 3001:3000
  - Volumen: api-datos:/app
  - Depende de mongo-sesion6
- **mongo-sesion6**
  - MongoDB 7
  - Puerto: 27019:27017
  - Volumen: mongo-datos:/data/db

## Red
- Nombre: red-sesion6
- Tipo: bridge
- Contenedores conectados: api-sesion6, mongo-sesion6

## Comandos principales
```bash
docker-compose up -d --build
docker ps
docker-compose logs api
docker-compose logs mongo
docker inspect api-sesion6
docker inspect mongo-sesion6
docker network inspect desafio-sesion6_red-sesion6
docker volume ls
docker exec -it api-sesion6 ping mongo-sesion6
docker-compose down -v