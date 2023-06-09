DROP TABLE IF EXISTS results, quizes, clientusers, correctanswers, answers, quiestions, account_mails;


BEGIN;

CREATE TABLE IF NOT EXISTS public.account_mails
(
    id bigint NOT NULL,
    "from" character varying(55) COLLATE pg_catalog."default" NOT NULL,
    users character varying(55) COLLATE pg_catalog."default" NOT NULL,
    password character varying(55) COLLATE pg_catalog."default" NOT NULL,
    host character varying(55) COLLATE pg_catalog."default" NOT NULL,
    addr character varying(55) COLLATE pg_catalog."default" NOT NULL,
    "to" character varying(55) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT account_mails_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.answers
(
    id bigint NOT NULL,
    answer1 character varying(255) COLLATE pg_catalog."default" NOT NULL,
    answer2 character varying(255) COLLATE pg_catalog."default" NOT NULL,
    answer3 character varying(255) COLLATE pg_catalog."default" NOT NULL,
    answer4 character varying(255) COLLATE pg_catalog."default" NOT NULL,
    quiestionid bigint NOT NULL,
    CONSTRAINT answers_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.clientusers
(
    id bigint NOT NULL,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    ip character varying(17) COLLATE pg_catalog."default" NOT NULL,
    browser character varying(20) COLLATE pg_catalog."default" NOT NULL,
    platform character varying(20) COLLATE pg_catalog."default",
    CONSTRAINT clientusers_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.correctanswers
(
    id bigint NOT NULL,
    questionid bigint NOT NULL,
    answercorrect character varying(255) COLLATE pg_catalog."default" NOT NULL,
    correct boolean NOT NULL DEFAULT true,
    CONSTRAINT correctanswers_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quiestions
(
    id bigint NOT NULL,
    question character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT quiestions_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quizes
(
    id bigint NOT NULL,
    userid bigint NOT NULL,
    started timestamp with time zone NOT NULL,
    CONSTRAINT quizes_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.results
(
    id bigint NOT NULL,
    questionid bigint NOT NULL,
    answerid bigint NOT NULL,
    quizid bigint NOT NULL,
    answered timestamp with time zone NOT NULL,
    point bigint NOT NULL DEFAULT 0,
    CONSTRAINT results_pkey PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.answers
    ADD CONSTRAINT fk_answers_quiestions FOREIGN KEY (quiestionid)
        REFERENCES public.quiestions (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION;


ALTER TABLE IF EXISTS public.correctanswers
    ADD CONSTRAINT fk_correctanswers_quiestions FOREIGN KEY (questionid)
        REFERENCES public.quiestions (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION;


ALTER TABLE IF EXISTS public.quizes
    ADD CONSTRAINT fk_quizes_clientusers FOREIGN KEY (userid)
        REFERENCES public.clientusers (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION;


ALTER TABLE IF EXISTS public.results
    ADD CONSTRAINT fk_results_answers FOREIGN KEY (answerid)
        REFERENCES public.answers (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION;


ALTER TABLE IF EXISTS public.results
    ADD CONSTRAINT fk_results_quiestions FOREIGN KEY (questionid)
        REFERENCES public.quiestions (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION;


ALTER TABLE IF EXISTS public.results
    ADD CONSTRAINT fk_results_quizes FOREIGN KEY (quizid)
        REFERENCES public.quizes (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION;


END;

INSERT INTO public.account_mails (id, "from", users, password, host, addr, "to")
VALUES ('1', 'Vorobeyenglish@ya.ru', 'Vorobeyenglish@ya.ru', '4ff-k9S-47A-5yY', 'smtp.yandex.ru', 'smtp.yandex.ru:587', 'Vorobeyenglish@ya.ru');