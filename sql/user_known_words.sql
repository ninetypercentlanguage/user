CREATE TABLE user_known_word (
    id              integer PRIMARY KEY,
    "user_id"       integer,
    string          varchar(50) NOT NULL,
    part_of_speech  varchar(30) NOT NULL,
    CONSTRAINT      fk_user
        FOREIGN KEY("user_id")
                REFERENCES "user"(id)
);