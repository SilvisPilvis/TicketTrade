CREATE TABLE IF NOT EXISTS tickets (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,

    eventId int NOT NULL,
    eventName varchar(255) NOT NULL,
    eventDescription text NOT NULL,
    ticketType int NOT NULL DEFAULT 1,
    ticketUID binary(16) UNIQUE NOT NULL DEFAULT (uuid_to_bin(uuid())),
    ticketSeat char(4),
    used boolean NOT NULL DEFAULT false,
    qrPath varchar(1024) DEFAULT NULL,

    ticketDate DATETIME NOT NULL,
    ticketLocation varchar(255) NOT NULL,

    createdAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (eventId) REFERENCES events(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (ticketType) REFERENCES ticketTypes(id) ON UPDATE CASCADE ON DELETE CASCADE
    -- FOREIGN KEY (eventName) REFERENCES events(eventName) ON UPDATE CASCADE ON DELETE CASCADE,
    -- FOREIGN KEY (ticketDate) REFERENCES events(eventDate) ON UPDATE CASCADE ON DELETE CASCADE,
    -- FOREIGN KEY (ticketTime) REFERENCES events(eventTime) ON UPDATE CASCADE ON DELETE CASCADE,
    -- FOREIGN KEY (ticketLocation) REFERENCES events(eventLocation) ON UPDATE CASCADE ON DELETE CASCADE
);