-- COMANDOS SQL PARA DESENVOLVIMENTO --

-- Modo mais antigo para ser criar um DB e criar uma tabela

-- IDENTITY(1,1) Configura a coluna para auto-incrementar. O primeiro número (1) é o valor inicial, e o segundo número (1) é o incremento.

IF NOT EXISTS (SELECT name FROM sys.databases WHERE name = 'devbook')
BEGIN
    CREATE DATABASE devbook;
END;

IF EXISTS (SELECT * FROM usuarios)
BEGIN
    DROP TABLE usuarios;
END;

CREATE TABLE usuarios(
    id INT IDENTITY(1,1) PRIMARY KEY, 
    nome VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    senha VARCHAR(100) NOT NULL,
    criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Modo mais recente para ser criar um DB e criar uma tabela

CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS seguidores;

CREATE TABLE usuarios(
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    senha VARCHAR(100) NOT NULL,
    criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
) ENGINE=INNODB;

CREATE TABLE seguidores(
    usuario_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    
    seguidor_id int not null,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    primary key(usuario_id, seguidor_id)
) ENGINE=INNODB;


-- ON DELETE CASCADE => Se algum ID da tabela usuario for deletada, todas as referencias dele nesta tabela tb serão
-- primary key(usuario_id, seguidor_id) => criacao da chave primaria com a junção das 2 foreign key.