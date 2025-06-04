-- Modify "hotdog_type" table
ALTER TABLE "public"."hotdog_type" ADD CONSTRAINT "hotdog_type_name_key" UNIQUE USING INDEX "hotdog_type_name_key";
-- Modify "stand" table
ALTER TABLE "public"."stand" ADD CONSTRAINT "stand_name_key" UNIQUE USING INDEX "stand_name_key";
