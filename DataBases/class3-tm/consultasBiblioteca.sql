
-- 1 Listar los datos de los autores.
select * from autor;
-- 2 Listar nombre y edad de los estudiantes
select nombre, apellido, edad from estudiante;
-- 3 ¿Qué estudiantes pertenecen a la carrera informática?
select nombre, apellido from estudiante
where carrera like 'Informatica';
-- 4 ¿Qué autores son de nacionalidad francesa o italiana?
select * from autor
where nacionalidad like 'Italiana';
-- 5 ¿Qué libros no son del área de internet?
select * from libro
where area not like 'Internet';
-- 6 Listar los libros de la editorial Salamandra.
select * from libro
where editorial not like 'Salamandra';
-- 7 Listar los datos de los estudiantes cuya edad es mayor al promedio.
select * from estudiante
where edad > 
(select avg(edad) from estudiante);
-- 8 Listar los nombres de los estudiantes cuyo apellido comience con la letra G.
select nombre, apellido from estudiante
where apellido like 'G%';
-- 9 Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).
select aut.nombre from autor aut
where aut.idAutor in
(select autlib.idAutor from libroautor autlib
where autlib.idLibro in
(select lib.idLibro from libro lib
where lib.titulo like 'El Universo: Guía de viaje'));

-- 10 ¿Qué libros se prestaron al lector “Filippo Galli”?

-- 11 Listar el nombre del estudiante de menor edad.
-- 12 Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.
-- 13 Listar los libros que pertenecen a la autora J.K. Rowling.
-- 14 Listar títulos de los libros que debían devolverse el 16/07/2021.