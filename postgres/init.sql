CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nombre_usuario TEXT NOT NULL UNIQUE,
    contraseña TEXT NOT NULL,
    imagen_perfil TEXT
);

CREATE TABLE categorias (
    id SERIAL PRIMARY KEY,
    nombre TEXT NOT NULL,
    descripcion TEXT
);

CREATE TABLE tareas (
    id SERIAL PRIMARY KEY,
    texto TEXT NOT NULL,
    fecha_creacion TIMESTAMP DEFAULT NOW(),
    fecha_finalizacion DATE,
    estado TEXT CHECK (estado IN ('Sin Empezar', 'Empezada', 'Finalizada')),
    id_categoria INT REFERENCES categorias(id),
    id_usuario INT REFERENCES usuarios(id)
);
