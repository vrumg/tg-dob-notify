create table if not exists "group"
(
    group_id bigserial not null
        constraint group_pkey
            primary key,
    telegram_id bigint not null
        constraint group_pk
            unique
);

alter table "group" owner to postgres;

create table if not exists "user"
(
    user_id bigserial not null
        constraint user_pkey
            primary key,
    login varchar(50) not null
        constraint user_pk
            unique,
    birthdate date not null
);

alter table "user" owner to postgres;

create table if not exists user_group
(
    user_id bigint not null
        constraint user_group_user_user_id_fk
            references "user"
            on delete cascade,
    group_id bigint not null
        constraint user_group_group_group_id_fk
            references "group"
            on delete cascade,
    user_group_id bigserial not null
        constraint user_group_pkey
            primary key,
    constraint user_group_pk
        unique (user_id, group_id)
);

alter table user_group owner to postgres;