SELECT * FROM internet_pack;

SELECT * FROM customer;

SELECT first_name, last_name FROM customer WHERE plan_id = 1234;

SELECT first_name, last_name FROM customer WHERE city = 'Bogota';

SELECT customer_id, first_name, last_name, date_of_birth 
FROM customer ORDER BY date_of_birth ASC;

SELECT plan_id, discount FROM internet_pack WHERE discount = 0.10;

SELECT plan_id, price FROM internet_pack ORDER BY price DESC LIMIT 3;