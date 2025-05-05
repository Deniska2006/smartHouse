CREATE TABLE IF NOT EXISTS public.devices (
    id SERIAL PRIMARY KEY,
    house_id INTEGER REFERENCES public.houses(id),
    room_id INTEGER REFERENCES public.rooms(id),
    uuid UUID NOT NULL UNIQUE,
    serial_number TEXT NOT NULL,
    characteristics TEXT,
    category TEXT NOT NULL,
    units TEXT,
    power_consumption TEXT,
    created_date TIMESTAMPTZ NOT NULL,
    updated_date TIMESTAMPTZ NOT NULL,
    deleted_date TIMESTAMPTZ
);