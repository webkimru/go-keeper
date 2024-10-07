-- +goose Up
CREATE TABLE IF NOT EXISTS keyvalues (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGSERIAL NOT NULL,
    title VARCHAR(100) NOT NULL,
    key VARCHAR(100) NOT NULL,
    value VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION updated_at()
RETURNS TRIGGER AS $$
BEGIN
	NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';
-- +goose StatementEnd

-- +goose StatementBegin
DO
$$BEGIN
    CREATE TRIGGER keyvalue_updated_at
        BEFORE UPDATE
        ON
            keyvalues
        FOR EACH ROW
    EXECUTE PROCEDURE updated_at();
EXCEPTION
	WHEN duplicate_object THEN
		NULL;
END;$$;
-- +goose StatementEnd

-- +goose Down
DROP TABLE keyvalues;