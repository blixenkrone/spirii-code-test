-- CREATE TABLE IF NOT EXISTS categories (
--     id SERIAL PRIMARY KEY NOT NULL,
--     label TEXT NOT NULL,
--     description TEXT NOT NULL
-- );
-- CREATE TABLE IF NOT EXISTS courses (
--     id UUID PRIMARY KEY NOT NULL,
--     -- category_id SERIAL NOT NULL REFERENCES categories(id),
--     is_active BOOLEAN NOT NULL,
--     course_name TEXT NOT NULL,
--     created_at TIMESTAMP NOT NULL,
--     updated_at TIMESTAMP NOT NULL
-- );

-- CREATE TABLE IF NOT EXISTS modules (
--     id UUID PRIMARY KEY NOT NULL,
--     course_id UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
--     experience_level SMALLINT NOT NULL
-- );

-- CREATE TABLE IF NOT EXISTS material (
--     id UUID PRIMARY KEY NOT NULL,
--     module_id UUID NOT NULL REFERENCES modules(id) ON DELETE CASCADE,
--     description TEXT NOT NULL,
--     explanation TEXT NOT NULL,
--     object_url TEXT NOT NULL
-- );

CREATE TABLE IF NOT EXISTS foo (
    id UUID PRIMARY KEY NOT NULL,
    value TEXT NOT NULL
);

