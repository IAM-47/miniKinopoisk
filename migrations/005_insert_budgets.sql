-- =====================================================
-- Добавление бюджетов и сборов
-- =====================================================

-- Создаём уникальный индекс, чтобы не было дублей
CREATE INDEX IF NOT EXISTS idx_budget_id_movie ON budget_and_fees (id_movie);

INSERT INTO budget_and_fees (id_movie, total_budget, fees_in_prod_country, fees_in_other)
SELECT m.id, b.total_budget, b.fees_in_prod_country, b.fees_in_other
FROM movies m
         JOIN (VALUES
                   ('Inception',                        160000000, 292576195, 543000000),
                   ('The Dark Knight',                  185000000, 534858444, 469700000),
                   ('Interstellar',                     165000000, 188020017, 487000000),
                   ('The Matrix',                        63000000, 171479930, 292037453),
                   ('Pulp Fiction',                       8000000, 107928762, 106100000),
                   ('Game of Thrones',                  100000000, 500000000, 900000000),
                   ('Fight Club',                        63000000,  37000000, 100000000),
                   ('Gone Girl',                         61000000,  83000000, 286000000),
                   ('The Social Network',                40000000,  96000000, 148000000),
                   ('Mr. & Mrs. Smith',                 110000000, 186000000, 266000000),
                   ('Ocean''s Eleven',                   85000000, 183000000, 267000000),
                   ('The Departed',                      90000000, 132000000, 157000000),
                   ('The Wolf of Wall Street',          100000000, 116000000, 276000000),
                   ('Joker',                             55000000, 335000000, 739000000),
                   ('Marriage Story',                    18000000,  10000000,  15000000),
                   ('Barbie',                           145000000, 636000000, 750000000),
                   ('Once Upon a Time in Hollywood',     90000000, 142000000, 232000000),
                   ('Django Unchained',                 100000000, 162000000, 263000000),
                   ('Inglourious Basterds',              70000000, 120000000, 221000000),
                   ('The Revenant',                     135000000, 183000000, 352000000),
                   ('Shutter Island',                    80000000, 128000000, 166000000)
) AS b(title, total_budget, fees_in_prod_country, fees_in_other)
              ON m.title = b.title
-- Если бюджет для фильма уже есть, обновляем его
ON CONFLICT (id_movie) DO UPDATE
    SET total_budget = EXCLUDED.total_budget,
        fees_in_prod_country = EXCLUDED.fees_in_prod_country,
        fees_in_other = EXCLUDED.fees_in_other;