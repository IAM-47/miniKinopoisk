-- =====================================================
-- Добавление актёров Game of Thrones
-- =====================================================

INSERT INTO actors (first_name, last_name, birth_date, salary)
SELECT * FROM (VALUES
    ('Emilia',      'Clarke',       '1986-10-23', 2000000),
    ('Kit',         'Harington',    '1986-12-26', 2000000),
    ('Peter',       'Dinklage',     '1969-06-11', 2500000),
    ('Lena',        'Headey',       '1973-10-03', 1800000),
    ('Nikolaj',     'Coster-Waldau','1970-07-27', 1500000),
    ('Sophie',      'Turner',       '1996-02-21', 1200000),
    ('Maisie',      'Williams',     '1997-04-15', 1200000),
    ('Jason',       'Momoa',        '1979-08-01', 1500000),
    ('Charles',     'Dance',        '1946-10-10', 1000000),
    ('Iain',        'Glen',         '1961-06-24', 900000)
) AS new_actors(first_name, last_name, birth_date, salary)
WHERE NOT EXISTS (
    SELECT 1 FROM actors
    WHERE actors.first_name = new_actors.first_name
      AND actors.last_name = new_actors.last_name
);
