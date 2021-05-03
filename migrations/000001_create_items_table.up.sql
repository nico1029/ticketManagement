CREATE TABLE IF NOT EXISTS registered_tickets(
namei SERIAL PRIMARY KEY, 
id VARCHAR(256), 
usuario VARCHAR(256), 
fechacreacion VARCHAR(256), 
fechaactualizacion VARCHAR(256), 
estatus VARCHAR(256)
);

CREATE TABLE IF NOT EXISTS deleted_tickets(
namei SERIAL PRIMARY KEY, 
id VARCHAR(256), 
usuario VARCHAR(256), 
fechacreacion VARCHAR(256), 
fechaactualizacion VARCHAR(256), 
estatus VARCHAR(256)
);