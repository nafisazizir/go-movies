CREATE TABLE movies (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  release_date DATE NOT NULL,
  poster_url TEXT NOT NULL,
  age_rating VARCHAR(10) NOT NULL,
  ticket_price INTEGER NOT NULL
);