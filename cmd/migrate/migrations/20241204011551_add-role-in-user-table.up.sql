CREATE TYPE user_role AS ENUM ('admin', 'user');

ALTER TABLE customers ADD COLUMN role user_role NOT NULL DEFAULT 'user';

UPDATE customers SET role = 'user' WHERE role IS NULL;

