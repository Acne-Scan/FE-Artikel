CREATE Table users(
    user_id int NOT NULL PRIMARY KEY,
    username VARCHAR(256),
    email VARCHAR(256),
    password VARCHAR(256),
    created_at DATE,
    updated_at DATE
);

CREATE Table pruduct_recommendation(
    recommendation_id int NOT NULL PRIMARY KEY,
    image VARCHAR(256),
    link VARCHAR(256),
    description VARCHAR(512),
    created_at DATE,
    updated_at DATE
);

CREATE Table article(
    article_id int NOT NULL PRIMARY KEY,
    image VARCHAR(256),
    description VARCHAR(512),
    created_at DATE,
    updated_at DATE
);


CREATE Table history(
    history_id int NOT NULL PRIMARY KEY,
    user_id int NOT NULL,
    image VARCHAR(256),
    result VARCHAR(256),
    created_at DATE,
    updated_at DATE,
    Foreign Key (user_id) REFERENCES users(user_id) on delete CASCADE
);

CREATE Table ingredients(
    ingredients_id int NOT NULL PRIMARY KEY,
    history_id int NOT NULL,
    ingredients VARCHAR(256),
    created_at DATE,
    updated_at DATE,
    Foreign Key (history_id) REFERENCES history(history_id) on delete CASCADE
);