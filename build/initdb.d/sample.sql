CREATE DATABASE IF NOT EXISTS sushidb;
USE sushidb;

CREATE TABLE IF NOT EXISTS sushi
(
    id       INT PRIMARY KEY AUTO_INCREMENT, 
    category VARCHAR(255) NOT NULL,
    name     VARCHAR(255) NOT NULL,
    price    INT,
    calorie  INT
);

INSERT INTO sushi VALUES 
(0, "サーモンマウンテン", "期間限定/エリア限定", 330, 211), 
(0, "えび天マウンテン", "期間限定/エリア限定", 330, 383), 
(0, "天然インド鮪中落ちてんこ盛り", "期間限定/エリア限定", 165, 73), 
(0, "うなきゅう巻", "期間限定/エリア限定", 165, 199), 
(0, "大切り煮穴子", "期間限定/エリア限定", 165, 66);
