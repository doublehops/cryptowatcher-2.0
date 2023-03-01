CREATE TABLE IF NOT EXISTS currency (
    id INT(11) NOT NULL AUTO_INCREMENT,
    name VARCHAR(200) NOT NUll,
    symbol VARCHAR(10) NOT NULL,
    image_url VARCHAR(200),
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    PRIMARY KEY (id)
    ) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4;
------------------
CREATE TABLE IF NOT EXISTS aggregator (
    id INT(11) NOT NULL AUTO_INCREMENT,
    name VARCHAR(200) NOT NUll,
    label VARCHAR(25) NOT NULL,
    is_enabled TINYINT(4),
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    PRIMARY KEY (id)
    ) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4;
------------------
INSERT INTO aggregator (name, label, is_enabled, created_at, updated_at) VALUES ('coinmarketcap', 'CoinMarketcap', 1, NOW(), NOW());
------------------
INSERT INTO aggregator (name, label, is_enabled, created_at, updated_at) VALUES ('coingecko', 'CoinGecko', 1, NOW(), NOW());
------------------
CREATE TABLE IF NOT EXISTS history (
    id INT(11) NOT NULL AUTO_INCREMENT,
    aggregator_id INT(11) NOT NULL,
    currency_id INT(11) NOT NULL,
    name VARCHAR(200) NOT NUll,
    symbol VARCHAR(10) NOT NULL,
    slug VARCHAR(100),
    all_time_high FLOAT(20,2),
    num_market_pairs INT(11),
    date_added VARCHAR(100),
    max_supply FLOAT(20, 2),
    circulating_supply FLOAT(20, 2),
    total_supply FLOAT(20, 2),
    quote_price FLOAT(15, 2),
    volume_24h FLOAT(15, 2),
    high_24hr FLOAT(15,2),
    low_24hr FLOAT(15,2),
    percent_change_1h FLOAT(6, 2),
    percent_change_24h FLOAT(6, 2),
    percent_change_7d FLOAT(6, 2),
    percent_change_30d FLOAT(6, 2),
    percent_change_60d FLOAT(6, 2),
    percent_change_90d FLOAT(6, 2),
    market_cap FLOAT(20, 2),
    rank INT(11) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    PRIMARY KEY (id),
    FOREIGN KEY (aggregator_id) REFERENCES aggregator(id)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4;