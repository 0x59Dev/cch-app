--
-- Roles
--

CREATE USER testdev
  WITH PASSWORD 'testdev';

--
-- Database creation
--

CREATE DATABASE testdb
  WITH OWNER testdev;

--
-- Access rights
--

REVOKE ALL ON DATABASE testdb FROM PUBLIC;
GRANT ALL ON DATABASE testdb TO testdev;

\connect testdb

REVOKE ALL ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO testdev;

--  Tables creation
--

CREATE TABLE domains(
  id serial primary key,
  domain_name varchar(100),
  expire_at timestamp without time zone,
  sign_algorithm varchar,
  cert_version integer,
  serial_number bigint,
  organization varchar
);


--
-- Test data insertion
--



GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO testdev;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO testdev;
