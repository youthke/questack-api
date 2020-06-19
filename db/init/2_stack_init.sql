DROP TABLE IF EXISTS `stacks`;

CREATE TABLE `stacks`(
     `id` VARCHAR(255) NOT NULL,
     `name` VARCHAR (255) NOT NULL ,
     `owner_id`   INT(11) NOT NULL ,
     PRIMARY KEY (`id`),
     FOREIGN KEY (`owner_id`) REFERENCES `owners`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
