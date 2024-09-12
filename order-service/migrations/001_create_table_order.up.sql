CREATE TABLE orders (
    id SERIAL PRIMARY KEY,  
    user_id VARCHAR NOT NULL, 
    location VARCHAR NOT NULL, 
    status INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_products (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE, 
    product_id VARCHAR NOT NULL 
);