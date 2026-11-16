--
-- Name: dog; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dog (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name text NOT NULL,
    age integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by integer NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_by integer NOT NULL,
    deleted_at timestamp with time zone,
    deleted_by integer
);


--
-- Name: training_type; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.training_type (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name text NOT NULL,
    description text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by integer NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_by integer NOT NULL,
    deleted_at timestamp with time zone,
    deleted_by integer
);


--
-- Name: training_session; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.training_session (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    dog_id integer NOT NULL REFERENCES public.dog(id),
    training_type_id integer NOT NULL REFERENCES public.training_type(id),
    started_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    duration_seconds integer NOT NULL,
    notes text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by integer NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_by integer NOT NULL,
    deleted_at timestamp with time zone,
    deleted_by integer
);




