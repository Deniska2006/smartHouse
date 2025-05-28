ALTER TABLE public.measurements
DROP COLUMN IF EXISTS device_uuid;

ALTER TABLE public.events
DROP COLUMN IF EXISTS device_uuid;