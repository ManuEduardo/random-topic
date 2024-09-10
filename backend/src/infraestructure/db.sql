CREATE TABLE users (
    id bigint primary key generated always as identity,
    name text not null,
    password text not null,
    birth_date date not null,
    gender text check (gender in ('male', 'female', 'other')) not null
);

CREATE TABLE card_types (
    id bigint primary key generated always as identity,
    name text not null unique,
    description text
);

CREATE TABLE cards (
    id bigint primary key generated always as identity,
    title text not null,
    content text not null,
    is_default boolean not null,
    user_id bigint references users(id),
    type_id bigint references card_types(id)
);

-- Insert default card types
INSERT INTO card_types (name, description) VALUES
('question', 'A card that poses a question to the user'),
('topic', 'A card that introduces a topic for discussion'),
('dynamic', 'A card that suggests an activity or dynamic to engage with others');