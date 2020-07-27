-- schema.sql
-- Since we might run the import many times we'll drop if exists
DROP DATABASE IF EXISTS books;

CREATE DATABASE books;

\c books;

CREATE TABLE books (
  id SERIAL PRIMARY KEY,
  title text not null,
  author text not null
);
