CREATE TABLE currency (
    id int(11) NOT NULL AUTO_INCREMENT,
    name varchar(200) NOT NUll,
    symbol varchar(10) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4;

CREATE TABLE cmc_history (
    id int(11) NOT NULL AUTO_INCREMENT,
    currency_id INT(11) NOT NULL,
    name varchar(200) NOT NUll,
    symbol varchar(10) NOT NULL,
    slug varchar(100),
    num_market_pairs int(11),
    date_added varchar(100),
    max_supply float(15, 2),
    circulating_supply float(15, 2),
    total_supply float(15, 2),
    cmc_rank int(11),
    quote_price float(15, 2),
    volume_24h float(15, 2),
    percent_change_1h float(15, 2),
    percent_change_24h float(15, 2),
    percent_change_7d float(15, 2),
    percent_change_30d float(15, 2),
    percent_change_60d float(15, 2),
    percent_change_90d float(15, 2),
    market_cap float(15, 2),
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4;