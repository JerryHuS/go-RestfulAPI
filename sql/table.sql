-- Table: public.user

-- DROP TABLE IF EXISTS public."user";

CREATE TABLE IF NOT EXISTS public."user"
(
    id          serial                                              NOT NULL,
    name        character varying(255) COLLATE pg_catalog."default" NOT NULL,
    dob         date                                                NOT NULL,
    address     jsonb,
    description text,
    following   integer[],
    followers   integer[],
    itime       timestamp with time zone                            NOT NULL,
    utime       timestamp with time zone                            NOT NULL,
    UNIQUE (name)
);