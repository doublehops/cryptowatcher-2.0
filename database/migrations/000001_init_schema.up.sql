CREATE TABLE coins (
    id int(11) NOT NULL AUTO_INCREMENT,
    name varchar(200) NOT NUll,
    symbol varchar(10) NOT NULL,
    created_at int(32),
    updated_at int(32),
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4;

CREATE TABLE cmc_history (
    id int(11) NOT NULL AUTO_INCREMENT,
    name varchar(200) NOT NUll,
    symbol varchar(10) NOT NULL,
    slug varchar(100),
    num_market_pairs int(11),
    date_added varchar(100),
    max_supply int(11),
    circulating_supply int(11),
    total_supply int(11),
    cmc_rank int(11),
    quote_price float(12, 2),
    volume_24h float(12, 2),
    percent_change_1h float(12, 2),
    percent_change_24h float(12, 2),
    percent_change_7d float(12, 2),
    percent_change_30d float(12, 2),
    percent_change_60d float(12, 2),
    percent_change_90d float(12, 2),
    market_cap float(12, 2),
    created_at int(32),
    updated_at int(32),
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4;