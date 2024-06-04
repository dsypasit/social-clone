-- migrate:up
ALTER TABLE comment
  ADD COLUMN updated_at timestamp DEFAULT current_timestamp;

ALTER TABLE comment
  ADD COLUMN deleted_at DATE;


-- migrate:down
ALTER TABLE comment
  DROP COLUMN deleted_at;

ALTER TABLE comment
  DROP COLUMN updated_at;
