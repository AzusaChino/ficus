CREATE DATABASE db_ficus ENCODING UTF8;

create table if not exists tb_tag (
    id serial primary key,
    name varchar(128) NOT NULL,
    created_at timestamp DEFAULT now(),
    created_by varchar(128) DEFAULT '',
    modified_at timestamp DEFAULT NULL,
    modified_by varchar(128) DEFAULT '',
    is_deleted boolean NOT NULL DEFAULT false,
    deleted_at timestamp DEFAULT NULL,
    status boolean NOT NULL DEFAULT true
);

create table if not exists tb_article (
    id serial primary key,
    title varchar(128) NOT NULL,
    brief varchar(255) NOT NULL,
    cover_image_url varchar(255) NOT NULL,
    content TEXT,
    created_at timestamp DEFAULT now(),
    created_by varchar(128) DEFAULT '',
    modified_at timestamp DEFAULT NULL,
    modified_by varchar(128) DEFAULT '',
    is_deleted boolean NOT NULL DEFAULT false,
    deleted_at timestamp DEFAULT NULL,
    status boolean NOT NULL DEFAULT true
);

create table if not exists tb_article_tag (
    id serial primary key,
    tag_id int not null,
    article_id int not null,
    created_at timestamp DEFAULT now(),
    created_by varchar(128) DEFAULT '',
    modified_at timestamp DEFAULT NULL,
    modified_by varchar(128) DEFAULT '',
    is_deleted boolean NOT NULL DEFAULT false,
    deleted_at timestamp DEFAULT NULL
);
