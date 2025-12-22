CREATE SCHEMA IF NOT EXISTS "Finance";

CREATE TABLE "Finance"."Credit_Card" (
    "id" SERIAL PRIMARY KEY,
    "bank_name" VARCHAR(20) NOT NULL,
    "card_name" VARCHAR(50) NOT NULL,
    "card_number" VARCHAR(16) NOT NULL,
    "cvv" INTEGER NOT NULL,
    "pin" INTEGER NOT NULL,
    "expiary_date" DATE NOT NULL,
    "usage" VARCHAR(100),
    "user_id" CHAR(16) NOT NULL
);

CREATE SEQUENCE "Finance"."salary_splits_id_seq"
    INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "Finance"."Salary_splits" (
    "id" INTEGER DEFAULT nextval('"Finance".salary_splits_id_seq') NOT NULL,
    "user_id" CHAR(16) NOT NULL,
    "month" DATE NOT NULL,
    "total_salary" NUMERIC(12,0) NOT NULL,
    "notes" TEXT,
    "is_fully_transferred" BOOLEAN DEFAULT FALSE,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "salary_splits_pkey" PRIMARY KEY ("id")
);

CREATE SEQUENCE "Finance"."salary_split_items_id_seq"
    INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "Finance"."Salary_split_items" (
    "id" INTEGER DEFAULT nextval('"Finance".salary_split_items_id_seq') NOT NULL,
    "split_id" INTEGER NOT NULL,
    "category_name" VARCHAR(100) NOT NULL,
    "amount" NUMERIC(12,0) NOT NULL,
    "move_to" VARCHAR(255),
    "is_transferred" BOOLEAN DEFAULT FALSE,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "salary_split_items_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "salary_split_items_split_id_fkey"
        FOREIGN KEY ("split_id")
        REFERENCES "Finance"."Salary_splits"("id")
        ON DELETE CASCADE
);