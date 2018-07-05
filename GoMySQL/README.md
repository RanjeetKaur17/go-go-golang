###Create Schema

   create database student;

   use student;

   CREATE TABLE IF NOT EXISTS `student`(
     `id` INT(11) NOT NULL AUTO_INCREMENT,
      `name` VARCHAR(50) DEFAULT NULL,
     `age` INT(11) DEFAULT 0,
     UNIQUE KEY(`id`),
     PRIMARY KEY(`id`)
   ) ENGINE=InnoDB DEFAULT CHARSET=utf8;