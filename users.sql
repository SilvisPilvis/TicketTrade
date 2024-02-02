CREATE TABLE IF NOT EXISTS users (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username varchar(50) UNIQUE NOT NULL,
    password text NOT NULL,
    email varchar(255) NOT NULL,
    roleId int NOT NULL DEFAULT 1,
    lastLogin varchar(15),
    createdAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (roleId) REFERENCES roles(id) ON UPDATE CASCADE ON DELETE CASCADE
);