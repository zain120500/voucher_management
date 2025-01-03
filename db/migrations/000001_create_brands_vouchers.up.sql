CREATE TABLE brands (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE vouchers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    cost_in_point SMALLINT NOT NULL,
    brand_id INT NOT NULL,
    FOREIGN KEY (brand_id) REFERENCES brands(id)
);

CREATE TABLE redemptions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    customer_name VARCHAR(100) NOT NULL,
    total_points SMALLINT NOT NULL
);

CREATE TABLE redemption_vouchers (
    redemption_id INT NOT NULL,
    voucher_id INT NOT NULL,
    PRIMARY KEY (redemption_id, voucher_id),
    FOREIGN KEY (redemption_id) REFERENCES redemptions(id),
    FOREIGN KEY (voucher_id) REFERENCES vouchers(id)
);
