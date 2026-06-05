-- Очистка перед вставкой (если нужно начать чисто)
TRUNCATE TABLE movie_actor, budget_and_fees, movies, actors RESTART IDENTITY CASCADE;

-- Актёры (25 штук)
INSERT INTO actors (first_name, last_name, birth_date, salary) VALUES
('Leonardo',    'DiCaprio',       '1974-11-11', 20000000),
('Joseph',      'Gordon-Levitt',  '1981-02-17',  5000000),
('Elliot',      'Page',           '1987-02-21',  4000000),
('Tom',         'Hardy',          '1977-09-15',  8000000),
('Ken',         'Watanabe',       '1959-10-21',  3000000),
('Christian',   'Bale',           '1974-01-30', 15000000),
('Heath',       'Ledger',         '1979-04-04', 10000000),
('Aaron',       'Eckhart',        '1968-03-12',  5000000),
('Gary',        'Oldman',         '1958-03-21',  7000000),
('Maggie',      'Gyllenhaal',     '1977-11-16',  4000000),
('Matthew',     'McConaughey',    '1969-11-04', 15000000),
('Anne',        'Hathaway',       '1982-11-12', 10000000),
('Jessica',     'Chastain',       '1977-03-24',  8000000),
('Michael',     'Caine',          '1933-03-14',  6000000),
('Mackenzie',   'Foy',            '2000-11-10',  1000000),
('Keanu',       'Reeves',         '1964-09-02', 12000000),
('Carrie-Anne', 'Moss',           '1967-08-21',  5000000),
('Laurence',    'Fishburne',      '1961-07-30',  6000000),
('Hugo',        'Weaving',        '1960-04-04',  5000000),
('Joe',         'Pantoliano',     '1951-09-12',  2000000),
('John',        'Travolta',       '1954-02-18',  8000000),
('Uma',         'Thurman',        '1970-04-29',  7000000),
('Samuel L.',   'Jackson',        '1948-12-21', 10000000),
('Bruce',       'Willis',         '1955-03-19',  9000000),
('Harvey',      'Keitel',         '1939-05-13',  5000000);

-- Фильмы
INSERT INTO movies (title, producer, director, release_year) VALUES
('Inception',        'Emma Thomas',     'Christopher Nolan',  2010),
('The Dark Knight',  'Emma Thomas',     'Christopher Nolan',  2008),
('Interstellar',     'Emma Thomas',     'Christopher Nolan',  2014),
('The Matrix',       'Joel Silver',     'Lana Wachowski',     1999),
('Pulp Fiction',     'Lawrence Bender', 'Quentin Tarantino',  1994);

-- Бюджет и сборы
INSERT INTO budget_and_fees (id_movie, total_budget, fees_in_prod_country, fees_in_other) VALUES
(1, 160000000,  292576195,  543000000),
(2, 185000000,  534858444,  469700000),
(3, 165000000,  188020017,  487000000),
(4,  63000000,  171479930,  292037453),

(5,   8000000,  107928762,  106100000);

-- Актёры фильмов
-- Inception (id=1): DiCaprio, Gordon-Levitt, Page, Hardy, Watanabe
INSERT INTO movie_actor (id_movie, id_actor) VALUES
(1, 1), (1, 2), (1, 3), (1, 4), (1, 5);

-- The Dark Knight (id=2): Bale, Ledger, Eckhart, Oldman, Gyllenhaal
INSERT INTO movie_actor (id_movie, id_actor) VALUES
(2, 6), (2, 7), (2, 8), (2, 9), (2, 10);

-- Interstellar (id=3): McConaughey, Hathaway, Chastain, Caine, Foy
INSERT INTO movie_actor (id_movie, id_actor) VALUES
(3, 11), (3, 12), (3, 13), (3, 14), (3, 15);

-- The Matrix (id=4): Reeves, Moss, Fishburne, Weaving, Pantoliano
INSERT INTO movie_actor (id_movie, id_actor) VALUES
(4, 16), (4, 17), (4, 18), (4, 19), (4, 20);

-- Pulp Fiction (id=5): Travolta, Thurman, Jackson, Willis, Keitel
-- + DiCaprio и Caine повторяются как пасхалка
INSERT INTO movie_actor (id_movie, id_actor) VALUES
(5, 21), (5, 22), (5, 23), (5, 24), (5, 25);
