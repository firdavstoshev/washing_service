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