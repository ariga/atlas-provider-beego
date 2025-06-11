-- Create "hotdog_stand" table
CREATE TABLE "public"."hotdog_stand" (
  "id" serial NOT NULL,
  "name" text NOT NULL DEFAULT '',
  "address" text NOT NULL,
  "description" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "hotdog_stand_name_key" UNIQUE ("name")
);
-- Drop "stand" table
DROP TABLE "public"."stand";
