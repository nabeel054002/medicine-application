CREATE TABLE coupons (
    coupon_code VARCHAR PRIMARY KEY,
    expiry_date TIMESTAMP NOT NULL,  -- To mark the actual expiration date (optional)
    usage_type VARCHAR CHECK (usage_type IN ('one_time', 'multi_use', 'time_based')),
    min_order_value DECIMAL(10, 2),
    max_usage_per_user INT,
    start_window TIMESTAMP,  -- Start time for coupon validity
    end_window TIMESTAMP,    -- End time for coupon validity
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE coupon_discounts (
    coupon_code VARCHAR PRIMARY KEY,
    discount_type VARCHAR CHECK (discount_type IN ('items', 'charges')),
    discount_value DECIMAL(10, 2),
    FOREIGN KEY (coupon_code) REFERENCES coupons (coupon_code) ON DELETE CASCADE
);

CREATE TABLE coupon_applicable_medicines (
    coupon_code VARCHAR,
    medicine_id VARCHAR,
    PRIMARY KEY (coupon_code, medicine_id),
    FOREIGN KEY (coupon_code) REFERENCES coupons (coupon_code) ON DELETE CASCADE
);

CREATE TABLE coupon_applicable_categories (
    coupon_code VARCHAR,
    category VARCHAR,
    PRIMARY KEY (coupon_code, category),
    FOREIGN KEY (coupon_code) REFERENCES coupons (coupon_code) ON DELETE CASCADE
);

CREATE TABLE coupon_usage (
    coupon_code VARCHAR,
    user_id VARCHAR,
    usage_count INT DEFAULT 0,
    PRIMARY KEY (coupon_code, user_id),
    FOREIGN KEY (coupon_code) REFERENCES coupons (coupon_code) ON DELETE CASCADE
);

CREATE TABLE medicines (
    medicine_id VARCHAR PRIMARY KEY,  -- Unique identifier for each medicine
    name VARCHAR(255) NOT NULL,       -- Name of the medicine
    price DECIMAL(10, 2),             -- Price of the medicine
    stock INT,                        -- Quantity available in stock
    category_id VARCHAR,              -- Foreign key to the categories table
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- When the medicine was added
    FOREIGN KEY (category_id) REFERENCES categories (category_id) ON DELETE SET NULL  -- Ensures referential integrity
);

CREATE TABLE categories (
    category_id VARCHAR PRIMARY KEY,  -- Unique identifier for each category
    name VARCHAR(255) NOT NULL        -- Name of the category (e.g., "Painkiller", "Antibiotics")
);


-- a bit of a bad assumption that category and medicines are 1 to many => 1 category for many medicines but it can be many to many

