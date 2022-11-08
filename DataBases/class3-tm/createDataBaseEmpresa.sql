# creando la base de datos
DROP DATABASE IF EXISTS empresa;
CREATE DATABASE empresa;
USE empresa;

# creando las tablas
DROP TABLE IF EXISTS departamento;
CREATE TABLE departamento (
  depto_nro varchar(15) NOT NULL,
  nombre_depto varchar(32) NOT NULL,
  localidad varchar(32) NOT NULL,
  PRIMARY KEY (depto_nro)
);

DROP TABLE IF EXISTS empleado;
CREATE TABLE empleado (
  cod_emp varchar(15) NOT NULL,
  nombre varchar(100) NOT NULL,
  apellido varchar(100) NOT NULL,
  puesto varchar(100) NOT NULL,
  fecha_alta timestamp NULL DEFAULT NULL,
  salario int(32) NOT NULL,
  comision int(32),
  depto_nro varchar(15) NOT NULL,
  PRIMARY KEY (cod_emp),
  CONSTRAINT empleado_depto_nro_foreign 
  FOREIGN KEY (depto_nro) 
  REFERENCES departamento (depto_nro)
);

# insertando los datos
INSERT INTO departamento(depto_nro, nombre_depto, localidad) VALUES 
('D-000-1', 'Software', 'Los Tigres'),
('D-000-2', 'Sistemas', 'Guadalupe'),
('D-000-3', 'Contabilidad', 'La Roca'),
('D-000-4', 'Ventas', 'Planta');

INSERT INTO empleado(cod_emp, nombre, apellido, puesto, fecha_alta, salario, comision, depto_nro) 
VALUES 
('E-0001', 'César', 'Piñero', 'Vendedor', '2018-05-12 00:01:01', 80000, 15000, 'D-000-4'),
('E-0002', 'Yosep', 'Kowaleski', 'Analista', '2015-07-14 00:01:01', 140000, 0, 'D-000-2'),
('E-0003', 'Mariela', 'Barrios', 'Director', '2014-06-05 00:01:01', 185000, 0, 'D-000-3'),
('E-0004', 'Jonathan', 'Aguilera', 'Vendedor', '2015-06-03 00:01:01', 85000, 10000, 'D-000-4'),
('E-0005', 'Daniel', 'Brezezicki', 'Vendedor', '2018-06-03 00:01:01', 83000, 10000, 'D-000-4'),
('E-0006', 'Mito', 'Barchuk', 'Presidente', '2014-05-05 00:01:01', 190000, 0, 'D-000-3'),
('E-0007', 'Emilio', 'Galarza', 'Desarrollador', '2014-08-02 00:01:01', 60000, 0, 'D-000-1');
