-- up
CREATE TABLE publications (
	meta_keywords    VARCHAR (255) NOT NULL,
	meta_description VARCHAR (255) NOT NULL,
	title            VARCHAR (255) PRIMARY KEY,
	content          TEXT NOT NULL,
	created_at       DATETIME NOT NULL
);

-- down
DROP TABLE publications;
