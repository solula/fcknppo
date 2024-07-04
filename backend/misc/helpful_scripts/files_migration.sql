ALTER TABLE files
    ADD COLUMN IF NOT EXISTS type text,
    ADD COLUMN IF NOT EXISTS temp boolean,
    ADD COLUMN IF NOT EXISTS sequence_number integer
;

UPDATE files f
SET type = CASE
               WHEN f.mime_type LIKE 'image/jpeg' OR
                    f.mime_type LIKE 'image/png' AND (f.object_type = 'users' OR f.object_type = 'public') THEN 'avatar'
               WHEN f.mime_type LIKE 'image/jpeg' OR
                    f.mime_type LIKE 'image/png' AND f.object_type != 'users' THEN 'image'
               ELSE 'unknown'
    END
  , temp = false
WHERE type IS NULL
;

ALTER TABLE files
    ALTER COLUMN type SET NOT NULL,
ALTER
COLUMN temp SET NOT NULL
;