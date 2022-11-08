SELECT * FROM movies_db.movies;

SELECT * FROM actor_movie;


#INNER JOIN
SELECT m.id, m.title, a.first_name, a.last_name, a.favorite_movie_id 
FROM movies m
INNER JOIN actors a 
ON m.id = a.favorite_movie_id;

SELECT m.id, m.title,
CONCAT(a.first_name," ", a.last_name) AS "nombre y apellido", a.favorite_movie_id 
FROM movies m
INNER JOIN actors a 
ON m.id = a.favorite_movie_id;

#LEFT JOIN
SELECT m.title AS "Pelicula Favorita", CONCAT(a.first_name," ", a.last_name) AS "Actor"
FROM movies m
LEFT JOIN actors a ON m.id = a.favorite_movie_id;

SELECT m.title AS "Pelicula Favorita", CONCAT(a.first_name," ", a.last_name) AS "Actor"
FROM movies m
LEFT JOIN actors a ON m.id = a.favorite_movie_id
WHERE a.favorite_movie_id IS NULL;

#RIGHT JOIN 
SELECT m.title as "Pelicula Favorita", CONCAT(a.first_name," ", a.last_name) AS "Actor"
FROM movies m
RIGHT JOIN actors a ON m.id = a.favorite_movie_id;

SELECT m.title as "Pelicula Favorita", CONCAT(a.first_name," ", a.last_name) AS "Actor"
FROM movies m
RIGHT JOIN actors a ON m.id = a.favorite_movie_id
WHERE a.favorite_movie_id IS NULL;

#Otras consultas
SELECT * FROM movies;

SELECT m.title, SUM(m.awards) AS total_awards
FROM movies m
GROUP BY id HAVING total_awards  > 2;

#SUBCONSULTAS

SELECT * FROM actor_movie am WHERE am.movie_id IN (SELECT am.id FROM actor_movie am WHERE am.actor_id = 1);


SELECT title, rating FROM movies_db.movies
WHERE title LIKE '%Toy Story';