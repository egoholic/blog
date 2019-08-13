-- up
CREATE TABLE rubrics (
	meta_keywords    VARCHAR(255) NOT NULL,
	meta_description VARCHAR(255) NOT NULL,
	title           VARCHAR(255) PRIMARY KEY,
	description     TEXT NOT NULL
);

-- down
DROP TABLE rubrics;