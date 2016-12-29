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

CREATE TABLE users(
  id serial primary key,
  username varchar(100) not null,
  userpass bytea not null,
  mail varchar(100) not null
);


--
-- Test data insertion
--



GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO testdev;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO testdev;
