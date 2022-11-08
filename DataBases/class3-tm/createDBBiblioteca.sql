/*==============================================================*/
/* DBMS name:      Sybase SQL Anywhere 12                       */
/* Created on:     1/11/2022 11:17:57 p.�m.                     */
/*==============================================================*/


DROP DATABASE IF EXISTS biblioteca_db;
CREATE DATABASE biblioteca_db;
USE biblioteca_db;

/*==============================================================*/
/* Table: AUTOR                                                 */
/*==============================================================*/
CREATE TABLE autor 
(
   id_autor             INTEGER                        NOT NULL,
   nombre               LONG VARCHAR                   NOT NULL,
   nacionalidad         LONG VARCHAR                   NOT NULL,
   CONSTRAINT pk_autor PRIMARY KEY (id_autor)
);

/*==============================================================*/
/* Index: AUTOR_PK                                              */
/*==============================================================*/
CREATE UNIQUE INDEX autor_pk ON autor (
id_autor ASC
);

INSERT INTO autor (id_autor, nombre, nacionalidad)
VALUES
  (1, "Judah Levy", "Francia"),
  (2, "Brendan Knox", "Italia"),
  (3, "Zahir Jimenez", "Chile"),
  (4, "Peter Colon", "Belize"),
  (5, "Demetrius Snow", "Nicaragua");

/*==============================================================*/
/* Table: ESTUDIANTE                                            */
/*==============================================================*/
CREATE TABLE estudiante 
(
   id_lector            INTEGER                        NOT NULL,
   nombre               LONG VARCHAR                   NOT NULL,
   apellido             LONG VARCHAR                   NOT NULL,
   direccion            LONG VARCHAR                   NOT NULL,
   carrera              LONG VARCHAR                   NOT NULL,
   edad                 INTEGER                        NOT NULL,
   CONSTRAINT pk_estudiante PRIMARY KEY (id_lector)
);

/*==============================================================*/
/* Index: ESTUDIANTE_PK                                         */
/*==============================================================*/
CREATE UNIQUE INDEX estudiante_pk ON estudiante (
id_lector ASC
);

INSERT INTO estudiante (id_lector, nombre, apellido, direccion, carrera, edad) 
VALUES (1, "Dara", "Williams", "Ap #472-9803 Proin Ave", "Customer Service", 18), 
(2, "Xyla", "Ferguson", "P.O. Box 612, 5517 Diam Ave", "Informatica", 15), 
(3, "Daria", "Roy", "Ap #499-4640 Etiam St.", "Customer Relations", 14), 
(4, "Chava", "Cline", "P.O. Box 720, 840 Blandit. Ave", "Informatica", 21), 
(5, "Kaitlin", "Guzman", "Ap #257-5297 Sem, Road", "Research and Development", 16);

/*==============================================================*/
/* Table: LIBRO                                                 */
/*==============================================================*/
CREATE TABLE libro 
(
   id_libro             INTEGER                        NOT NULL,
   titulo               LONG VARCHAR                   NOT NULL,
   editorial            LONG VARCHAR                   NOT NULL,
   area                 LONG VARCHAR                   NOT NULL,
   CONSTRAINT pk_libro PRIMARY KEY (id_libro)
);

/*==============================================================*/
/* Index: LIBRO_PK                                              */
/*==============================================================*/
CREATE UNIQUE INDEX libro_pk ON libro (
id_libro ASC
);

INSERT INTO libro (id_libro, titulo,editorial,area)
VALUES
  (1, "Fusce dolor quam,", "adipiscing", "Customer Relations"),
  (2, "nunc nulla vulputate", "lectus", "Finances"),
  (3, "sed, facilisis vitae,", "egestas.", "Legal Department"),
  (4, "metus. Aenean sed", "enim", "Customer Service"),
  (5, "ipsum ac mi", "odio", "Tech Support"),
  (6, "libero. Proin sed", "Etiam", "Tech Support"),
  (7, "dolor quam, elementum", "Mauris", "Human Resources"),
  (8, "Etiam imperdiet dictum", "eros.", "Asset Management"),
  (9, "a purus. Duis", "mi.", "Finances"),
  (10, "velit eu sem.", "nisl", "Human Resources"),
  (11, "non massa non", "tempor", "Asset Management"),
  (12, "orci. Ut sagittis", "lacinia.", "Human Resources"),
  (13, "eu dolor egestas", "Donec", "Payroll"),
  (14, "elementum, dui quis", "commodo", "Media Relations"),
  (15, "enim diam vel", "at", "Human Resources");

/*==============================================================*/
/* Table: LIBRO_AUTOR                                           */
/*==============================================================*/
CREATE TABLE libro_autor 
(
   id                   INTEGER                        NOT NULL,
   id_libro             INTEGER                        NOT NULL,
   id_autor             INTEGER                        NOT NULL,
   CONSTRAINT pk_libro_autor PRIMARY KEY (id)
);

