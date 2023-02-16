DROP TABLE IF EXISTS results, quizes, clientusers, correctanswers, answers, quiestions;


BEGIN;


CREATE TABLE IF NOT EXISTS public.answers
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    answer1 character varying(255) COLLATE pg_catalog."default" NOT NULL,
    answer2 character varying(255) COLLATE pg_catalog."default" NOT NULL,
    answer3 character varying(255) COLLATE pg_catalog."default" NOT NULL,
    answer4 character varying(255) COLLATE pg_catalog."default" NOT NULL,
    quiestionid bigint NOT NULL,
    CONSTRAINT answer_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.clientusers
(
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    CONSTRAINT client_users_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.correctanswers
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    questionid bigint NOT NULL,
    answercorrect character varying(255) COLLATE pg_catalog."default" NOT NULL,
    correct boolean NOT NULL DEFAULT true,
    CONSTRAINT correct_answers_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quiestions
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    question character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT quiestions_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quizes
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    userid bigint NOT NULL,
    started timestamp(6) with time zone NOT NULL,
    CONSTRAINT quizes_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.results
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    questionid bigint NOT NULL,
    answerid bigint NOT NULL,
    quizid bigint NOT NULL,
    answered timestamp(6) with time zone NOT NULL,
    point integer NOT NULL DEFAULT 0,
    CONSTRAINT results_pkey PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.answers
    ADD CONSTRAINT question_id FOREIGN KEY (quiestionid)
        REFERENCES public.quiestions (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID;


ALTER TABLE IF EXISTS public.correctanswers
    ADD CONSTRAINT quiestion_fkey FOREIGN KEY (questionid)
        REFERENCES public.quiestions (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID;


ALTER TABLE IF EXISTS public.quizes
    ADD CONSTRAINT uswers_id FOREIGN KEY (userid)
        REFERENCES public.clientusers (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID;


ALTER TABLE IF EXISTS public.results
    ADD CONSTRAINT answer_id FOREIGN KEY (answerid)
        REFERENCES public.answers (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID;


ALTER TABLE IF EXISTS public.results
    ADD CONSTRAINT question_id FOREIGN KEY (questionid)
        REFERENCES public.quiestions (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID;


ALTER TABLE IF EXISTS public.results
    ADD CONSTRAINT quiz_id FOREIGN KEY (quizid)
        REFERENCES public.quizes (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID;

END;