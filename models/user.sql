CREATE TABLE `user`
(
    `id`          bigint(10) unsigned NOT NULL AUTO_INCREMENT primary key,
    `name`        varchar(255) COLLATE utf8mb4_general_ci NULL COMMENT 'The username',
    `password`    varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'The \n user password',
    `mobile`      varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'The mobile phone number',
    `gender`      char(10) COLLATE utf8mb4_general_ci      NOT NULL COMMENT 'gender,male|female|unknown',
    `nickname`    varchar(255) COLLATE utf8mb4_general_ci          DEFAULT '' COMMENT 'The nickname',
    `type`        tinyint(1) COLLATE utf8mb4_general_ci DEFAULT 0 COMMENT 'The user type, 0:normal,1:vip, for test golang keyword',
    `create_at` timestamp NULL,
    `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `name_index` (`name`),
    UNIQUE KEY `mobile_index` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'user table' COLLATE=utf8mb4_general_ci;

-- fn: FindOne
select * from user where id = ? limit 1;

-- fn: UpdateOne
update user set name = ? where id = ?;

-- fn: DeleteOne
delete from user where id = ? limit 1;

-- fn: FindLimit
select * from user where id > ? limit ?;

-- fn: Count
select count(id) AS count from user where id > ?;
