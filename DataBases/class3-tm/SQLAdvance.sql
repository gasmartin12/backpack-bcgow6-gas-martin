# actica SQL avanzado
#¿A qué se denomina JOIN en una base de datos y para qué se utiliza?
# Se utiliza para obtener datos de varias tablas relacionadas entre sí. Consiste en combinar datos de una tabla con datos de la otra tabla, a partir de una o varias condiciones en común.

#Explicar dos tipos de JOIN.
# Inner Join se utiliza para traer los datos relacionados de dos o más tablas.
# Left Join se utiliza para traer los datos de la tabla izquierda más los relacionados de la tabla derecha.

#¿Para qué se utiliza el GROUP BY?
# Agrupa los resultados según las columnas indicadas. 
# Genera un solo registro por cada grupo de filas que compartan las columnas indicadas.
# Reduce la cantidad de filas de la consulta.
# Se suele utilizar en conjunto con funciones de agregación, para obtener datos resumidos y agrupados por las columnas que se necesiten.

#¿Para qué se utiliza el HAVING? 
# La cláusula HAVING se utiliza para incluir condiciones con algunas funciones SQL.
# Afecta a los resultados traidos por Group By.
# Escribir una consulta genérica para cada uno de los siguientes diagramas:

# PARTE 2
# Se propone realizar las siguientes consultas a la base de datos movies_db.sql trabajada en la primera clase.

# Mostrar el título y el nombre del género de todas las series.
SELECT movies.title AS movie_title, genres.name AS genre_name FROM movies
INNER JOIN genres ON movies.genre_id = genres.id;

# Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT episodes.title, actors.first_name, actors.last_name FROM episodes
LEFT OUTER JOIN actor_episode ON episodes.id = actor_episode.episode_id
LEFT OUTER JOIN actors ON actor_episode.actor_id = actors.id;

# Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT series.title AS serie_title, COUNT(seasons.id) AS total_seasons FROM seasons
INNER JOIN series ON seasons.serie_id = series.id
GROUP BY series.title;

# Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT genres.name AS genre_name, COUNT(movies.id) AS total_movies FROM movies
INNER JOIN genres ON movies.genre_id = genres.id
GROUP BY genres.name
HAVING total_movies >= 3;

# Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
SELECT DISTINCT(CONCAT(actors.first_name, ' ', actors.last_name)) fullname FROM actors
LEFT OUTER JOIN actor_movie ON actors.id = actor_movie.actor_id
LEFT OUTER JOIN movies ON actor_movie.movie_id = movies.id
WHERE movies.title LIKE '%Guerra de las galaxias%';