-- DDL
DROP DATABASE IF EXISTS `storage_desafio_db`;

CREATE DATABASE `storage_desafio_db`;

USE `storage_desafio_db`;

-- Table: customers
CREATE TABLE `customers` (
    `id` int NOT NULL AUTO_INCREMENT,
    `first_name` varchar(45) NULL,
    `last_name` varchar(45) NULL,
    `condition` boolean NULL,
    -- constraints
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Table: invoices
CREATE TABLE `invoices` (
    `id` int NOT NULL AUTO_INCREMENT,
    `datetime` datetime NULL,
    `total` float NULL,
    `customer_id` int NULL,
    -- constraints
    PRIMARY KEY (`id`),
    KEY `idx_invoices_customer_id` (`customer_id`),
    CONSTRAINT `fk_invoices_customer_id` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Table: products
CREATE TABLE `products` (
    `id` int NOT NULL AUTO_INCREMENT,
    `description` varchar(100) NULL,
    `price` float NULL,
    -- constraints
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Table: sales
CREATE TABLE `sales` (
    `id` int NOT NULL AUTO_INCREMENT,
    `quantity` int NULL,
    `invoice_id` int NULL,
    `product_id` int NULL,
    -- constraints
    PRIMARY KEY (`id`),
    KEY `idx_sales_invoice_id` (`invoice_id`),
    KEY `idx_sales_product_id` (`product_id`),
    CONSTRAINT `fk_sales_invoice_id` FOREIGN KEY (`invoice_id`) REFERENCES `invoices` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_sales_product_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;