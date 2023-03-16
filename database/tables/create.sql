CREATE TABLE IF NOT EXISTS users(
    id SERIAL NOT NULL PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    role TEXT NOT NULL,
    phone TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS products(
    product_id SERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    price decimal NOT NULL
);

CREATE TABLE IF NOT EXISTS orders(
    order_id SERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS product_order(
    product_id int NOT NULL,
    order_id int  NOT NULL,
    PRIMARY KEY (product_id, order_id),
    FOREIGN KEY (product_id) REFERENCES products ON DELETE CASCADE,
    FOREIGN KEY (order_id) REFERENCES orders ON DELETE CASCADE
);

INSERT INTO users (id,username, password, role, phone) VALUES (0, 'admin', 'password', 'Admin', '+380000000')
ON CONFLICT DO NOTHING;