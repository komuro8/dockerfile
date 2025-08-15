-- Crear tabla de prueba
CREATE TABLE IF NOT EXISTS usuarios (
    id SERIAL PRIMARY KEY,
    nombre TEXT NOT NULL
);

-- Insertar datos iniciales
INSERT INTO usuarios (nombre) VALUES ('Juan'), ('María'), ('Pedro');
