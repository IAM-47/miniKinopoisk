-- =====================================================
-- Связи актёров с фильмами (много пересечений)
-- =====================================================

-- Inception (ID 1)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'Inception'
  AND a.first_name IN ('Leonardo', 'Joseph', 'Elliot', 'Tom', 'Ken', 'Brad', 'Matt', 'Jennifer', 'Ryan', 'Margot')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- The Dark Knight (ID 2)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'The Dark Knight'
  AND a.first_name IN ('Christian', 'Heath', 'Aaron', 'Gary', 'Maggie', 'Brad', 'George', 'Matt', 'Scarlett', 'Jennifer', 'Ryan')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- Interstellar (ID 3)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'Interstellar'
  AND a.first_name IN ('Matthew', 'Anne', 'Jessica', 'Michael', 'Mackenzie', 'Brad', 'Matt', 'Mark', 'Ryan', 'Leonardo')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- The Matrix (ID 4)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'The Matrix'
  AND a.first_name IN ('Keanu', 'Carrie-Anne', 'Laurence', 'Hugo', 'Joe', 'Brad', 'George', 'Scarlett', 'Leonardo')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- Pulp Fiction (ID 5)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'Pulp Fiction'
  AND a.first_name IN ('John', 'Uma', 'Samuel L.', 'Bruce', 'Harvey', 'Brad', 'Ryan', 'Margot', 'Jennifer', 'Leonardo')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- Game of Thrones (ID 6)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'Game of Thrones'
  AND a.first_name IN ('Brad', 'Matt', 'Ben', 'Jennifer', 'Mark', 'Jeremy', 'Tom', 'Ryan', 'Margot', 'Idris', 'Willem', 'George', 'Scarlett')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- Fight Club (ID 7)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'Fight Club'
  AND a.first_name IN ('Brad', 'George', 'Matt', 'Joseph', 'Tom', 'Christian', 'Ryan', 'Leonardo')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- Gone Girl (ID 8)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'Gone Girl'
  AND a.first_name IN ('Ben', 'Jennifer', 'Jessica', 'Matthew', 'Joseph', 'Ryan', 'Brad')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- The Social Network (ID 9)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'The Social Network'
  AND a.first_name IN ('George', 'Matt', 'Ben', 'Elliot', 'Leonardo', 'Jennifer', 'Ryan')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- Mr. & Mrs. Smith (ID 10)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'Mr. & Mrs. Smith'
  AND a.first_name IN ('Brad', 'Angelina', 'Jennifer', 'Ryan', 'Tom', 'George', 'Matt')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- Ocean's Eleven (ID 11)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'Ocean''s Eleven'
  AND a.first_name IN ('George', 'Brad', 'Matt', 'Scarlett', 'Leonardo', 'Christian', 'Ryan', 'Margot')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- The Departed (ID 12)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'The Departed'
  AND a.first_name IN ('Leonardo', 'Matt', 'Mark', 'Brad', 'Tom', 'Gary', 'Jennifer', 'Ryan')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- The Wolf of Wall Street (ID 13)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'The Wolf of Wall Street'
  AND a.first_name IN ('Leonardo', 'Joseph', 'Tom', 'Jennifer', 'Ryan', 'Margot', 'Brad', 'Matthew')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- Joker (ID 14)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'Joker'
  AND a.first_name IN ('Joaquin', 'Michael', 'Brad', 'Heath', 'Leonardo', 'Ryan', 'Robert')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- Marriage Story (ID 15)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'Marriage Story'
  AND a.first_name IN ('Scarlett', 'Jeremy', 'Jennifer', 'Emma', 'Jessica', 'Ryan', 'Brad', 'Adam')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- Barbie (ID 16)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'Barbie'
  AND a.first_name IN ('Margot', 'Ryan', 'Emma', 'Scarlett', 'Timothee', 'Leonardo', 'Brad', 'Jennifer')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- Once Upon a Time in Hollywood (ID 17)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'Once Upon a Time in Hollywood'
  AND a.first_name IN ('Leonardo', 'Brad', 'Margot', 'Joseph', 'Samuel L.', 'Ryan', 'Al')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- Django Unchained (ID 18)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'Django Unchained'
  AND a.first_name IN ('Samuel L.', 'Brad', 'Bruce', 'Leonardo', 'Tom', 'Matthew', 'Ryan', 'Jamie')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- Inglourious Basterds (ID 19)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'Inglourious Basterds'
  AND a.first_name IN ('Brad', 'Michael', 'Gary', 'Samuel L.', 'Ken', 'Aaron', 'Ryan', 'Christoph')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- The Revenant (ID 20)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'The Revenant'
  AND a.first_name IN ('Leonardo', 'Tom', 'Matt', 'Jennifer', 'Mark', 'Brad', 'Ryan', 'Will')
ON CONFLICT (id_movie, id_actor) DO NOTHING;

-- Shutter Island (ID 21)
INSERT INTO movie_actor (id_movie, id_actor)
SELECT m.id, a.id
FROM movies m
         CROSS JOIN actors a
WHERE m.title = 'Shutter Island'
  AND a.first_name IN ('Leonardo', 'Mark', 'Christian', 'Jessica', 'Joseph', 'Brad', 'Ryan', 'Michelle')
ON CONFLICT (id_movie, id_actor) DO NOTHING;