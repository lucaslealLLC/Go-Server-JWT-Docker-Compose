
CREATE DATABASE IF NOT EXISTS `UserDatabase_dev`;
GRANT ALL ON `UserDatabase` .* TO 'root'@'%';

CREATE TABLE IF NOT EXISTS UserDatabase.users (
    id int not null AUTO_INCREMENT Primary Key,
    user varchar(255) not null,
    name varchar(255) not null,
    surname varchar(255) not null,
    createdAt TIMESTAMP DEFAULT NOW(),
	updatedAt TIMESTAMP DEFAULT NOW(),
    UNIQUE(user)
);

CREATE TABLE IF NOT EXISTS UserDatabase_dev.users (
    id int not null AUTO_INCREMENT Primary Key,
    user varchar(255) not null,
    name varchar(255) not null,
    surname varchar(255) not null,
    createdAt TIMESTAMP DEFAULT NOW(),
	updatedAt TIMESTAMP DEFAULT NOW(),
    UNIQUE(user)
);
