CREATE TABLE languages (
    id uuid primary key,
    email varchar not null,
    hours decimal not null,
    name varchar not null,
    date timestamp not null
);

CREATE INDEX languages_email_idx ON languages(email);
CREATE UNIQUE INDEX languages_date_unq ON languages (email, date, name);
