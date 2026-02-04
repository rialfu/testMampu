--
-- PostgreSQL database dump
--

\restrict e4LAoAPE5FtM4BNLV8Xz722z4beFZCgmlgXRamyxNov2Vm7hXwV9fUzCOzUTYyd

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

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: deposits; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.deposits (
    id bigint NOT NULL,
    transaction_id bigint,
    source text,
    paid_at timestamp with time zone
);


ALTER TABLE public.deposits OWNER TO postgres;

--
-- Name: deposits_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.deposits_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.deposits_id_seq OWNER TO postgres;

--
-- Name: deposits_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.deposits_id_seq OWNED BY public.deposits.id;


--
-- Name: information_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.information_users (
    id bigint NOT NULL,
    nik text,
    is_verified boolean DEFAULT false,
    user_id bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.information_users OWNER TO postgres;

--
-- Name: information_users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.information_users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.information_users_id_seq OWNER TO postgres;

--
-- Name: information_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.information_users_id_seq OWNED BY public.information_users.id;


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
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id bigint NOT NULL,
    reference_no character varying(25) NOT NULL,
    wallet_id bigint,
    user_id bigint,
    transaction_type character varying(20),
    status smallint DEFAULT 3 NOT NULL,
    date_trans timestamp with time zone,
    created_at timestamp with time zone
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.transactions_id_seq OWNER TO postgres;

--
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


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
-- Name: wallet_ledgers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wallet_ledgers (
    id bigint NOT NULL,
    wallet_id bigint,
    transaction_id bigint,
    direction text,
    amount text,
    balance_before text,
    balance_after text,
    created_at timestamp with time zone
);


ALTER TABLE public.wallet_ledgers OWNER TO postgres;

--
-- Name: wallet_ledgers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wallet_ledgers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.wallet_ledgers_id_seq OWNER TO postgres;

--
-- Name: wallet_ledgers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wallet_ledgers_id_seq OWNED BY public.wallet_ledgers.id;


--
-- Name: wallets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wallets (
    id bigint NOT NULL,
    user_id bigint,
    balance numeric(16,2) DEFAULT '0'::numeric,
    updated_at timestamp with time zone
);


ALTER TABLE public.wallets OWNER TO postgres;

--
-- Name: wallets_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wallets_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.wallets_id_seq OWNER TO postgres;

--
-- Name: wallets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wallets_id_seq OWNED BY public.wallets.id;


--
-- Name: withdrawals; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.withdrawals (
    id bigint NOT NULL,
    transaction_id bigint,
    amount numeric(16,2) NOT NULL,
    fee numeric(16,2) NOT NULL,
    target_bank bigint,
    target_account character varying(50)
);


ALTER TABLE public.withdrawals OWNER TO postgres;

--
-- Name: withdrawals_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.withdrawals_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.withdrawals_id_seq OWNER TO postgres;

--
-- Name: withdrawals_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.withdrawals_id_seq OWNED BY public.withdrawals.id;


--
-- Name: deposits id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.deposits ALTER COLUMN id SET DEFAULT nextval('public.deposits_id_seq'::regclass);


--
-- Name: information_users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.information_users ALTER COLUMN id SET DEFAULT nextval('public.information_users_id_seq'::regclass);


--
-- Name: master_banks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_banks ALTER COLUMN id SET DEFAULT nextval('public.master_banks_id_seq'::regclass);


--
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: wallet_ledgers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallet_ledgers ALTER COLUMN id SET DEFAULT nextval('public.wallet_ledgers_id_seq'::regclass);


--
-- Name: wallets id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets ALTER COLUMN id SET DEFAULT nextval('public.wallets_id_seq'::regclass);


--
-- Name: withdrawals id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.withdrawals ALTER COLUMN id SET DEFAULT nextval('public.withdrawals_id_seq'::regclass);


--
-- Data for Name: deposits; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.deposits VALUES (1, 1, 'va', '0001-01-01 00:00:00+00');
INSERT INTO public.deposits VALUES (2, 5, 'va', '0001-01-01 00:00:00+00');


--
-- Data for Name: information_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.information_users VALUES (1, '', false, 1, '2026-02-03 09:04:49.609884+00', '2026-02-03 09:04:49.609884+00');


--
-- Data for Name: master_banks; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.master_banks VALUES (1, '014', 'Bank BCA', true, NULL, NULL);


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.transactions VALUES (1, '20260204647V0MJ0V20V56EED', 1, 1, 'deposit', 1, '2026-02-04 04:50:15.322376+00', '2026-02-04 04:50:15.323502+00');
INSERT INTO public.transactions VALUES (5, '20260204D3DH9382E99JQTX2S', 1, 1, 'deposit', 1, '2026-02-04 04:58:33.110141+00', '2026-02-04 04:58:33.111171+00');
INSERT INTO public.transactions VALUES (10, '20260204R0WJVD7T7WSYKILXV', 1, 1, 'withdraw', 1, '2026-02-04 05:01:44.184676+00', '2026-02-04 05:01:44.185364+00');


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES (1, 'rema', 'rema@gmail.com', '0822109355', '$2a$10$NpS/nuHtRtvG1h9neJ72n.Y0ThFNb/B4zyWxXg5mLUJmp3JqRtPqK', true, '2026-02-03 09:04:49.586236+00', '2026-02-03 09:04:49.586236+00');


--
-- Data for Name: wallet_ledgers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.wallet_ledgers VALUES (1, 1, 1, 'credit', '50000', '0', '50000', '2026-02-04 04:50:15.329968+00');
INSERT INTO public.wallet_ledgers VALUES (2, 1, 5, 'credit', '50000', '50000', '100000', '2026-02-04 04:58:33.112221+00');
INSERT INTO public.wallet_ledgers VALUES (3, 1, 10, 'debit', '10000', '100000', '90000', '2026-02-04 05:01:44.205933+00');


--
-- Data for Name: wallets; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.wallets VALUES (1, 1, 90000.00, '2026-02-04 05:01:44.20671+00');


--
-- Data for Name: withdrawals; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.withdrawals VALUES (1, 10, 10000.00, 0.00, 1, '1234');


--
-- Name: deposits_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.deposits_id_seq', 2, true);


--
-- Name: information_users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.information_users_id_seq', 1, true);


--
-- Name: master_banks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.master_banks_id_seq', 1, true);


--
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 10, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- Name: wallet_ledgers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wallet_ledgers_id_seq', 3, true);


--
-- Name: wallets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wallets_id_seq', 1, true);


--
-- Name: withdrawals_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.withdrawals_id_seq', 1, true);


--
-- Name: deposits deposits_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.deposits
    ADD CONSTRAINT deposits_pkey PRIMARY KEY (id);


--
-- Name: information_users information_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.information_users
    ADD CONSTRAINT information_users_pkey PRIMARY KEY (id);


--
-- Name: master_banks master_banks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_banks
    ADD CONSTRAINT master_banks_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: wallet_ledgers wallet_ledgers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallet_ledgers
    ADD CONSTRAINT wallet_ledgers_pkey PRIMARY KEY (id);


--
-- Name: wallets wallets_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets
    ADD CONSTRAINT wallets_pkey PRIMARY KEY (id);


--
-- Name: withdrawals withdrawals_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.withdrawals
    ADD CONSTRAINT withdrawals_pkey PRIMARY KEY (id);


--
-- Name: idx_information_users_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_information_users_user_id ON public.information_users USING btree (user_id);


--
-- Name: idx_master_banks_bank_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_master_banks_bank_code ON public.master_banks USING btree (bank_code);


--
-- Name: idx_transactions_reference_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_transactions_reference_no ON public.transactions USING btree (reference_no);


--
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_users_email ON public.users USING btree (email);


--
-- Name: idx_users_telp_number; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_users_telp_number ON public.users USING btree (telp_number);


--
-- Name: idx_wallets_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wallets_user_id ON public.wallets USING btree (user_id);


--
-- Name: deposits fk_deposits_transaction; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.deposits
    ADD CONSTRAINT fk_deposits_transaction FOREIGN KEY (transaction_id) REFERENCES public.transactions(id);


--
-- Name: transactions fk_transactions_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_transactions_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: information_users fk_users_information; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.information_users
    ADD CONSTRAINT fk_users_information FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: wallet_ledgers fk_wallet_ledgers_transaction; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallet_ledgers
    ADD CONSTRAINT fk_wallet_ledgers_transaction FOREIGN KEY (transaction_id) REFERENCES public.transactions(id);


--
-- Name: wallet_ledgers fk_wallet_ledgers_wallet; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallet_ledgers
    ADD CONSTRAINT fk_wallet_ledgers_wallet FOREIGN KEY (wallet_id) REFERENCES public.wallets(id);


--
-- Name: transactions fk_wallets_transaction; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_wallets_transaction FOREIGN KEY (wallet_id) REFERENCES public.wallets(id);


--
-- Name: wallets fk_wallets_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets
    ADD CONSTRAINT fk_wallets_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: withdrawals fk_withdrawals_bank; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.withdrawals
    ADD CONSTRAINT fk_withdrawals_bank FOREIGN KEY (target_bank) REFERENCES public.master_banks(id);


--
-- Name: withdrawals fk_withdrawals_transaction; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.withdrawals
    ADD CONSTRAINT fk_withdrawals_transaction FOREIGN KEY (transaction_id) REFERENCES public.transactions(id);


--
-- PostgreSQL database dump complete
--

\unrestrict e4LAoAPE5FtM4BNLV8Xz722z4beFZCgmlgXRamyxNov2Vm7hXwV9fUzCOzUTYyd

