-- 创建gin_demo数据库
CREATE DATABASE `gin_demo` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 创建gin_demo数据库的用户
CREATE USER 'gin_demo'@'%' IDENTIFIED WITH mysql_native_password BY 'gin_demo' PASSWORD EXPIRE NEVER;
-- 授权gin_demo数据库的用户
GRANT ALL PRIVILEGES ON gin_demo.* TO 'gin_demo'@'%';
-- 刷新权限
FLUSH PRIVILEGES;
