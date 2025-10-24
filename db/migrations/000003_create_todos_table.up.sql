CREATE TABLE todos(
  id SERIAL,
  title VARCHAR(200) NOT NULL,
  description TEXT,
  category_id int REFERENCES categories(id),
  priority priority,
  due_date TIMESTAMP,
  is_completed BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP   
);