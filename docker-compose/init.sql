-- base_stage.orders definition

CREATE TABLE `orders` (
  `id` varchar(255) NOT NULL,
  `price` decimal(10,2) NOT NULL,
  `tax` decimal(10,2) NOT NULL,
  `final_price` decimal(10,2) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

INSERT INTO base_stage.orders (id, price, tax, final_price) VALUES
('aaa1', 100.00, 10.00, 110.00),
('aaa2', 100.00, 10.00, 110.00),
('aaa3', 100.00, 10.00, 110.00),
('aaaa', 100.00, 10.00, 110.00),
('bbbbb', 50.00, 40.00, 90.00),
('ccc', 12.20, 5.00, 17.20),
('h5', 85.00, 85.00, 170.00),
('l5', 852.00, 70.00, 922.00),
('l7', 96.00, 74.00, 170.00),
('l8', 85.00, 85.00, 170.00),
('p7', 85.00, 85.00, 170.00),
('t7', 85.00, 85.00, 170.00),
('w4', 12.20, 5.00, 17.20),
('wwww', 12.20, 5.00, 17.20),
('wwww4', 12.20, 5.00, 17.20),
('wwww7', 12.20, 5.00, 17.20),
('wwww8', 12.20, 5.00, 17.20),
('zzz', 100.00, 10.00, 110.00);
