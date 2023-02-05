# -- Database: Dandelions

CREATE DATABASE "Dandelions"
WITH
OWNER = postgres
ENCODING = 'UTF8'
LC_COLLATE = 'Russian_Russia.1251'
LC_CTYPE = 'Russian_Russia.1251'
TABLESPACE = pg_default
CONNECTION LIMIT = -1
IS_TEMPLATE = False;

# -- Table: public.quiestions

CREATE TABLE IF NOT EXISTS public.quiestions
(
id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
question character varying(255) COLLATE pg_catalog."default" NOT NULL,
CONSTRAINT quiestions_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.quiestions
OWNER to postgres;

# -- Table: public.answers

CREATE TABLE IF NOT EXISTS public.answers
(
id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
answer1 character varying(255) COLLATE pg_catalog."default" NOT NULL,
answer2 character varying(255) COLLATE pg_catalog."default" NOT NULL,
answer3 character varying(255) COLLATE pg_catalog."default" NOT NULL,
answer4 character varying(255) COLLATE pg_catalog."default" NOT NULL,
quiestionid bigint NOT NULL,
CONSTRAINT answer_pkey PRIMARY KEY (id),
CONSTRAINT question_id FOREIGN KEY (quiestionid)
REFERENCES public.quiestions (id) MATCH SIMPLE
ON UPDATE NO ACTION
ON DELETE NO ACTION
NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.answers
OWNER to postgres;

# -- Table: public.correctanswers

CREATE TABLE IF NOT EXISTS public.correctanswers
(
id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
questionid bigint NOT NULL,
answercorrect character varying(255) COLLATE pg_catalog."default" NOT NULL,
correct boolean NOT NULL DEFAULT true,
CONSTRAINT correct_answers_pkey PRIMARY KEY (id),
CONSTRAINT quiestion_fkey FOREIGN KEY (questionid)
REFERENCES public.quiestions (id) MATCH SIMPLE
ON UPDATE NO ACTION
ON DELETE NO ACTION
NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.correctanswers
OWNER to postgres;

# -- Table: public.clientusers

CREATE TABLE IF NOT EXISTS public.clientusers
(
name character varying(255) COLLATE pg_catalog."default" NOT NULL,
id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
CONSTRAINT client_users_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.clientusers
OWNER to postgres;

# -- Table: public.quizes

CREATE TABLE IF NOT EXISTS public.quizes
(
id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
userid bigint NOT NULL,
started timestamp(6) with time zone NOT NULL,
CONSTRAINT quizes_pkey PRIMARY KEY (id),
CONSTRAINT uswers_id FOREIGN KEY (userid)
REFERENCES public.clientusers (id) MATCH SIMPLE
ON UPDATE NO ACTION
ON DELETE NO ACTION
NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.quizes
OWNER to postgres;

# -- Table: public.results

CREATE TABLE IF NOT EXISTS public.results
(
id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
questionid bigint NOT NULL,
answerid bigint NOT NULL,
quizid bigint NOT NULL,
answered timestamp(6) with time zone NOT NULL,
point integer NOT NULL DEFAULT 0,
CONSTRAINT results_pkey PRIMARY KEY (id),
CONSTRAINT answer_id FOREIGN KEY (answerid)
REFERENCES public.answers (id) MATCH SIMPLE
ON UPDATE NO ACTION
ON DELETE NO ACTION
NOT VALID,
CONSTRAINT question_id FOREIGN KEY (questionid)
REFERENCES public.quiestions (id) MATCH SIMPLE
ON UPDATE NO ACTION
ON DELETE NO ACTION
NOT VALID,
CONSTRAINT quiz_id FOREIGN KEY (quizid)
REFERENCES public.quizes (id) MATCH SIMPLE
ON UPDATE NO ACTION
ON DELETE NO ACTION
NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.results
OWNER to postgres;