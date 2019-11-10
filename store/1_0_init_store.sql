-- +migrate Up
CREATE TABLE jobs(
  id		 SERIAL PRIMARY KEY,
  title   CHARACTER VARYING(255) NOT NULL,
  description TEXT NOT NULL,
  city   CHARACTER VARYING(255) NOT NULL,
  email   CHARACTER VARYING(255) NOT NULL,
  company   CHARACTER VARYING(255) NOT NULL,
  phone   CHARACTER VARYING(255) NOT NULL,
  contact_name   CHARACTER VARYING(255) NOT NULL,
  currency_type  NUMERIC NOT NULL,
  employment_type  NUMERIC NOT NULL,
  salary NUMERIC DEFAULT NULL,
  is_remote   BOOLEAN DEFAULT FALSE,
  is_approved  BOOLEAN DEFAULT FALSE,
  created_at  TIMESTAMP NOT NULL
);

CREATE TABLE events(
  id		 SERIAL PRIMARY KEY,
  title   CHARACTER VARYING(255) NOT NULL,
  description TEXT NOT NULL,
  city   CHARACTER VARYING(255) NOT NULL,
  registration_link   CHARACTER VARYING(255) NOT NULL,
  contact_email   CHARACTER VARYING(255) NOT NULL,
  contact_phone   CHARACTER VARYING(255) NOT NULL,
  address   CHARACTER VARYING(255) NOT NULL,
  event_date   TIMESTAMP NOT NULL,
  is_approved  BOOLEAN DEFAULT FALSE,
  created_at  TIMESTAMP NOT NULL
);

-- +migrate Down
DROP TABLE jobs;
DROP TABLE events;