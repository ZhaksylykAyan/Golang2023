ALTER TABLE accessories ADD CONSTRAINT accessories_runtime_check CHECK (runtime >= 0);
ALTER TABLE accessories ADD CONSTRAINT accessories_year_check CHECK (year BETWEEN 1888 AND date_part('year', now()));
ALTER TABLE accessories ADD CONSTRAINT accessories_price_check CHECK (price > 0);
