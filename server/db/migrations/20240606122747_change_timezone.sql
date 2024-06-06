-- migrate:up
ALTER DATABASE social
SET TIMEZONE TO 'Asia/Bangkok';

-- migrate:down
ALTER DATABASE social
SET TIMEZONE TO 'UTC';
