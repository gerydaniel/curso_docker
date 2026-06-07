# Desafío Opcional - Sesión 1

## Imagen elegida
- postgres:16-alpine
- Razón: ligera, estable, buena para laboratorios y práctica con bases de datos.

## Variables de entorno
- POSTGRES_USER=curso
- POSTGRES_PASSWORD=curso123
- POSTGRES_DB=labdb

## Puerto
- Interno: 5432
- Mapeado en host: 5433

## Comandos principales
```bash
docker pull postgres:16-alpine
docker run -d --name lab-postgres -e POSTGRES_USER=curso -e POSTGRES_PASSWORD=curso123 -e POSTGRES_DB=labdb -p 5433:5432 postgres:16-alpine
docker ps --filter "name=lab-"
docker logs lab-postgres
docker rm -f $(docker ps -aq --filter "name=lab-")