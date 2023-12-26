-- drop schema go_test;
-- create schema go_test;
use go_test;
-- Tabela address_searchs
CREATE TABLE address_searchs (
    Id INT AUTO_INCREMENT PRIMARY KEY,
    District_Id INT,
    Address LONGTEXT,
    CEP LONGTEXT,
    Latitude LONGTEXT,
    Longitude LONGTEXT
);
-- Tabela BankAccounts
CREATE TABLE bank_Accounts (
    Id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Bank SMALLINT UNSIGNED NOT NULL,
    Agency LONGTEXT NOT NULL,
    AccountNumber LONGTEXT NOT NULL,
    IsSavingsAccount BOOL NOT NULL
);
-- Tabela cities
CREATE TABLE cities (
    Id INT AUTO_INCREMENT PRIMARY KEY,
    State_Id INT NOT NULL,
    Name LONGTEXT
);
-- Tabela districts
CREATE TABLE districts (
    Id INT AUTO_INCREMENT PRIMARY KEY,
    City_Id INT NOT NULL,
    Name LONGTEXT
);
-- Tabela states
CREATE TABLE states (
    Id INT AUTO_INCREMENT PRIMARY KEY,
    Name LONGTEXT
);
-- Tabela Users
CREATE TABLE Users (
    Id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    HashedPassword LONGBLOB NOT NULL,
    Name VARCHAR(30) NOT NULL,
    Email VARCHAR(30) NOT NULL UNIQUE,
    Phone VARCHAR(15) NOT NULL UNIQUE,
    CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    Filiated BOOL DEFAULT FALSE,
    PersonIdentifier LONGTEXT,
    Role TINYINT NOT NULL
);
-- Tabela Colaborators
CREATE TABLE Colaborators (
    Id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    IsAccepted BOOL NOT NULL,
    Disponible BOOL NOT NULL,
    EmergencyPhone VARCHAR(15) NOT NULL,
    UserId BIGINT UNSIGNED NOT NULL,
    BankAccountId BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (BankAccountId) REFERENCES bank_Accounts(Id) ON DELETE CASCADE,
    FOREIGN KEY (UserId) REFERENCES Users(Id) ON DELETE CASCADE
);
-- Tabela user_address
CREATE TABLE user_address (
    Id INT AUTO_INCREMENT PRIMARY KEY,
    UserId BIGINT UNSIGNED NOT NULL,
    AdressId INT NOT NULL,
    Number INT UNSIGNED NOT NULL,
    Reference LONGTEXT,
    NickName LONGTEXT,
    FOREIGN KEY (UserId) REFERENCES Users(Id),
    FOREIGN KEY (AdressId) REFERENCES address_searchs(Id)
);
-- Tabela user_fcm
CREATE TABLE user_fcm (
    Id INT AUTO_INCREMENT PRIMARY KEY,
    UserId BIGINT UNSIGNED NOT NULL,
    Token LONGTEXT NOT NULL,
    FOREIGN KEY (UserId) REFERENCES Users(Id)
);

-- Tabela restaurants
CREATE TABLE restaurants (
    Id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(30) NOT NULL,
    Partner BOOL DEFAULT FALSE,
    Segment smallint unsigned NOT NULL,
    Reviews MEDIUMINT UNSIGNED NOT NULL,
    Rating TINYINT UNSIGNED NOT NULL,
    IsOpen BOOL NOT NULL,
    CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    UserId BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (UserId) REFERENCES Users(Id)
);
-- Tabela restaurants_info	
CREATE TABLE restaurants_info (
    Id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    IsAccepted BOOL NOT NULL,
    CorporateIdentifier VARCHAR(20) NOT NULL,
    CorporateName VARCHAR(30) NOT NULL,
    BankAccountId BIGINT UNSIGNED NOT NULL,
    RestaurantId BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (BankAccountId) REFERENCES bank_Accounts(Id) ON DELETE CASCADE,
    FOREIGN KEY (RestaurantId) REFERENCES restaurants(Id)
);

