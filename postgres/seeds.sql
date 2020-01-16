-- psql "postgres://postgres:$PGPW@172.16.2.201:5432/meetmeup_dev" -f ./postgres/seeds.sql

INSERT INTO users(username, email) VALUES ('bob', 'bob@gmail.com');
INSERT INTO users(username, email) VALUES ('jon', 'jon@gmail.com');
INSERT INTO users(username, email) VALUES ('jane', 'jane@gmail.com');

INSERT INTO meetups(name, description, user_id) VALUES('My first meetup','This is a description', 1);
INSERT INTO meetups(name, description, user_id) VALUES('My second meetup','This is another description', 1);
