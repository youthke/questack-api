DROP TABLE IF EXISTS `owners`;

CREATE TABLE `owners`(
     `id` INT(11) NOT NULL AUTO_INCREMENT,
     `name` VARCHAR (255) NOT NULL ,
     `mail` VARCHAR(255) NOT NULL UNIQUE,
     `password` VARCHAR(255) NOT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
