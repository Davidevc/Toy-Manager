# Usa la imagen oficial de PostgreSQL como base
FROM postgres:latest

# Variables de entorno para configurar la contraseña de PostgreSQL
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD password
ENV POSTGRES_DB toys_db_development

# Copia el script SQL para inicializar la base de datos
COPY init.sql /docker-entrypoint-initdb.d/
CMD ["postgres"]