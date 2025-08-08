CREATE TABLE IF NOT EXISTS item(
    id UUID PRIMARY KEY,  
    name VARCHAR NOT NULL,
    description VARCHAR,
    price NUMERIC(10, 2),
    taxcode VARCHAR -- digitax taxcode -
);

