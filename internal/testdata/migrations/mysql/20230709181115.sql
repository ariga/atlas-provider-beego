-- Create "hotdog_stock" table
CREATE TABLE `hotdog_stock` (
  `id` int NOT NULL AUTO_INCREMENT,
  `quantity` int NOT NULL DEFAULT 0,
  `hotdog_id` int NOT NULL,
  `stand_id` int NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `hotdog_stock_hotdog_id` (`hotdog_id`),
  INDEX `hotdog_stock_stand_id` (`stand_id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "hotdog_type" table
CREATE TABLE `hotdog_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT "",
  `description` longtext NOT NULL,
  `price` decimal(10,2) NOT NULL DEFAULT 0.00,
  PRIMARY KEY (`id`),
  INDEX `hotdog_type_price` (`price`),
  UNIQUE INDEX `name` (`name`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "stand" table
CREATE TABLE `stand` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT "",
  `address` longtext NOT NULL,
  `description` longtext NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name` (`name`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
