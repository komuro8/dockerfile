-- Crear tabla de usuarios
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nombre_usuario TEXT NOT NULL UNIQUE,
    contraseña TEXT NOT NULL,
    imagen_perfil TEXT
);

-- Crear tabla de categorías
CREATE TABLE categorias (
    id SERIAL PRIMARY KEY,
    nombre TEXT NOT NULL,
    descripcion TEXT
);

-- Crear tabla de tareas
CREATE TABLE tareas (
    id SERIAL PRIMARY KEY,
    texto TEXT NOT NULL,
    fecha_creacion TIMESTAMP DEFAULT NOW(),
    fecha_finalizacion DATE,
    estado TEXT CHECK (estado IN ('Sin Empezar', 'Empezada', 'Finalizada')),
    id_categoria INT REFERENCES categorias(id),
    id_usuario INT REFERENCES usuarios(id)
);

-- Insertar usuarios de ejemplo
INSERT INTO usuarios (nombre_usuario, contraseña, imagen_perfil)
VALUES
    ('juan', '1234', NULL),
    ('maria', '1234', NULL);

-- Insertar categorías de ejemplo
INSERT INTO categorias (nombre, descripcion)
VALUES
    ('Trabajo', 'Tareas relacionadas con el trabajo'),
    ('Hogar', 'Tareas domésticas'),
    ('Estudio', 'Tareas y actividades académicas');

-- Insertar tareas de ejemplo
INSERT INTO tareas (texto, fecha_finalizacion, estado, id_categoria, id_usuario)
VALUES
    ('Preparar informe mensual', '2025-08-20', 'Sin Empezar', 1, 1),
    ('Limpiar cocina', '2025-08-16', 'Empezada', 2, 1),
    ('Estudiar para el examen', '2025-08-25', 'Sin Empezar', 3, 2);
