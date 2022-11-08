# creando la base de datos
DROP DATABASE IF EXISTS empresa_internet;
CREATE DATABASE empresa_internet;
USE empresa_internet;

# creando las tablas
DROP TABLE IF EXISTS internet_pack;
CREATE TABLE internet_pack (
  plan_id int(10) NOT NULL,
  speed decimal(3, 2) NOT NULL,
  price decimal(3, 2) NOT NULL,
  PRIMARY KEY (plan_id)
);

DROP TABLE IF EXISTS customer;
CREATE TABLE customer (
  customer_id int(10) NOT NULL,
  first_name varchar(100) NOT NULL,
  last_name varchar(100) NOT NULL,
  date_of_birth timestamp NULL DEFAULT NULL,
  city varchar(100) NOT NULL,
  plan_id int(10) NOT NULL,
  PRIMARY KEY (`customer_id`),
  CONSTRAINT customer_internet_pack_foreign FOREIGN KEY (plan_id) REFERENCES internet_pack (plan_id)
);

# insertando los datos
INSERT INTO internet_pack (plan_id, speed, price) VALUES 
(1234, 3.5, 8.8),
(1235, 3.4, 8.7),
(1236, 3.3, 8.5),
(1237, 3.2, 8.6),
(1238, 3.1, 8.9)
;

INSERT INTO customer (customer_id, first_name, last_name, date_of_birth, city, plan_id) VALUES 
(34546, 'Aylen', 'Tejas', '2008-12-01 00:01:01', 'California', 1234),
(34547, 'Aiser', 'Beza', '2008-10-01 00:02:01', 'Colorado', 1235),
(34548, 'Ai', 'Soda', '2008-03-01 00:04:01', 'Madrid', 1236),
(34549, 'Lola', 'Mento', '2008-04-01 00:03:01', 'Cordoba', 1237),
(34556, 'Seko', 'Me', '2008-05-01 00:06:01', 'Paz', 1238),
(34557, 'Laper', 'Dida', '2008-06-01 00:09:01', 'Valley', 1234),
(34558, 'Eldo', 'Mingo', '2008-07-01 00:10:01', 'LA', 1235),
(34559, 'Aiko', 'Mida', '2008-09-01 00:11:01', 'Quito', 1236),
(34551, 'Ayse', 'Ayse', '2008-08-01 00:50:01', 'Patagonia', 1237),
(34552, 'Ayse', 'Tilla', '2008-11-01 00:40:01', 'Estonia', 1238)
;