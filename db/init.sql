CREATE USER 'sqluser'@'%' IDENTIFIED BY 'sqlpass' ;
GRANT ALL ON *.* TO 'sqluser'@'%' WITH GRANT OPTION ;

CREATE DATABASE `credentials`;

USE credentials;

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `firstname` varchar(80) DEFAULT NULL,
  `surname` varchar(80) DEFAULT NULL,
  `country` varchar(40) DEFAULT NULL,
  `email` varchar(80) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


