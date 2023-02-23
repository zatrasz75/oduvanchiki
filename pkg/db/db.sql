DROP TABLE IF EXISTS results, quizes, clientusers, correctanswers, answers, quiestions;


BEGIN;


CREATE TABLE IF NOT EXISTS public.answers
(
    id bigint NOT NULL DEFAULT nextval('answers_id_seq'::regclass),
    answer1 character varying(255) COLLATE pg_catalog."default" NOT NULL,
    answer2 character varying(255) COLLATE pg_catalog."default" NOT NULL,
    answer3 character varying(255) COLLATE pg_catalog."default" NOT NULL,
    answer4 character varying(255) COLLATE pg_catalog."default" NOT NULL,
    quiestionid bigint NOT NULL,
    CONSTRAINT answers_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.clientusers
(
    id bigint NOT NULL DEFAULT nextval('clientusers_id_seq'::regclass),
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    ip character varying(17) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT clientusers_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.correctanswers
(
    id bigint NOT NULL DEFAULT nextval('correctanswers_id_seq'::regclass),
    questionid bigint NOT NULL,
    answercorrect character varying(255) COLLATE pg_catalog."default" NOT NULL,
    correct boolean NOT NULL DEFAULT true,
    CONSTRAINT correctanswers_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quiestions
(
    id bigint NOT NULL DEFAULT nextval('quiestions_id_seq'::regclass),
    question character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT quiestions_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quizes
(
    id bigint NOT NULL DEFAULT nextval('quizes_id_seq'::regclass),
    userid bigint NOT NULL,
    started timestamp with time zone NOT NULL,
    CONSTRAINT quizes_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.results
(
    id bigint NOT NULL DEFAULT nextval('results_id_seq'::regclass),
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