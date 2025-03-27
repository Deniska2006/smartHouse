CREATE TABLE IF NOT EXISTS public.houses (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES public.users(id),
    name TEXT NOT NULL,
    description TEXT,
    city TEXT NOT NULL,
    address TEXT NOT NULL,
    lat DOUBLE PRECISION NOT NULL,
    lon DOUBLE PRECISION NOT NULL,
    created_date timestamptz NOT NULL,
    updated_date timestamptz NOT NULL,
    deleted_date timestamptz
);