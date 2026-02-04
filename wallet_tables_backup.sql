--
-- PostgreSQL database dump
--

\restrict 3edqMdYd9IOuU1PjKYOqT0fQGtj5Ube5b7DLZJyQ8YkdIaemjUV5y2teRJcF9zJ

-- Dumped from database version 18.1 (Debian 18.1-1.pgdg13+2)
-- Dumped by pg_dump version 18.1 (Debian 18.1-1.pgdg13+2)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: master_banks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.master_banks (
    id bigint NOT NULL,
    bank_code character varying(10) NOT NULL,
    bank_name character varying(100) NOT NULL,
    is_active boolean DEFAULT true,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.master_banks OWNER TO postgres;

--
-- Name: master_banks_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.master_banks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.master_banks_id_seq OWNER TO postgres;

--
-- Name: master_banks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.master_banks_id_seq OWNED BY public.master_banks.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(255) NOT NULL,
    telp_number character varying(20) NOT NULL,
    password character varying(255) NOT NULL,
    is_active boolean DEFAULT true,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: master_banks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_banks ALTER COLUMN id SET DEFAULT nextval('public.master_banks_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: master_banks; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.master_banks VALUES (1, '014', 'Bank BCA', true, NULL, NULL);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES (1, 'rema', 'rema@gmail.com', '0822109355', '$2a$10$NpS/nuHtRtvG1h9neJ72n.Y0ThFNb/B4zyWxXg5mLUJmp3JqRtPqK', true, '2026-02-03 09:04:49.586236+00', '2026-02-03 09:04:49.586236+00');


--
-- Name: master_banks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.master_banks_id_seq', 1, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- Name: master_banks master_banks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_banks
    ADD CONSTRAINT master_banks_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_master_banks_bank_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_master_banks_bank_code ON public.master_banks USING btree (bank_code);


--
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_users_email ON public.users USING btree (email);


--
-- Name: idx_users_telp_number; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_users_telp_number ON public.users USING btree (telp_number);


--
-- PostgreSQL database dump complete
--

\unrestrict 3edqMdYd9IOuU1PjKYOqT0fQGtj5Ube5b7DLZJyQ8YkdIaemjUV5y2teRJcF9zJ

