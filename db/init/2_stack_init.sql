DROP TABLE IF EXISTS `stacks`;

CREATE TABLE `stacks`(
     `id` INT(11) NOT NULL AUTO_INCREMENT,
     `name` VARCHAR (255) NOT NULL ,
     `owner_id`   INT(11) NOT NULL ,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
