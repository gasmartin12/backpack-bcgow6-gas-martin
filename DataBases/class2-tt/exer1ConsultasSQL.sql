# consultas clase 2 tt
select * from movies;
select first_name, last_name, rating from actors;
select title as titulo from series as series;
select title, rating, awards from movies where rating > '7.5' and awards > 2;
select title, rating from movies order by rating;
select title from movies limit 3;
select title, rating, awards, release_date from movies order by rating limit 5;
select first_name, last_name from actors limit 10;
select title, rating from movies where title like '%Toy Story%';
select first_name, last_name from actors where first_name like 'Sam%';
select title from movies where release_date between '2004-01-01' and '2008-01-01';
select title from movies 
	where rating > 3 
		and awards > 1 
        and release_date between '1988-01-01' and '2009-01-01' 
	order by rating;