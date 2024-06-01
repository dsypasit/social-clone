-- migrate:up
ALTER TABLE commend
  ADD COLUMN updated_at timestamp DEFAULT current_timestamp;

ALTER TABLE commend
  ADD COLUMN deleted_at DATE;


-- migrate:down
ALTER TABLE commend
  DROP COLUMN deleted_at;

ALTER TABLE commend
  DROP COLUMN updated_at;
