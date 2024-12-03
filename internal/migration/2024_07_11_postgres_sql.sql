--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1
-- Dumped by pg_dump version 16.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
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
-- Name: book_stocks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.book_stocks (
    book_id character varying(36) NOT NULL,
    code character varying(50) NOT NULL,
    status character varying(50) NOT NULL,
    borrower_id character varying(36),
    borrowed_at timestamp(6) without time zone,
    returned_at timestamp(6) without time zone
);


--
-- Name: books; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.books (
    id character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    title character varying(255) NOT NULL,
    description text,
    isbn character varying(100) NOT NULL,
    created_at timestamp(6) without time zone,
    updated_at timestamp(6) without time zone,
    deleted_at timestamp(6) without time zone
);


--
-- Name: customers; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.customers (
    id character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    code character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp(6) without time zone,
    updated_at timestamp(6) without time zone,
    deleted_at timestamp(6) without time zone
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL
);


--
-- Data for Name: book_stocks; Type: TABLE DATA; Schema: public; Owner: -
--



--
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.books (id, title, description, isbn, created_at, updated_at, deleted_at) VALUES ('a3a5bb3d-4d89-4adf-9bae-82701247415b', 'Thing and grow rich', 'Book about thingking and critial', '0000001', '2024-07-11 15:16:37.218247', '2024-07-11 15:17:26.976471', '2024-07-11 15:17:45.243556');


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.customers (id, code, name, created_at, updated_at, deleted_at) VALUES ('efee29c8-7a47-49a7-a053-a296d07e8d1a', 'M-0001', 'Senandika Selesa', '2024-07-10 21:54:06', '2024-07-10 21:54:09', NULL);
INSERT INTO public.customers (id, code, name, created_at, updated_at, deleted_at) VALUES ('99fee649-5ca3-46da-bba6-b90ff06bb4ad', 'MT-0004', 'Hendriawan Susanto', '2024-07-11 00:26:34.029608', '2024-07-11 00:38:49.537663', '2024-07-11 03:01:24.578502');


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.users (id, email, password) VALUES ('4246fb58-ff45-4d2a-8946-93e541fc39fd', 'admin@shellrean.id', '$2a$12$Rvslxj25D4OU7w3Ercz/IucMiDkEp1dOCSwq902oWpy0mqcUx2GAq');


--
-- Name: book_stocks book_stocks_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.book_stocks
    ADD CONSTRAINT book_stocks_pk PRIMARY KEY (code);


--
-- Name: books books_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pk PRIMARY KEY (id);


--
-- Name: customers customers_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pk PRIMARY KEY (id);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--
