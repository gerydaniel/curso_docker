# Desafío Opcional - Sesión 2

## Imagen elegida

Elegí la imagen `mongo:7`.

MongoDB es una base de datos NoSQL orientada a documentos. La imagen permite ejecutar un servidor MongoDB dentro de un contenedor Docker.

## Variables de entorno

Variables usadas:

- `MONGO_INITDB_ROOT_USERNAME=curso`
- `MONGO_INITDB_ROOT_PASSWORD=curso123`

Estas variables crean un usuario root inicial en la base de autenticación `admin`.

## Puerto expuesto

MongoDB expone internamente el puerto:

- `27017`

En este laboratorio se publicó en el host como:

- `27018:27017`

## Directorio de datos persistentes

MongoDB almacena los datos por defecto en:

- `/data/db`

Por eso el volumen nombrado se montó así:

```bash
-v datos-desafio:/data/db