/*==============================================================*/
/* Index: LIBRO_AUTOR_PK                                        */
/*==============================================================*/
CREATE UNIQUE INDEX libro_autor_pk ON libro_autor (
id ASC
);

/*==============================================================*/
/* Index: RELATIONSHIP_2_FK                                     */
/*==============================================================*/
CREATE INDEX RELATIONSHIP_2_FK ON LIBRO_AUTOR (
ID_AUTOR ASC
);

/*==============================================================*/
/* Index: RELATIONSHIP_3_FK                                     */
/*==============================================================*/
CREATE INDEX relationship_3_fk ON libro_autor (
id_libro ASC
);

/*==============================================================*/
/* Table: PRESTAMO                                              */
/*==============================================================*/
CREATE TABLE prestamo 
(
   id_prestamo          INTEGER                        NOT NULL,
   id_libro             INTEGER                        NOT NULL,
   id_lector            INTEGER                        NOT NULL,
   fecha_prestamo       DATE                           NOT NULL,
   fecha_devolucion     DATE                           NULL,
   devuelto             SMALLINT                       NULL,
   CONSTRAINT pk_prestamo PRIMARY KEY (id_prestamo)
);

/*==============================================================*/
/* Index: PRESTAMO_PK                                           */
/*==============================================================*/
CREATE UNIQUE INDEX prestamo_pk ON prestamo (
id_prestamo ASC
);

/*==============================================================*/
/* Index: RELATIONSHIP_5_FK                                     */
/*==============================================================*/
CREATE INDEX relationship_5_fk ON prestamo (
id_lector ASC
);

/*==============================================================*/
/* Index: RELATIONSHIP_6_FK                                     */
/*==============================================================*/
CREATE INDEX relationhip_6_fk ON prestamo (
id_libro ASC
);

ALTER TABLE libro_autor
   ADD CONSTRAINT fk_libro_au_relations_autor FOREIGN KEY (id_autor)
      REFERENCES autor (id_autor)
      ON UPDATE restrict
      ON DELETE restrict;

ALTER TABLE libro_autor
   ADD CONSTRAINT fk_libro_au_relations_libro FOREIGN KEY (id_libro)
      REFERENCES libro (id_libro)
      ON UPDATE restrict
      ON DELETE restrict;

ALTER TABLE prestamo
   ADD CONSTRAINT fk_prestamo_relations_estudiante FOREIGN KEY (id_lector)
      REFERENCES estudiante (id_lector)
      ON UPDATE restrict
      ON DELETE restrict;

ALTER TABLE prestamo
   ADD CONSTRAINT fk_prestamo_relations_libro FOREIGN KEY (id_libro)
      REFERENCES libro (id_libro)
      ON UPDATE restrict
      ON DELETE restrict;

INSERT INTO libro_autor (id,id_libro,id_autor)
VALUES (1,2,1), (2,9,1), (3,1,3), (4,9,5), (5,14,5), 
(6,4,3), (7,13,4), (8,4,5), (9,4,2), (10,12,4), 
(11,12,5), (12,8,3), (13,4,3), (14,11,1), (15,14,4);

INSERT INTO prestamo (id_prestamo,id_libro,id_lector,fecha_prestamo,fecha_devolucion,devuelto)
VALUES
  (1,1,5,"2021-11-27","2022-04-08",1),
  (2,10,1,"2022-07-05","2022-10-27",0),
  (3,7,5,"2021-11-19","2022-05-27",0),
  (4,10,3,"2020-11-20","2021-05-09",1),
  (5,13,3,"2022-02-19","2022-05-27",1),
  (6,2,4,"2021-04-27","2021-11-01",1),
  (7,3,4,"2022-09-05","2022-12-10",1),
  (8,10,2,"2021-01-02","2022-10-08",1),
  (9,14,4,"2020-12-06",'2021-07-16',0),
  (10,2,4,"2021-01-10","2022-08-07",1),
  (11,1,2,"2022-07-10","2022-07-19",1),
  (12,11,2,"2021-04-20","2021-08-16",1),
  (13,6,2,"2021-11-08","2022-01-27",0),
  (14,13,2,"2022-08-11","2022-12-27",0),
  (15,7,5,"2022-06-11","2022-09-27",0),
  (16,11,2,"2021-10-13","2021-12-26",1),
  (17,6,2,"2020-01-20","2021-07-16",0),
  (18,6,3,"2021-10-17","2021-12-14",1),
  (19,12,3,"2020-11-27","2020-12-14",1),
  (20,14,3,"2021-09-07","2022-03-17",1),
  (21,5,3,"2021-10-15","2022-05-27",0),
  (22,6,2,"2021-01-25","2022-03-23",1),
  (23,4,4,"2021-12-28","2022-09-06",0),
  (24,15,2,"2022-10-19","2022-12-27",0),
  (25,4,5,"2021-12-02","2022-02-27",0);