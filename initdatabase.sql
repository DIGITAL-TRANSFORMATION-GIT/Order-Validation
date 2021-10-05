-- GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost'
--     IDENTIFIED BY 'password'  
--     WITH GRANT OPTION;
-- FLUSH PRIVILEGES;

CREATE DATABASE orderValidator;
use orderValidator;


CREATE TABLE orders(
    id integer PRIMARY KEY,
    title varchar(50)
);

CREATE TABLE requirements(
    requirementid integer PRIMARY KEY AUTO_INCREMENT,
    request varchar(50),
    expectedoutcome varchar(50),
    order_id integer,
    status bool,
    FOREIGN KEY(order_id) REFERENCES orders(id)
);

