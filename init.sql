-- init.sql
CREATE TABLE toys (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(50) NOT NULL,
    color VARCHAR(50),
    size VARCHAR(50),
    material VARCHAR(50),
    location VARCHAR(100),
    status VARCHAR(50)
);
INSERT INTO toys (name, type, color, size, material, location, status) VALUES 
    ('Muñeca', 'Muñeca', 'Rosa', 'Pequeña', 'Plástico', 'Cuarto de juegos', 'Nuevo'),
    ('Carro de Carreras', 'Auto', 'Rojo', 'Mediano', 'Metal', 'Caja de juguetes', 'Usado'),
    ('Bloques de Construcción', 'Bloques', 'Varios', 'Grande', 'Plástico', 'Sala de estar', 'Nuevo'),
    ('Pelota', 'Pelota', 'Azul', 'Pequeña', 'Goma', 'Patio trasero', 'Usado');
