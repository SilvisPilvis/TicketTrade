CREATE TABLE IF NOT EXISTS roles (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    roleName varchar(255) NOT NULL,
    
    createReviews boolean NOT NULL DEFAULT FALSE,
    createUsers boolean NOT NULL DEFAULT FALSE,
    createRoles boolean NOT NULL DEFAULT FALSE,
    createEvents boolean NOT NULL DEFAULT FALSE,
    createTickets boolean NOT NULL DEFAULT FALSE,
    createCategories boolean NOT NULL DEFAULT FALSE,
    createTicketTypes boolean NOT NULL DEFAULT FALSE,

    deleteReviews boolean NOT NULL DEFAULT FALSE,
    deleteUsers boolean NOT NULL DEFAULT FALSE,
    deleteRoles boolean NOT NULL DEFAULT FALSE,
    deleteEvents boolean NOT NULL DEFAULT FALSE,
    deleteTickets boolean NOT NULL DEFAULT FALSE,
    deleteCategories boolean NOT NULL DEFAULT FALSE,
    deleteTicketTypes boolean NOT NULL DEFAULT FALSE,

    updateReviews boolean NOT NULL DEFAULT FALSE,
    updateUsers boolean NOT NULL DEFAULT FALSE,
    updateRoles boolean NOT NULL DEFAULT FALSE,
    updateEvents boolean NOT NULL DEFAULT FALSE,
    updateTickets boolean NOT NULL DEFAULT FALSE,
    updateCategories boolean NOT NULL DEFAULT FALSE,
    updateTicketTypes boolean NOT NULL DEFAULT FALSE,

    createdAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);