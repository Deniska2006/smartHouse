CREATE TABLE IF NOT EXISTS public.events (
    id SERIAL PRIMARY KEY,
    device_id INTEGER NOT NULL,
    room_id INTEGER  NOT NULL,
    device_uuid UUID NOT NULL, 
    action TEXT NOT NULL,
    created_date TIMESTAMPTZ NOT NULL
);