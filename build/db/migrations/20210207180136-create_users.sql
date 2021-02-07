
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    token VARCHAR(255)
);

-- +migrate Down
DROP TABLE IF EXISTS users;
