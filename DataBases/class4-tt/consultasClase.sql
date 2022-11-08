-- Active: 1667417196879@@127.0.0.1@3306@movies_db

# Ejercicio 1
#Con la base de datos "movies", se propone crear una tabla temporal llamada "TWD" y guardar en la misma los episodios de todas las temporadas de "The Walking Dead".
SELECT e.*
FROM episodes e
    INNER JOIN seasons s on e.season_id = s.id
where s.serie_id = (
        SELECT id
        from series
        where
            title = "The Walking Dead"
    );

#Realizar una consulta a la tabla temporal para ver los episodios de la primera temporada.
create TEMPORARY
table `temp_TWD`
SELECT
    e.title,
    e.number,
    e.release_date,
    e.rating,
    e.season_id
FROM episodes e
    INNER JOIN seasons s on e.season_id = s.id
where s.serie_id = (
        SELECT id
        from series
        where
            title = "The Walking Dead"
    );

show CREATE table temp_TWD;

select * from temp_TWD order by id desc;

# Ejercicio 2
# En la base de datos "movies", seleccionar una tabla donde crear un índice y luego chequear la creación del mismo. 
# Analizar por qué crearía un índice en la tabla indicada y con qué criterio se elige/n el/los campos.
select *
from series;

alter table series add UNIQUE INDEX `idx_series_title` (`title`);

select * from series where title = "The Walking Dead";

