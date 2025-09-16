-- +goose Up
INSERT INTO users (email,name,role,password_hash)
VALUES
 ('admin@example.com','Admin','ADMIN','$2a$10$kQyipbW7n6XyQd3w8Jc7euYpS1p5K0D2m0iGvYw2vY6r1sT8rSPyG'); -- bcrypt("admin123")

INSERT INTO exams (title,description,duration_minutes,price_paise)
VALUES
 ('Demo Math','Basics of arithmetic',60,9900),
 ('Reasoning Basics','Logical reasoning',45,5900);

-- +goose Down
DELETE FROM attempts;
DELETE FROM exams;
DELETE FROM users WHERE email='admin@example.com';
