Para levantar la imagen docker con la base de datos 

docker run --rm --name my-postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres \
  && docker exec -it my-postgres psql -U postgres -W

\l #para listar nuestras bases de datos; Aqui debe aparecer en la columa Name nuestra dase de datos
\c toys_db_development #Ingresamos a la base de datos utilizando la password configurada en el dockerfile
\dt #lista las tablas en la base de datos
select * from toys; # trae el resultado de todas las filas de la tabla