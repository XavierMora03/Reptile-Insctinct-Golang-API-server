
CREATE TYPE edad_t AS ENUM('Cria','Juvenil','Adulta') ;
CREATE TYPE genero_t as ENUM('M','H');
CREATE TABLE reptiles (
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
CREATE TABLE brgvlzm6ifhza3x0u4rs.geneticas(
  id  SERIAL,
  genName varchar(64) NOT NULL,
  tipo genetica_t NOT NULL,
  color char(6),
  PRIMARY KEY(id)
);

CREATE TABLE brgvlzm6ifhza3x0u4rs.reptiles_geneticas (
  reptileID integer,
  geneticID integer,
  FOREIGN KEY(reptileID) REFERENCES brgvlzm6ifhza3x0u4rs.reptiles(id),
  FOREIGN KEY(geneticID) REFERENCES brgvlzm6ifhza3x0u4rs.geneticas(id) ,
  UNIQUE (reptileID,geneticID)
);

CREATE TABLE brgvlzm6ifhza3x0u4rs.fotos_reptiles(
  id  SERIAL,
  img MEDIUMBLOB NOT NULL,
  PRIMARY KEY(id),
  FOREIGN KEY(id) REFERENCES brgvlzm6ifhza3x0u4rs.reptiles(id)
);


