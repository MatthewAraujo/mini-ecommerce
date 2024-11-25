CREATE TABLE IF NOT EXISTS stock (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL UNIQUE,
    available_quantity INT NOT NULL CHECK (available_quantity >= 0),
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);