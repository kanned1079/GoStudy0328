## 这是一个测试项目

### 数据库表结构
- 导出 db0 的数据库结构

CREATE DATABASE IF NOT EXISTS `db0` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `db0`;

- 导出  表 db0.myusers 结构

CREATE TABLE IF NOT EXISTS `myusers` (
`user_id` int NOT NULL AUTO_INCREMENT,
`name` varchar(255) NOT NULL,
`gender` varchar(3) DEFAULT NULL,
`age` int DEFAULT NULL,
`phone_number` varchar(45) DEFAULT NULL,
`email` varchar(255) NOT NULL,
`password` varchar(255) NOT NULL,
`premium` char(10) DEFAULT NULL,
`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NULL DEFAULT NULL,
PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


