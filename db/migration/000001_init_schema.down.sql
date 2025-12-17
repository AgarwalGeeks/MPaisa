-- DOWN migration for the given UP migration

-- 1. Drop foreign key constraint first (because salary_split_items depends on salary_splits)
ALTER TABLE IF EXISTS "Finance"."Salary_split_items"
DROP CONSTRAINT IF EXISTS "salary_split_items_split_id_fkey";

-- 2. Drop the tables in reverse order of creation
DROP TABLE IF EXISTS "Finance"."Salary_split_items";
DROP TABLE IF EXISTS "Finance"."Salary_splits";
DROP TABLE IF EXISTS "Finance"."Credit_Card";

-- 3. Drop the sequences after tables that use them are dropped
DROP SEQUENCE IF EXISTS "Finance"."salary_split_items_id_seq";
DROP SEQUENCE IF EXISTS "Finance"."salary_splits_id_seq";