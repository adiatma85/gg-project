-- [DML] Insert Dummy Data for User and for Category Tables
INSERT INTO `user` (`email`, `username`, `password`, `display_name`)
VALUES 
('adiatma85@gmail.com', 'adiatma85', '$2a$10$hEU84tig1.W0TcoSe5.zwushMbDarsTnaXadC5/Y/difKiatAHuGO', 'Luki');

INSERT INTO `category` (`name`)
VALUES 
('Work'),
('Hobby'),
('School');
