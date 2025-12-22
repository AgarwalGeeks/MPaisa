CREATE TABLE "Finance"."users" (
    "id" CHAR(16) PRIMARY KEY,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "username" VARCHAR(255) NOT NULL UNIQUE,
    "password" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Trigger to update the updated_at column
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW."updated_at" = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON "Finance"."users"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Trigger to generate custom ID format
CREATE OR REPLACE FUNCTION generate_custom_id()
RETURNS TRIGGER AS $$
BEGIN
    NEW."id" := 'a01-' || lpad((FLOOR(RANDOM() * 100000000000000)::BIGINT)::TEXT, 12, '0');
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_custom_id
BEFORE INSERT ON "Finance"."users"
FOR EACH ROW
EXECUTE FUNCTION generate_custom_id();

ALTER TABLE "Finance"."Credit_Card"
ADD CONSTRAINT "credit_card_user_id_fkey"
FOREIGN KEY ("user_id")
REFERENCES "Finance"."users"("id")
ON DELETE CASCADE;

ALTER TABLE "Finance"."Salary_splits"
ADD CONSTRAINT "salary_splits_user_id_fkey"
FOREIGN KEY ("user_id")
REFERENCES "Finance"."users"("id")
ON DELETE CASCADE;
