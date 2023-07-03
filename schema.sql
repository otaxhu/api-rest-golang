DROP DATABASE IF EXISTS api_rest_golang;
CREATE DATABASE api_rest_golang;
USE api_rest_golang;

CREATE TABLE movies(
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(80) NOT NULL,
    date DATE,
    cover_url VARCHAR(255) NOT NULL
);