-- =====================================================
-- Добавление фильмов (21 штука)
-- =====================================================

INSERT INTO movies (title, producer, director, release_year)
SELECT * FROM (VALUES
                   ('Inception',                        'Emma Thomas',      'Christopher Nolan',   2010),
                   ('The Dark Knight',                  'Emma Thomas',      'Christopher Nolan',   2008),
                   ('Interstellar',                     'Emma Thomas',      'Christopher Nolan',   2014),
                   ('The Matrix',                       'Joel Silver',      'Lana Wachowski',      1999),
                   ('Pulp Fiction',                     'Lawrence Bender',  'Quentin Tarantino',   1994),
                   ('Game of Thrones',                  'Brian Kirk',       'David Benioff',       2011),
                   ('Fight Club',                       'Art Linson',       'David Fincher',       1999),
                   ('Gone Girl',                        'Ceán Chaffin',     'David Fincher',       2014),
                   ('The Social Network',               'Scott Rudin',      'David Fincher',       2010),
                   ('Mr. & Mrs. Smith',                 'Lucas Foster',     'Doug Liman',          2005),
                   ('Ocean''s Eleven',                  'Jerry Weintraub',  'Steven Soderbergh',   2001),
                   ('The Departed',                     'Brad Pitt',        'Martin Scorsese',     2006),
                   ('The Wolf of Wall Street',          'Martin Scorsese',  'Martin Scorsese',     2013),
                   ('Joker',                            'Todd Phillips',    'Todd Phillips',       2019),
                   ('Marriage Story',                   'Noah Baumbach',    'Noah Baumbach',       2019),
                   ('Barbie',                           'Margot Robbie',    'Greta Gerwig',        2023),
                   ('Once Upon a Time in Hollywood',    'David Heyman',     'Quentin Tarantino',   2019),
                   ('Django Unchained',                 'Stacey Sher',      'Quentin Tarantino',   2012),
                   ('Inglourious Basterds',             'Lawrence Bender',  'Quentin Tarantino',   2009),
                   ('The Revenant',                     'Arnon Milchan',    'Alejandro G.',        2015),
                   ('Shutter Island',                   'Mike Medavoy',     'Martin Scorsese',     2010)
              ) AS new_movies(title, producer, director, release_year)
WHERE NOT EXISTS (
    SELECT 1 FROM movies WHERE movies.title = new_movies.title
);