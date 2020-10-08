DROP DATABASE healthcheck;
CREATE DATABASE healthcheck;
USE healthcheck;
CREATE TABLE url_info ( id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                        url varchar(225),
                        time int);