-- Tabela product_categories
CREATE TABLE product_categories (
    Id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(20) NOT NULL UNIQUE,
    RestaurantId BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (RestaurantId) REFERENCES restaurants(Id)
);
-- Tabela products
CREATE TABLE products (
    Id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(27) NOT NULL,
    Description VARCHAR(150),
    CategoryId BIGINT UNSIGNED NOT NULL,
    RestaurantId BIGINT UNSIGNED NOT NULL,
    HasSubchoice BOOL,
    price DECIMAL(6,2) NOT NULL,
    ImageUrl VARCHAR(300) DEFAULT 'https://static.wikia.nocookie.net/whentheycry/images/e/e9/Ri_pc_sch.png',
    CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (CategoryId) REFERENCES product_categories(Id),
    FOREIGN KEY (RestaurantId) REFERENCES restaurants(Id) 
);
-- Tabela product_tags
CREATE TABLE product_tags (
    Id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Tag VARCHAR(255) NOT NULL,
    Maximum TINYINT UNSIGNED NOT NULL,
    IsRequired BOOL DEFAULT FALSE,
    ProductId BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (ProductId) REFERENCES products(Id)
);
-- Tabela ProductSubchoices
CREATE TABLE product_subchoices (
    Id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(40) NOT NULL,
    Price DECIMAL(5,2),
    ProductTagId BIGINT UNSIGNED,
    CONSTRAINT FK_ProductSubchoices_ProductTag FOREIGN KEY (ProductTagId) REFERENCES product_tags(Id)
);

-- Tabela orders
CREATE TABLE orders (
    Id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    total_price DECIMAL(7,2) NOT NULL,
    Status tinyint unsigned DEFAULT 0,
    Rating TINYINT UNSIGNED,
    ClientExperience VARCHAR(144),
    ClientDetails VARCHAR(144),
    UserId BIGINT UNSIGNED NOT NULL,
    RestaurantId BIGINT UNSIGNED NOT NULL,
    ColaboratorId BIGINT UNSIGNED NOT NULL,
    CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    PaidAt DATETIME,
    AcceptedAt DATETIME,
    PreparedAt DATETIME,
    DeliveredAt DATETIME,
    AddressId BIGINT UNSIGNED NOT NULL,
    UserAdressId INT,
    FOREIGN KEY (ColaboratorId) REFERENCES Colaborators(Id) ,
    FOREIGN KEY (UserId) REFERENCES Users(Id),
    FOREIGN KEY (RestaurantId) REFERENCES restaurants(Id),
    FOREIGN KEY (UserAdressId) REFERENCES user_address(Id)
);
-- Tabela order_products
CREATE TABLE order_products (
    Id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    OrderId BIGINT UNSIGNED NOT NULL,
    ProductId BIGINT UNSIGNED NOT NULL,
    Quantity TINYINT UNSIGNED NOT NULL,
    FOREIGN KEY (OrderId) REFERENCES orders(Id),
    FOREIGN KEY (ProductId) REFERENCES products(Id)
);
-- Tabela orderProductSubchoices
CREATE TABLE order_product_subchoices (
    Id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    OrderProductId BIGINT UNSIGNED,
    ProductSubchoiceId BIGINT UNSIGNED,
    Quantity TINYINT UNSIGNED,
    FOREIGN KEY (OrderProductId) REFERENCES order_products(Id),	
    FOREIGN KEY (ProductSubchoiceId) REFERENCES product_subchoices(Id)
);
-- Tabela orders_tip
CREATE TABLE orders_tip (
    Id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    OrderId BIGINT UNSIGNED NOT NULL,
    value DECIMAL(6,2) NOT NULL,
    Percentage TINYINT UNSIGNED NOT NULL,
    FOREIGN KEY (OrderId) REFERENCES orders(Id)
);

-- Tabela likes
CREATE TABLE likes (
    Id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    UserId BIGINT UNSIGNED NOT NULL,
    RestaurantId BIGINT UNSIGNED NOT NULL,
    IsLiked BOOL DEFAULT TRUE,
    FOREIGN KEY (UserId) REFERENCES Users(Id),
    FOREIGN KEY (RestaurantId) REFERENCES restaurants(Id)
);
