
CREATE TABLE IF NOT EXISTS "hobbies" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "user_hobbies" (
    "user_id" UUID REFERENCES "users" ("id") NOT NULL,
    "hobby_id" UUID REFERENCES "hobbies" ("id") NOT NULL
);
