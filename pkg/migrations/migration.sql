CREATE TABLE IF NOT EXISTS unit_types -- Типы единиц измерения
(
    id   SERIAL PRIMARY KEY,  -- Уникальный идентификатор типа единицы измерения
    name VARCHAR(25) NOT NULL -- Название типа единицы измерения
);

CREATE TABLE IF NOT EXISTS service_types
(
    id        SERIAL PRIMARY KEY,                                  -- Уникальный идентификатор типа услуги
    name      VARCHAR(255) NOT NULL,                               -- Название типа услуги
    unit_type INTEGER REFERENCES unit_types (id) ON DELETE CASCADE -- Тип единицы измерения (UnitType)
);

CREATE TABLE IF NOT EXISTS services
(
    id         SERIAL PRIMARY KEY,                                      -- Уникальный идентификатор услуги
    name       VARCHAR(255)   NOT NULL,                                 -- Название услуги
    type       INTEGER REFERENCES service_types (id) ON DELETE CASCADE, -- Тип услуги (ServiceType)
    unit_price DECIMAL(10, 2) NOT NULL                                  -- Цена за единицу (UnitPrice)
);

CREATE TABLE IF NOT EXISTS customers
(
    id   SERIAL PRIMARY KEY,   -- Уникальный идентификатор клиента
    name VARCHAR(255) NOT NULL -- Имя клиента
);

CREATE TABLE IF NOT EXISTS orders
(
    id             SERIAL PRIMARY KEY,                                  -- Уникальный идентификатор заказа
    customer_id    INTEGER REFERENCES customers (id) ON DELETE CASCADE, -- Ссылка на клиента (FOREIGN KEY)
    is_child_items BOOLEAN        NOT NULL DEFAULT false,               -- Наличие детских вещей в заказе
    wait_days      INTEGER,                                             -- Количество дней, которые клиент готов подождать
    express        BOOLEAN                 DEFAULT FALSE,               -- Экспресс-услуга
    discount       DECIMAL(5, 2)           DEFAULT 0,                   -- Скидка на заказ в процентах
    total_price    DECIMAL(10, 2) NOT NULL                              -- Общая стоимость заказа
);

CREATE TABLE IF NOT EXISTS service_items
(
    id         SERIAL PRIMARY KEY,                                 -- Уникальный идентификатор записи
    order_id   INTEGER REFERENCES orders (id) ON DELETE CASCADE,   -- Ссылка на заказ (FOREIGN KEY)
    service_id INTEGER REFERENCES services (id) ON DELETE CASCADE, -- Ссылка на услугу (FOREIGN KEY)
    amount     DECIMAL(10, 2) NOT NULL                             -- Количество (Amount)
);

INSERT INTO unit_types (name)
VALUES ('шт'),
       ('кг');

SELECT *
FROM unit_types;

INSERT INTO service_types (name, unit_type)
VALUES ('Химчистка', 1),
       ('Ручная стирка', 2),
       ('Общие услуги по стирке', 2),
       ('Гладильные услуги', 1),
       ('Ремонт одежды', 1),
       ('Удаление пятен', 1);

SELECT *
FROM service_types;

INSERT INTO services (name, type, unit_price)
VALUES ('Пальто', 1, 20),
       ('Брюки', 1, 10),
       ('Костюм', 1, 15),
       ('Сюртук', 1, 15),
       ('Вечерное Платье', 1, 15),
       ('Свадебное Платье', 1, 15),
       ('Ручная стирка вещей', 2, 5),
       ('Белые', 3, 10),
       ('Цветные', 3, 8),
       ('Шерсть', 3, 12),
       ('Вещи из Шелка', 3, 15),
       ('Мягкие Игрушки', 3, 9),
       ('Постельное Бельё', 3, 10),
       ('Рубашки', 4, 5),
       ('Брюки', 4, 3),
       ('Юбки', 4, 2),
       ('Платья', 4, 6),
       ('Костюмы', 4, 8),
       ('Исправление шва или штопки', 5, 3.5),
       ('Пятна от масла', 6, 5),
       ('Пятна от крови', 6, 3),
       ('Общая грязь', 6, 2);

SELECT *
FROM services;

INSERT INTO customers (name)
VALUES ('Тошев Фирдавс Исроилович'),
       ('Петров Петр Петрович');

SELECT *
FROM customers;