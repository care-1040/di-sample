CREATE DATABASE IF NOT EXISTS app;

CREATE USER IF NOT EXISTS 'di'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON `app`.* TO 'di'@'%' WITH GRANT OPTION;
