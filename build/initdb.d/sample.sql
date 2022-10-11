CREATE DATABASE IF NOT EXISTS sushidb;
USE sushidb;

CREATE TABLE IF NOT EXISTS sushi
(
    id       INT PRIMARY KEY AUTO_INCREMENT,
    category VARCHAR(255) NOT NULL,
    name     VARCHAR(255) NOT NULL,
    price    INT NOT NULL,
    calorie  INT NOT NULL
);

INSERT INTO sushi (category, name, price, calorie) VALUES
("期間限定/エリア限定", "サーモンマウンテン", 330, 211),
("期間限定/エリア限定", "えび天マウンテン", 330, 383),
("期間限定/エリア限定", "天然インド鮪中落ちてんこ盛り", 165, 73),
("期間限定/エリア限定", "うなきゅう巻", 165, 199),
("期間限定/エリア限定", "大切り煮穴子", 165, 66);
