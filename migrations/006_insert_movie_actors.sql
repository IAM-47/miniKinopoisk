-- =====================================================
-- Связи актёров с фильмами (реальные касты)
-- Запускать после 003_insert_actors.sql и 007_insert_got_actors.sql
-- =====================================================

-- Inception (2010)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'Inception' AND (a.first_name, a.last_name) IN (
    ('Leonardo',  'DiCaprio'),
    ('Joseph',    'Gordon-Levitt'),
    ('Elliot',    'Page'),
    ('Tom',       'Hardy'),
    ('Ken',       'Watanabe'),
    ('Michael',   'Caine'),
    ('Cillian',   'Murphy')
) ON CONFLICT DO NOTHING;

-- The Dark Knight (2008)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'The Dark Knight' AND (a.first_name, a.last_name) IN (
    ('Christian', 'Bale'),
    ('Heath',     'Ledger'),
    ('Aaron',     'Eckhart'),
    ('Gary',      'Oldman'),
    ('Maggie',    'Gyllenhaal'),
    ('Michael',   'Caine'),
    ('Morgan',    'Freeman')
) ON CONFLICT DO NOTHING;

-- Interstellar (2014)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'Interstellar' AND (a.first_name, a.last_name) IN (
    ('Matthew',   'McConaughey'),
    ('Anne',      'Hathaway'),
    ('Jessica',   'Chastain'),
    ('Michael',   'Caine'),
    ('Mackenzie', 'Foy'),
    ('Matt',      'Damon'),
    ('Casey',     'Affleck')
) ON CONFLICT DO NOTHING;

-- The Matrix (1999)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'The Matrix' AND (a.first_name, a.last_name) IN (
    ('Keanu',       'Reeves'),
    ('Carrie-Anne', 'Moss'),
    ('Laurence',    'Fishburne'),
    ('Hugo',        'Weaving'),
    ('Joe',         'Pantoliano')
) ON CONFLICT DO NOTHING;

-- Pulp Fiction (1994)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'Pulp Fiction' AND (a.first_name, a.last_name) IN (
    ('John',      'Travolta'),
    ('Uma',       'Thurman'),
    ('Samuel L.', 'Jackson'),
    ('Bruce',     'Willis'),
    ('Harvey',    'Keitel'),
    ('Tim',       'Roth')
) ON CONFLICT DO NOTHING;

-- Game of Thrones (2011)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'Game of Thrones' AND (a.first_name, a.last_name) IN (
    ('Emilia',   'Clarke'),
    ('Kit',      'Harington'),
    ('Peter',    'Dinklage'),
    ('Lena',     'Headey'),
    ('Nikolaj',  'Coster-Waldau'),
    ('Sophie',   'Turner'),
    ('Maisie',   'Williams'),
    ('Jason',    'Momoa'),
    ('Charles',  'Dance'),
    ('Iain',     'Glen')
) ON CONFLICT DO NOTHING;

-- Fight Club (1999)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'Fight Club' AND (a.first_name, a.last_name) IN (
    ('Brad',    'Pitt'),
    ('Harvey',  'Keitel'),
    ('Meat',    'Loaf')
) ON CONFLICT DO NOTHING;

-- Gone Girl (2014)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'Gone Girl' AND (a.first_name, a.last_name) IN (
    ('Ben',      'Affleck'),
    ('Tyler',    'Perry'),
    ('Neil',     'Patrick Harris')
) ON CONFLICT DO NOTHING;

-- The Social Network (2010)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'The Social Network' AND (a.first_name, a.last_name) IN (
    ('Elliot',    'Page'),
    ('Justin',    'Timberlake'),
    ('Jeremy',    'Renner')
) ON CONFLICT DO NOTHING;

-- Mr. & Mrs. Smith (2005)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'Mr. & Mrs. Smith' AND (a.first_name, a.last_name) IN (
    ('Brad',     'Pitt'),
    ('Angelina', 'Jolie'),
    ('Vince',    'Vaughn')
) ON CONFLICT DO NOTHING;

