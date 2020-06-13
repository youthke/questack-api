DROP TABLE IF EXISTS `questions`;

CREATE TABLE `questions`(
     `id` INT(11) NOT NULL AUTO_INCREMENT,
     `title` VARCHAR (255) NOT NULL ,
     `author`   VARCHAR (255) NOT NULL ,
     `content` VARCHAR (255) NOT NULL,
     `stack_id` INT(255) NOT NULL ,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
