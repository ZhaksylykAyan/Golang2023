CREATE INDEX IF NOT EXISTS accessories_title_idx ON accessories USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS accessories_material_idx ON accessories USING GIN (to_tsvector('simple', material));