-- Ocean's Eleven (2001)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'Ocean''s Eleven' AND (a.first_name, a.last_name) IN (
    ('George',   'Clooney'),
    ('Brad',     'Pitt'),
    ('Matt',     'Damon'),
    ('Andy',     'Garcia'),
    ('Julia',    'Roberts'),
    ('Casey',    'Affleck'),
    ('Bernie',   'Mac')
) ON CONFLICT DO NOTHING;

-- The Departed (2006)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'The Departed' AND (a.first_name, a.last_name) IN (
    ('Leonardo', 'DiCaprio'),
    ('Matt',     'Damon'),
    ('Mark',     'Ruffalo'),
    ('Martin',   'Sheen'),
    ('Alec',     'Baldwin')
) ON CONFLICT DO NOTHING;

-- The Wolf of Wall Street (2013)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'The Wolf of Wall Street' AND (a.first_name, a.last_name) IN (
    ('Leonardo', 'DiCaprio'),
    ('Margot',   'Robbie'),
    ('Matthew',  'McConaughey'),
    ('Jonah',    'Hill'),
    ('Kyle',     'Chandler')
) ON CONFLICT DO NOTHING;

-- Joker (2019)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'Joker' AND (a.first_name, a.last_name) IN (
    ('Joaquin',  'Phoenix'),
    ('Robert',   'De Niro'),
    ('Zazie',    'Beetz'),
    ('Frances',  'Conroy')
) ON CONFLICT DO NOTHING;

-- Marriage Story (2019)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'Marriage Story' AND (a.first_name, a.last_name) IN (
    ('Scarlett', 'Johansson'),
    ('Adam',     'Driver'),
    ('Laura',    'Dern'),
    ('Alan',     'Alda'),
    ('Ray',      'Liotta')
) ON CONFLICT DO NOTHING;

-- Barbie (2023)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'Barbie' AND (a.first_name, a.last_name) IN (
    ('Margot',   'Robbie'),
    ('Ryan',     'Gosling'),
    ('America',  'Ferrera'),
    ('Kate',     'McKinnon'),
    ('Issa',     'Rae'),
    ('Simu',     'Liu')
) ON CONFLICT DO NOTHING;

-- Once Upon a Time in Hollywood (2019)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'Once Upon a Time in Hollywood' AND (a.first_name, a.last_name) IN (
    ('Leonardo', 'DiCaprio'),
    ('Brad',     'Pitt'),
    ('Margot',   'Robbie'),
    ('Al',       'Pacino'),
    ('Dakota',   'Fanning')
) ON CONFLICT DO NOTHING;

-- Django Unchained (2012)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'Django Unchained' AND (a.first_name, a.last_name) IN (
    ('Jamie',    'Foxx'),
    ('Leonardo', 'DiCaprio'),
    ('Samuel L.','Jackson'),
    ('Christoph','Waltz'),
    ('Kerry',    'Washington'),
    ('Bruce',    'Willis')
) ON CONFLICT DO NOTHING;

-- Inglourious Basterds (2009)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'Inglourious Basterds' AND (a.first_name, a.last_name) IN (
    ('Brad',      'Pitt'),
    ('Christoph', 'Waltz'),
    ('Michael',   'Fassbender'),
    ('Eli',       'Roth'),
    ('Diane',     'Kruger')
) ON CONFLICT DO NOTHING;

-- The Revenant (2015)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'The Revenant' AND (a.first_name, a.last_name) IN (
    ('Leonardo', 'DiCaprio'),
    ('Tom',      'Hardy'),
    ('Will',     'Poulter'),
    ('Domhnall', 'Gleeson')
) ON CONFLICT DO NOTHING;

-- Shutter Island (2010)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id FROM movies m CROSS JOIN actors a
WHERE m.title = 'Shutter Island' AND (a.first_name, a.last_name) IN (
    ('Leonardo', 'DiCaprio'),
    ('Mark',     'Ruffalo'),
    ('Ben',      'Kingsley'),
    ('Michelle', 'Williams'),
    ('Max',      'von Sydow')
) ON CONFLICT DO NOTHING;
