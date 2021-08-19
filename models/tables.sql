# 用户表
drop table if exists `user`;
create table `user` (
    `id` bigint(20) not null auto_increment,
    `user_id` bigint(20) collate utf8mb4_general_ci not null,
    `username` varchar(64) collate utf8mb4_general_ci not null,
    `password` varchar(64) collate utf8mb4_general_ci not null,
    `email` varchar(64) collate utf8mb4_general_ci,
    `gender` tinyint(1) not null default '0',
    `created_time` timestamp null default current_timestamp,
    `updated_time` timestamp null default current_timestamp on update current_timestamp,
    primary key (`id`),
    unique key `idx_username` (`username`) using btree,
    unique key `idx_user_id` (`user_id`) using btree
) engine=innodb default charset=utf8mb4 collate=utf8mb4_general_ci;

# 社区表
drop table if exists `category`;
create table `category` (
    `id` bigint(20) not null auto_increment,
    `category_id` bigint(20) collate utf8mb4_general_ci not null,
    `category_name` varchar(128) collate utf8mb4_general_ci not null,
    `introduction` varchar(256) collate utf8mb4_general_ci not null,
    `created_time` timestamp null default current_timestamp,
    `updated_time` timestamp null default current_timestamp on update current_timestamp,
    primary key (`id`),
    unique key `idx_category_id` (`category_id`) using btree,
    unique key `idx_category_name` (`category_name`) using btree
) engine=innodb default charset=utf8mb4 collate=utf8mb4_general_ci;

# 帖子表
drop table if exists `post`;
create table `post` (
    `id` bigint(20) not null auto_increment,
    `post_id` bigint(20) not null,
    `title` varchar(128) collate utf8mb4_general_ci not null comment '标题',
    `content` text collate utf8mb4_general_ci not null comment '内容',
    `author_id` bigint(20) not null comment '作者id',
    `category_id` bigint(20) not null comment '社区id',
    `status` tinyint(4) not null default 0 comment '帖子状态',
    `created_time` timestamp null default current_timestamp,
    `updated_time` timestamp null default current_timestamp on update current_timestamp,
    primary key (`id`),
    unique key `idx_post_id` (`post_id`),
    key `idx_author_id` (`author_id`),
    key `idx_category_id` (`category_id`)
) engine=innodb default charset=utf8mb4 collate=utf8mb4_general_ci;