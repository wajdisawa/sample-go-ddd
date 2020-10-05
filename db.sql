use ddd_sample;

CREATE TABLE IF NOT EXISTS users
(
    `id`         int          NOT NULL AUTO_INCREMENT,
    `name`       varchar(256) NOT NULL,
    `email`      varchar(256) not NULL,
    `created_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;