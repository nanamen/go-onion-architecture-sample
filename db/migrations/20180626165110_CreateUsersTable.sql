-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `users` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL COMMENT 'ユーザ名',
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `users`;