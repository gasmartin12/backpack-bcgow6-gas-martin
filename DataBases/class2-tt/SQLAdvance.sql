# ¿A qué se denomina JOIN en una base de datos y para qué se utiliza?
# Se utiliza para obtener datos de varias tablas relacionadas entre sí. 
#
# Explicar dos tipos de JOIN.
# Inner Join se utiliza para traer los datos de la intersección de dos o más tablas.
# Left Join se utiliza para traer los datos producidos por la union de la tabla izquierda más los relacionados 
# de la tabla derecha.
#
# ¿Para qué se utiliza el GROUP BY?
# Agrupa los resultados según las columnas indicadas, generando un solo registro por cada grupo de filas que compartan
# las columnas indicadas. Asimismo, reduce la cantidad de filas de la consulta.
#
# ¿Para qué se utiliza el HAVING? 
# HAVING permite filtrar los resultados traídos por el GROUP BY.
#
# Consultas genéricas inner y left joins
# Inner Join: SELECT <cols> FROM tableA ta INNER JOIN tableB tb ON ta.id = tb.id 
# Left Join: SELECT <cols> FROM tableA ta LEFT JOIN tableB tb ON ta.id = tb.id

-- Ejercicio 1
SELECT movies.title AS movie_title, genres.name AS genre_name FROM movies
INNER JOIN genres ON movies.genre_id = genres.id;

-- Ejercicio 2
SELECT episodes.title, actors.first_name, actors.last_name FROM episodes
LEFT OUTER JOIN actor_episode ON episodes.id = actor_episode.episode_id
LEFT OUTER JOIN actors ON actor_episode.actor_id = actors.id;

-- Ejercicio 3
SELECT series.title AS serie_title, COUNT(seasons.id) AS total_seasons FROM seasons
INNER JOIN series ON seasons.serie_id = series.id
GROUP BY series.title;

-- Ejercicio 4
SELECT genres.name AS genre_name, COUNT(movies.id) AS total_movies FROM movies
INNER JOIN genres ON movies.genre_id = genres.id
GROUP BY genres.name
HAVING total_movies >= 3;

-- Ejercicio 5
SELECT DISTINCT(CONCAT(actors.first_name, ' ', actors.last_name)) fullname FROM actors
LEFT OUTER JOIN actor_movie ON actors.id = actor_movie.actor_id
LEFT OUTER JOIN movies ON actor_movie.movie_id = movies.id
WHERE movies.title LIKE '%Guerra de las galaxias%';