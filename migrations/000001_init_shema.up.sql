CREATE TABLE
    posts (
        id int AUTO_INCREMENT PRIMARY KEY,
        title varchar(200),
        content text,
        category varchar(100),
        created_at timestamp,
        updated_at timestamp,
        status varchar(100)
        -- Publish | Draft | Thrash
    );