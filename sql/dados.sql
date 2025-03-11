insert into usuarios (nome, nick, email, senha)
values
("Usuário 1", "usuario1", "usuario1@gmail.com", "$2a$10$71idyMeaEuswLxMJuMmjkOCuW3ZqRNKAKpLYwD0k2hmuoDao.gEkS"),
("Usuário 2", "usuario2", "usuario2@gmail.com", "$2a$10$71idyMeaEuswLxMJuMmjkOCuW3ZqRNKAKpLYwD0k2hmuoDao.gEkS"),
("Usuário 3", "usuario3", "usuario3@gmail.com", "$2a$10$71idyMeaEuswLxMJuMmjkOCuW3ZqRNKAKpLYwD0k2hmuoDao.gEkS");
--$2a$10$71idyMeaEuswLxMJuMmjkOCuW3ZqRNKAKpLYwD0k2hmuoDao.gEkS => Hash para senha 123456,

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2), 
(3, 1),
(1, 3);


--(1, 2), usuario1 é seguido pelo 2
--(3, 1), usuario3 é seguido o usuario 1