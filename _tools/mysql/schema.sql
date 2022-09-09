CREATE TABLE `user≈`

(
    `id`        BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーの識別子',
    `name`      varchar(20)  NOT NULL COMMENT 'ユーザー名',
    `password`  varchar(80)  NOT NULL COMMENT 'パスワードハッシュ',
    `role`      varchar(80)  NOT NULL COMMENT 'ロール',
    `created`   DATETIME(6) NOT NULL  COMMENT 'レコード作成日時',
    `modified`  DATETIME(6) NOT NULL  COMMENT 'レコード修正日時',
    PRIMARY KEY(`id`),
    UNIQUE KEY (`name`) USING BTREE
)Engin=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';