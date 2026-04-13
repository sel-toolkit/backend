CREATE TABLE IF NOT EXISTS `teachers` (
    `id`         BIGINT       NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    `google_sub` VARCHAR(128) NOT NULL COMMENT 'Google unique user identifier',
    `email`      VARCHAR(255) NOT NULL COMMENT 'Google email',
    `name`       VARCHAR(255) NOT NULL COMMENT 'Display name',
    `avatar_url` VARCHAR(500)          DEFAULT NULL COMMENT 'Avatar URL',
    `created_at` DATETIME              DEFAULT NULL COMMENT 'Created at',
    `updated_at` DATETIME              DEFAULT NULL COMMENT 'Updated at',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_google_sub` (`google_sub`),
    UNIQUE KEY `uk_email` (`email`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='Teacher account data';
