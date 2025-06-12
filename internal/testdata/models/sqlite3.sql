-- atlas:pos hotdog_stand[type=table] [ABS_PATH]/internal/testdata/models/hotdog.go:15:6
-- atlas:pos hotdog_stock[type=table] [ABS_PATH]/internal/testdata/models/hotdog.go:23:6
-- atlas:pos hotdog_type[type=table] [ABS_PATH]/internal/testdata/models/hotdog.go:7:6

-- --------------------------------------------------
--  Table Structure for `ariga.io/atlas-provider-beego/internal/testdata/models.HotdogType`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `hotdog_type` (
    `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    `name` varchar(255) NOT NULL DEFAULT ''  UNIQUE,
    `description` text NOT NULL,
    `price` decimal NOT NULL DEFAULT 0 
);
CREATE INDEX `hotdog_type_price` ON `hotdog_type` (`price`);
-- --------------------------------------------------
--  Table Structure for `ariga.io/atlas-provider-beego/internal/testdata/models.Stand`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `hotdog_stand` (
    `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    `name` varchar(255) NOT NULL DEFAULT ''  UNIQUE,
    `address` text NOT NULL,
    `description` text NOT NULL
);
-- --------------------------------------------------
--  Table Structure for `ariga.io/atlas-provider-beego/internal/testdata/models.HotdogStock`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `hotdog_stock` (
    `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    `quantity` integer NOT NULL DEFAULT 0 ,
    `hotdog_id` integer NOT NULL,
    `stand_id` integer NOT NULL
);
CREATE INDEX `hotdog_stock_hotdog_id` ON `hotdog_stock` (`hotdog_id`);
CREATE INDEX `hotdog_stock_stand_id` ON `hotdog_stock` (`stand_id`);

