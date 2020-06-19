DROP TABLE IF EXISTS `questions`;

CREATE TABLE `questions`(
     `id` INT(11) NOT NULL AUTO_INCREMENT,
     `title` VARCHAR (255) NOT NULL ,
     `author`   VARCHAR (255) NOT NULL ,
     `content` VARCHAR (255) NOT NULL,
     `stack_id` VARCHAR (255) NOT NULL ,
     `deletedAt` DATETIME DEFAULT NULL ,
     PRIMARY KEY (`id`),
     FOREIGN KEY (`stack_id`) REFERENCES `stacks`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
