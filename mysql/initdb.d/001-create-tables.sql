---- drop ----
DROP TABLE IF EXISTS `users`;

---- create ----
create table IF not exists `users` (
 `id`               BIGINT AUTO_INCREMENT,
 `name`             VARCHAR(255) NOT NULL,
 `email`            VARCHAR(255) NOT NULL UNIQUE,
 `hashed_password`  VARCHAR(255) NOT NULL,
 `line_id`          VARCHAR(255) UNIQUE,
 `created_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into users(name, email, hashed_password) values("admin", "admin@test.com", "ca104e39a10ee19d78dcd1523b628c3028ba11343f355f81d4bea00065829161"); -- "password"
insert into users(name, email, hashed_password) values("test1", "test1@test.com", "a2c08eb261b5b9a3d59f016b910ccce10b85c4f7cb4a6dda9d4ffee3e14f3fe5"); -- "password"
insert into users(name, email, hashed_password) values("test2", "test2@test.com", "4925dda39aeac6606431b07042a2ca5f1b93924fe473a145972976b1a4ffe6f0"); -- "password"
insert into users(name, email, hashed_password) values("test3", "test3@test.com", "1d6cc6ca5412dd04c4a4dd09c52469ba99e4f27101407289b9ef68d2f0cc0a98"); -- "password"
insert into users(name, email, hashed_password) values("test4", "test4@test.com", "2be31d457ea77e14c14b0e29c1d1debbe4dfda41d999102ae58171d315c50a93"); -- "password"


---- drop ----
DROP TABLE IF EXISTS `rooms`;

---- create ----
create table IF not exists `rooms` (
 `id`               BIGINT AUTO_INCREMENT,
 `name`             VARCHAR(255) NOT NULL,
 `created_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into rooms(name) values("room1");
insert into rooms(name) values("room2");


---- drop ----
DROP TABLE IF EXISTS `user_rooms`;

---- create ----
create table IF not exists `user_rooms` (
 `id`               BIGINT AUTO_INCREMENT,
 `user_id`          BIGINT NOT NULL,
 `room_id`          BIGINT NOT NULL,
 `created_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into user_rooms(user_id, room_id) values(1, 1);
insert into user_rooms(user_id, room_id) values(2, 1);
insert into user_rooms(user_id, room_id) values(3, 1);
insert into user_rooms(user_id, room_id) values(4, 2);
insert into user_rooms(user_id, room_id) values(5, 2);


---- drop ----
DROP TABLE IF EXISTS `comments`;

---- create ----
create table IF not exists `comments` (
 `id`               BIGINT AUTO_INCREMENT,
 `text`             VARCHAR(255) NOT NULL,
 `user_id`          BIGINT NOT NULL,
 `created_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into comments(text, user_id) values("comment1", 1);
insert into comments(text, user_id) values("comment2", 1);
insert into comments(text, user_id) values("comment3", 2);
insert into comments(text, user_id) values("comment4", 2);


---- drop ----
DROP TABLE IF EXISTS `papers`;

---- create ----
create table IF not exists `papers` (
 `id`               BIGINT AUTO_INCREMENT,
 `url`              VARCHAR(255) NOT NULL,
 `text`             VARCHAR(255) NOT NULL,
 `user_id`          BIGINT NOT NULL,
 `is_deleted`       BOOLEAN DEFAULT 0,
 `created_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into papers(url, text, user_id) values("http://example.com/", "paper1", 1);
insert into papers(url, text, user_id) values("http://example.com/", "paper2", 1);
insert into papers(url, text, user_id) values("http://example.com/", "paper3", 2);
insert into papers(url, text, user_id) values("http://example.com/", "paper4", 2);
