CREATE TABLE IF NOT EXISTS cart_items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,        
    price NUMERIC NOT NULL,       
    quantity INT NOT NULL            
);