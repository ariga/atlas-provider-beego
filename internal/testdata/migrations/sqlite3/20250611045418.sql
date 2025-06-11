-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Drop "stand" table
DROP TABLE `stand`;
-- Create "hotdog_stand" table
CREATE TABLE `hotdog_stand` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `name` varchar NOT NULL DEFAULT '',
  `address` text NOT NULL,
  `description` text NOT NULL
);
-- Create index "hotdog_stand_name" to table: "hotdog_stand"
CREATE UNIQUE INDEX `hotdog_stand_name` ON `hotdog_stand` (`name`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
