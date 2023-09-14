CREATE TABLE if not exists links (
  id SERIAL PRIMARY KEY,
  link text NOT NULL,
  short text NOT NULL
);

INSERT INTO links (link, short) VALUES ('http://google.com', 'CmgHeu');
INSERT INTO links (link, short) VALUES ('http://youtube.com', 'TutuBe');
