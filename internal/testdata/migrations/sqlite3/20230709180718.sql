-- Create "hotdog_type" table
CREATE TABLE `hotdog_type` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `name` varchar NOT NULL DEFAULT '',
  `description` text NOT NULL,
  `price` decimal NOT NULL DEFAULT 0
);
-- Create index "hotdog_type_name" to table: "hotdog_type"
CREATE UNIQUE INDEX `hotdog_type_name` ON `hotdog_type` (`name`);
-- Create index "hotdog_type_price" to table: "hotdog_type"
CREATE INDEX `hotdog_type_price` ON `hotdog_type` (`price`);
-- Create "stand" table
CREATE TABLE `stand` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `name` varchar NOT NULL DEFAULT '',
  `address` text NOT NULL,
  `description` text NOT NULL
);
-- Create index "stand_name" to table: "stand"
CREATE UNIQUE INDEX `stand_name` ON `stand` (`name`);
-- Create "hotdog_stock" table
CREATE TABLE `hotdog_stock` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `quantity` integer NOT NULL DEFAULT 0,
  `hotdog_id` integer NOT NULL,
  `stand_id` integer NOT NULL
);
-- Create index "hotdog_stock_hotdog_id" to table: "hotdog_stock"
CREATE INDEX `hotdog_stock_hotdog_id` ON `hotdog_stock` (`hotdog_id`);
-- Create index "hotdog_stock_stand_id" to table: "hotdog_stock"
CREATE INDEX `hotdog_stock_stand_id` ON `hotdog_stock` (`stand_id`);
