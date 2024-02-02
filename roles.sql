CREATE TABLE IF NOT EXISTS roles (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    roleName varchar(255) NOT NULL,
    
    createReviews boolean NOT NULL DEFAULT FALSE,
    deleteReviews boolean NOT NULL DEFAULT FALSE,

    manageUsers boolean NOT NULL DEFAULT FALSE,
    manageRoles boolean NOT NULL DEFAULT FALSE,
    manageEvents boolean NOT NULL DEFAULT FALSE,
    manageTickets boolean NOT NULL DEFAULT FALSE,
    manageCategories boolean NOT NULL DEFAULT FALSE,
    manageTicketTypes boolean NOT NULL DEFAULT FALSE,

    createdAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);