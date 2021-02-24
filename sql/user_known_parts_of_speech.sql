CREATE TABLE user_known_parts_of_speech (
    id              bigserial PRIMARY KEY,
    "user"          bigint,
    part_of_speech  bigint,
    UNIQUE  ("user", part_of_speech),
    CONSTRAINT      fk_user
        FOREIGN KEY("user")
                REFERENCES "user"(id)
);