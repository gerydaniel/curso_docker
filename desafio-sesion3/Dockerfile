# Dockerfile del desafío 3
FROM python:3.12-alpine

# Información del autor
LABEL maintainer="Gery Rivera <geryriveraf@gmail.com>"

# Crear directorio de trabajo
WORKDIR /app

# Copiar archivos locales
COPY . /app

# Instalar dependencias
RUN pip install --no-cache-dir requests

# Comando por defecto
CMD ["python", "-c", "print('Hola desde el contenedor del desafío 3')"]