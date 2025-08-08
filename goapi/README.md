# 설치

```bash
mkdir goapi
cd goapi
go mod init github.com/anabasis/goapi
#go mod init

cat go.mod

go mod tidy
```

## Docker MySQL 설치

```docker
export MYSQL_DATABASE="mysql"
export MYSQL_CONTAINER_NAME="mysql"
export MYSQL_USER="admin"
export MYSQL_ROOT_PASSWORD="1"
export MYSQL_PASSWORD="1"

docker \
  run \
  --detach \
  --volume ~/Working/Containers/mysql:/var/lib/mysql \
  --env MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} \
  --env MYSQL_USER=${MYSQL_USER} \
  --env MYSQL_PASSWORD=${MYSQL_PASSWORD} \
  --env MYSQL_DATABASE=${MYSQL_DATABASE} \
  --name ${MYSQL_CONTAINER_NAME} \
  --hostname mysql.localdomain \
  --publish 3306:3306 \
  mysql:latest;
```

### 권한 및 테이블 생성

```sql
CREATE USER 'anabasis'@'%' IDENTIFIED BY 'Welcome!1';
CREATE DATABASE goapi_db;
#ALTER USER 'anabasis'@'%' IDENTIFIED BY 'Welcome!1';

GRANT ALL PRIVILEGES ON goapi_db.* TO 'anabasis'@'%';
FLUSH PRIVILEGES;

USE goapi_db;

CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    stock INT DEFAULT 0
);

INSERT INTO products (name, price, stock)
VALUES
('Laptop', 1200.50, 50),
('Mouse', 25.00, 200),
('Keyboard', 75.25, 120);

-- 모든 상품 정보 검색
SELECT * FROM products;

-- 가격이 50보다 높은 상품의 이름과 가격만 검색
SELECT name, price FROM products WHERE price > 50;

UPDATE products
SET stock = 250
WHERE id = 2;

DELETE FROM products
WHERE id = 3;

```
