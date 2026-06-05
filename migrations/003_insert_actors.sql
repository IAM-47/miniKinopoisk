-- =====================================================
-- Добавление актёров (50 штук)
-- =====================================================

INSERT INTO actors (first_name, last_name, birth_date, salary)
SELECT * FROM (VALUES
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
                   ('Harvey',      'Keitel',         '1939-05-13',  5000000),
                   ('Brad',        'Pitt',           '1963-12-18', 20000000),
                   ('Angelina',    'Jolie',          '1975-06-04', 15000000),
                   ('George',      'Clooney',        '1961-05-06', 20000000),
                   ('Matt',        'Damon',          '1970-10-08', 20000000),
                   ('Ben',         'Affleck',        '1972-08-15', 15000000),
                   ('Jennifer',    'Lawrence',       '1990-08-15', 15000000),
                   ('Robert',      'Downey Jr.',     '1965-04-04', 25000000),
                   ('Chris',       'Evans',          '1981-06-13', 15000000),
                   ('Scarlett',    'Johansson',      '1984-11-22', 15000000),
                   ('Chris',       'Hemsworth',      '1983-08-11', 15000000),
                   ('Mark',        'Ruffalo',        '1967-11-22', 10000000),
                   ('Jeremy',      'Renner',         '1971-01-07', 10000000),
                   ('Tom',         'Holland',        '1996-06-01', 10000000),
                   ('Zendaya',     'Coleman',        '1996-09-01',  8000000),
                   ('Joaquin',     'Phoenix',        '1974-10-28', 10000000),
                   ('Lady Gaga',   'Germanotta',     '1986-03-28',  8000000),
                   ('Ryan',        'Gosling',        '1980-11-12', 15000000),
                   ('Emma',        'Stone',          '1988-11-06', 15000000),
                   ('Margot',      'Robbie',         '1990-07-02', 15000000),
                   ('Idris',       'Elba',           '1972-09-06', 10000000),
                   ('Cate',        'Blanchett',      '1969-05-04', 12000000),
                   ('Timothee',    'Chalamet',       '1995-12-27',  8000000),
                   ('Florence',    'Pugh',           '1996-01-03',  6000000),
                   ('Anya',        'Taylor-Joy',     '1996-04-16',  6000000),
                   ('Willem',      'Dafoe',          '1955-07-22',  7000000)
              ) AS new_actors(first_name, last_name, birth_date, salary)
WHERE NOT EXISTS (
    SELECT 1 FROM actors
    WHERE actors.first_name = new_actors.first_name
      AND actors.last_name = new_actors.last_name
);