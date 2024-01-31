CREATE TABLE IF NOT EXISTS tickettypes (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    ticketName varchar(255) NOT NULL,
    typePrice decimal(6,2) NOT NULL,
    createdAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);