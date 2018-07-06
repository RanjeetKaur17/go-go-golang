go get -u github.com/go-sql-driver/mysql

#####Create Schema

create database `student`;

use `student`;

CREATE TABLE IF NOT EXISTS `student`(<br/>
    __`id` INT(11) NOT NULL AUTO_INCREMENT,<br/>
    `name` VARCHAR(50) DEFAULT NULL,<br/>
    `age` INT(11) DEFAULT 0,<br/>
    UNIQUE KEY(`id`),<br/>
    PRIMARY KEY(`id`)<br/>
) ENGINE=InnoDB DEFAULT CHARSET=utf8;