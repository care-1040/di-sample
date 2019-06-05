CREATE TABLE IF NOT EXISTS app.todos(
  `id` int(11) NOT NULL PRIMARY KEY,
  `body` text not null
) DEFAULT CHARSET=utf8;