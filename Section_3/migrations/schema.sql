--
-- PostgreSQL database dump
--

\restrict yJ26Ba6QoF4BeZ73K6c320rRS2ZFj9CwDvkzHDrMBWpiDJuR3tc6oaYB2DjWpfT

-- Dumped from database version 14.20 (Ubuntu 14.20-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.20 (Ubuntu 14.20-0ubuntu0.22.04.1)

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
-- Name: Restriction; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Restriction" (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public."Restriction" OWNER TO postgres;

--
-- Name: Restriction_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Restriction_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Restriction_id_seq" OWNER TO postgres;

--
-- Name: Restriction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Restriction_id_seq" OWNED BY public."Restriction".id;


--
-- Name: banglow; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.banglow (
    id integer NOT NULL,
    banglow_name character varying(255) DEFAULT ''::character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.banglow OWNER TO postgres;

--
-- Name: banglowRestriction; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."banglowRestriction" (
    id integer NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL,
    banglow_id integer NOT NULL,
    reservation_id integer NOT NULL,
    created_at date NOT NULL,
    updated_at date NOT NULL,
    restriction_id integer NOT NULL
);


ALTER TABLE public."banglowRestriction" OWNER TO postgres;

--
-- Name: banglowRestriction_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."banglowRestriction_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."banglowRestriction_id_seq" OWNER TO postgres;

--
-- Name: banglowRestriction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."banglowRestriction_id_seq" OWNED BY public."banglowRestriction".id;


--
-- Name: banglow_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.banglow_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.banglow_id_seq OWNER TO postgres;

--
-- Name: banglow_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.banglow_id_seq OWNED BY public.banglow.id;


--
-- Name: reservation; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.reservation (
    id integer NOT NULL,
    full_name character varying(255) DEFAULT ''::character varying NOT NULL,
    email character varying(255) NOT NULL,
    phone character varying(255) NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL,
    banglow_id integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.reservation OWNER TO postgres;

--
-- Name: reservation_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.reservation_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.reservation_id_seq OWNER TO postgres;

--
-- Name: reservation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reservation_id_seq OWNED BY public.reservation.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    full_name character varying(255) DEFAULT ''::character varying NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(60) NOT NULL,
    role integer DEFAULT 1 NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: Restriction id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Restriction" ALTER COLUMN id SET DEFAULT nextval('public."Restriction_id_seq"'::regclass);


--
-- Name: banglow id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banglow ALTER COLUMN id SET DEFAULT nextval('public.banglow_id_seq'::regclass);


--
-- Name: banglowRestriction id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."banglowRestriction" ALTER COLUMN id SET DEFAULT nextval('public."banglowRestriction_id_seq"'::regclass);


--
-- Name: reservation id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservation ALTER COLUMN id SET DEFAULT nextval('public.reservation_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: Restriction Restriction_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Restriction"
    ADD CONSTRAINT "Restriction_pkey" PRIMARY KEY (id);


--
-- Name: banglowRestriction banglowRestriction_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."banglowRestriction"
    ADD CONSTRAINT "banglowRestriction_pkey" PRIMARY KEY (id);


--
-- Name: banglow banglow_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banglow
    ADD CONSTRAINT banglow_pkey PRIMARY KEY (id);


--
-- Name: reservation reservation_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservation
    ADD CONSTRAINT reservation_pkey PRIMARY KEY (id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: banglowRestriction_banglow_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "banglowRestriction_banglow_id_idx" ON public."banglowRestriction" USING btree (banglow_id);


--
-- Name: banglowRestriction_reservation_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "banglowRestriction_reservation_id_idx" ON public."banglowRestriction" USING btree (reservation_id);


--
-- Name: banglowRestriction_restriction_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "banglowRestriction_restriction_id_idx" ON public."banglowRestriction" USING btree (restriction_id);


--
-- Name: banglowRestriction_start_date_end_date_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "banglowRestriction_start_date_end_date_idx" ON public."banglowRestriction" USING btree (start_date, end_date);


--
-- Name: reservation_banglow_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX reservation_banglow_id_idx ON public.reservation USING btree (banglow_id);


--
-- Name: reservation_email_phone_full_name_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX reservation_email_phone_full_name_idx ON public.reservation USING btree (email, phone, full_name);


--
-- Name: reservation_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX reservation_id_idx ON public.reservation USING btree (id);


--
-- Name: reservation_start_date_end_date_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX reservation_start_date_end_date_idx ON public.reservation USING btree (start_date, end_date);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: users_email_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX users_email_idx ON public.users USING btree (email);


--
-- Name: banglowRestriction banglowRestriction_Restriction_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."banglowRestriction"
    ADD CONSTRAINT "banglowRestriction_Restriction_id_fk" FOREIGN KEY (restriction_id) REFERENCES public."Restriction"(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: banglowRestriction banglowRestriction_banglow_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."banglowRestriction"
    ADD CONSTRAINT "banglowRestriction_banglow_id_fk" FOREIGN KEY (banglow_id) REFERENCES public.banglow(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: banglowRestriction banglowRestriction_reservation_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."banglowRestriction"
    ADD CONSTRAINT "banglowRestriction_reservation_id_fk" FOREIGN KEY (reservation_id) REFERENCES public.reservation(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: reservation reservation_banglow_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservation
    ADD CONSTRAINT reservation_banglow_id_fk FOREIGN KEY (banglow_id) REFERENCES public.banglow(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

\unrestrict yJ26Ba6QoF4BeZ73K6c320rRS2ZFj9CwDvkzHDrMBWpiDJuR3tc6oaYB2DjWpfT

