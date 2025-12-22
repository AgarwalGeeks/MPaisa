-- Remove foreign key constraints from existing tables
ALTER TABLE "Finance"."Credit_Card"
DROP CONSTRAINT IF EXISTS "credit_card_user_id_fkey";

ALTER TABLE "Finance"."Salary_splits"
DROP CONSTRAINT IF EXISTS "salary_splits_user_id_fkey";

-- Drop the trigger and function for custom ID generation
DROP TRIGGER IF EXISTS set_custom_id ON "Finance"."users";
DROP FUNCTION IF EXISTS generate_custom_id;

-- Drop the users table
DROP TABLE IF EXISTS "Finance"."users";
