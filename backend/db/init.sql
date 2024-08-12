DROP TABLE IF EXISTS mbti_user;

CREATE TABLE IF NOT EXISTS mbti_user (
    id SERIAL PRIMARY KEY,
    last_name VARCHAR(100) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    nickname VARCHAR(100),
    mbti INTEGER NOT NULL
);

INSERT INTO mbti_user (last_name, first_name, nickname, mbti) VALUES
('田中', '太郎', 'たなっち', 1),
('山田', '花子', 'はなちゃん', 2);
