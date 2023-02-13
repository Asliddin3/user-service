CREATE Table If NOT exists users(
  id serial PRIMARY KEY,
  name VARCHAR(300),
  age INT,
  phone VARCHAR(30)
);