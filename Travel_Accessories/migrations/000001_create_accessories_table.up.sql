CREATE TABLE IF NOT EXISTS accessories
(
    id         bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title      text                        NOT NULL,
    year       integer                     NOT NULL,
    runtime    integer                     NOT NULL,
    material   text                        NOT NULL,
    version    integer                     NOT NULL DEFAULT 1,
    color      text                        NOT NULL,
    price      float                       NOT NULL
);