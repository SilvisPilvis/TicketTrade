INSERT INTO roles (id, roleName) VALUES (1, 'buyer');
INSERT INTO `categories`(`categoryName`) VALUES ('Cinema');
INSERT INTO `categories` (`categoryName`) VALUES ('Nature');
INSERT INTO `categories` (`categoryName`) VALUES ('Theather');
INSERT INTO `categories` (`categoryName`) VALUES ('Expo');
INSERT INTO `categories` (`categoryName`) VALUES ('Automotive');
INSERT INTO tickettypes (ticketName, typePrice) VALUES ('Standart', 5.99);
INSERT INTO tickettypes (ticketName, typePrice) VALUES ('VIP', 10.99);
INSERT INTO tickettypes (ticketName, typePrice) VALUES ('Chungus Pass', 11.99);
INSERT INTO `events` (`id`, `eventName`, `eventDescription`, `eventCategory`, `eventDate`, `eventLocation`, `eventImage`, `eventBanner`, `eventCapacity`, `seatsRequired`, `createdAt`, `updatedAt`) VALUES (1, 'Ziedu mēness slepkavas', 'Filma.', 1, '2024-02-03', 'Rīgas iela 1, Kino \"Mogus\"', 'https://www.madona.lv/lat/box/pasakumi/6219a.jpg?v=0', 'https://media.cinemacloud.co.uk/imageFilm/1577_1_2.jpg', 26, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);