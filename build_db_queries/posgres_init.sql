
CREATE TYPE edad_t AS ENUM('Cria','Juvenil','Adulta') ;
CREATE TYPE genero_t as ENUM('M','H');
CREATE SCHEMA IF NOT EXISTS productos;
CREATE SCHEMA IF NOT EXISTS info_reptiles;
CREATE TABLE productos.reptiles (
  id  SERIAL,
  name varchar(64) NOT NULL,
  regularPrice DECIMAL(7,2),
  price DECIMAL(7,2) NOT NULL,
  age edad_t DEFAULT 'Cria',
  description varchar(256),
  genre genero_t DEFAULT 'M',

  PRIMARY KEY(id)
);

CREATE TYPE genetica_t as ENUM('Recesivo','Dominante','Codominante','Supercodominante','Otro');
CREATE TABLE info_reptiles.geneticas(
  id  SERIAL,
  genName varchar(64) NOT NULL,
  tipo genetica_t NOT NULL,
  color char(6),
  PRIMARY KEY(id)
);

CREATE TABLE info_reptiles.reptiles_geneticas (
  reptileID integer,
  geneticID integer,
  FOREIGN KEY(reptileID) REFERENCES productos.reptiles(id),
  FOREIGN KEY(geneticID) REFERENCES info_reptiles.geneticas(id) ,
  UNIQUE (reptileID,geneticID)
);

CREATE TABLE productos.fotos_reptiles(
  id  SERIAL,
  img BYTEA NOT NULL,
  PRIMARY KEY(id),
  FOREIGN KEY(id) REFERENCES productos.reptiles(id)
);


