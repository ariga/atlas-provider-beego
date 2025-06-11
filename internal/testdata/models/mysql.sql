-- atlas:pos hotdog_stand[type=table] internal/testdata/models/hotdog.go:15:6
-- atlas:pos hotdog_stock[type=table] internal/testdata/models/hotdog.go:23:6
-- atlas:pos hotdog_type[type=table] internal/testdata/models/hotdog.go:7:6

-- --------------------------------------------------
--  Table Structure for `ariga.io/atlas-provider-beego/internal/testdata/models.HotdogType`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `hotdog_type` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `name` varchar(255) NOT NULL DEFAULT ''  UNIQUE,
    `description` longtext NOT NULL,
    `price` numeric(10, 2) NOT NULL DEFAULT 0 
) ENGINE=INNODB;
CREATE INDEX `hotdog_type_price` ON `hotdog_type` (`price`);
-- --------------------------------------------------
--  Table Structure for `ariga.io/atlas-provider-beego/internal/testdata/models.Stand`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `hotdog_stand` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `name` varchar(255) NOT NULL DEFAULT ''  UNIQUE,
    `address` longtext NOT NULL,
    `description` longtext NOT NULL
) ENGINE=INNODB;
-- --------------------------------------------------
--  Table Structure for `ariga.io/atlas-provider-beego/internal/testdata/models.HotdogStock`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `hotdog_stock` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `quantity` integer NOT NULL DEFAULT 0 ,
    `hotdog_id` integer NOT NULL,
    `stand_id` integer NOT NULL
) ENGINE=INNODB;
CREATE INDEX `hotdog_stock_hotdog_id` ON `hotdog_stock` (`hotdog_id`);
CREATE INDEX `hotdog_stock_stand_id` ON `hotdog_stock` (`stand_id`);

