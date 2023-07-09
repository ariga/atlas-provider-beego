-- Create "hotdog_stock" table
CREATE TABLE "public"."hotdog_stock" (
  "id" serial NOT NULL,
  "quantity" integer NOT NULL DEFAULT 0,
  "hotdog_id" integer NOT NULL,
  "stand_id" integer NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "hotdog_stock_hotdog_id" to table: "hotdog_stock"
CREATE INDEX "hotdog_stock_hotdog_id" ON "public"."hotdog_stock" ("hotdog_id");
-- Create index "hotdog_stock_stand_id" to table: "hotdog_stock"
CREATE INDEX "hotdog_stock_stand_id" ON "public"."hotdog_stock" ("stand_id");
-- Create "hotdog_type" table
CREATE TABLE "public"."hotdog_type" (
  "id" serial NOT NULL,
  "name" text NOT NULL DEFAULT '',
  "description" text NOT NULL,
  "price" numeric(10,2) NOT NULL DEFAULT 0,
  PRIMARY KEY ("id")
);
-- Create index "hotdog_type_name_key" to table: "hotdog_type"
CREATE UNIQUE INDEX "hotdog_type_name_key" ON "public"."hotdog_type" ("name");
-- Create index "hotdog_type_price" to table: "hotdog_type"
CREATE INDEX "hotdog_type_price" ON "public"."hotdog_type" ("price");
-- Create "stand" table
CREATE TABLE "public"."stand" (
  "id" serial NOT NULL,
  "name" text NOT NULL DEFAULT '',
  "address" text NOT NULL,
  "description" text NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "stand_name_key" to table: "stand"
CREATE UNIQUE INDEX "stand_name_key" ON "public"."stand" ("name");
