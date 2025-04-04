DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM users) THEN
        INSERT INTO users (Name, Email, Password) VALUES
            ('Иван Иванов', 'ivan@example.com', 'password1'),
            ('Петр Петров', 'petr@example.com', 'password2'),
            ('Светлана Смирнова', 'svetlana@example.com', 'password3'),
            ('Алексей Алексев', 'aleksey@example.com', 'password4'),
            ('Мария Маринина', 'maria@example.com', 'password5'),
            ('Анна Антонова', 'anna@example.com', 'password6'),
            ('Дмитрий Дмитриев', 'dmitry@example.com', 'password7'),
            ('Елена Еленина', 'elena@example.com', 'password8'),
            ('Сергей Сергеев', 'sergey@example.com', 'password9'),
            ('Ольга Ольгина', 'olga@example.com', 'password10');
END IF;
END
$$;
