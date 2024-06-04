-- migrate:up
CREATE TABLE IF NOT EXISTS app_user (
  id SERIAL PRIMARY KEY,
  uuid uuid UNIQUE, -- Consider adding UNIQUE constraint
  app_username varchar(25) NOT NULL,
  password varchar(250) NOT NULL,
  email varchar(250),
  gender varchar(10),
  profile_image varchar(250),
  is_deleted boolean DEFAULT false,
  updated_at timestamp DEFAULT current_timestamp,
  delete_at date
);

CREATE TABLE IF NOT EXISTS follows (
  id SERIAL PRIMARY KEY,
  follower_id int NOT NULL,
  followed_id int NOT NULL,
  created_at timestamp DEFAULT current_timestamp,

  FOREIGN KEY(follower_id) REFERENCES app_user(id),
  FOREIGN KEY(followed_id) REFERENCES app_user(id)
);

CREATE TABLE IF NOT EXISTS visibility_type (
  id SERIAL PRIMARY KEY,
  name varchar(250)
);

CREATE TABLE IF NOT EXISTS post (
  id SERIAL PRIMARY KEY,
  uuid uuid UNIQUE, -- Consider adding UNIQUE constraint
  content text,
  num_like int,
  visibility_type_id int,
  app_user_id int,
  deleted_at date,
  updated_at timestamp DEFAULT current_timestamp,

  FOREIGN KEY(app_user_id) REFERENCES app_user(id),
  FOREIGN KEY(visibility_type_id) REFERENCES visibility_type(id)
);

CREATE TABLE IF NOT EXISTS post_image (
  id SERIAL PRIMARY KEY,
  uuid uuid UNIQUE, -- Consider adding UNIQUE constraint
  img_url varchar(250),
  post_id int,

  FOREIGN KEY(post_id) REFERENCES post(id)
);

CREATE TABLE IF NOT EXISTS comment (
  id SERIAL PRIMARY KEY,
  uuid uuid UNIQUE, -- Consider adding UNIQUE constraint
  content text,
  app_user_id int,
  post_id int,

  FOREIGN KEY(app_user_id) REFERENCES app_user(id),
  FOREIGN KEY(post_id) REFERENCES post(id)
);

-- migrate:down
DROP TABLE IF EXISTS app_user;
DROP TABLE IF EXISTS post;
DROP TABLE IF EXISTS visibility_type;
DROP TABLE IF EXISTS comment;
DROP TABLE IF EXISTS follows;
DROP TABLE IF EXISTS post_image;
