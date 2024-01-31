CREATE TABLE IF NOT EXISTS events (
    id int NOT NULL AUTO_INCREMENT UNIQUE PRIMARY KEY,
    eventName varchar(255) NOT NULL,
    eventDescription varchar(1024) NOT NULL,
    eventCategory int NOT NULL,
    eventDate DATETIME NOT NULL,
    eventLocation varchar(255) NOT NULL,
    eventImage varchar(1024) NOT NULL,
    eventBanner varchar(1024) NOT NULL,
    eventCapacity int NOT NULL,
    seatsRequired boolean NOT NULL DEFAULT 1,
    createdAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (eventCategory) REFERENCES categories(id) ON UPDATE CASCADE ON DELETE CASCADE
);