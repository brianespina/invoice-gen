-- CREATE TABLE client(
-- 	client_id 	INTEGER PRIMARY KEY,
-- 	name 		TEXT,
-- 	email 		TEXT,
-- 	rate 		REAL
-- );

-- CREATE TABLE timelog(
-- 	id 		INTEGER NOT NULL PRIMARY KEY,
-- 	name 		TEXT,
-- 	description	TEXT,
-- 	log 		REAL,
-- 	date 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
-- 	client		INTEGER NOT NULL,
-- 	FOREIGN KEY (client)
-- 		REFERENCES client (client_id)
-- );

-- INSERT INTO client (name, email, rate)
-- VALUES("Brian Espina", "espinabrian@gmail.com", 20.00);

-- INSERT INTO timelog(name, description, log, client)
-- VALUES ("MNY Site build", "new page created for MNY website", 16.0, 2)

-- ALTER TABLE timelog DROP COLUMN time;

-- SELECT * FROM  timelog;

-- DELETE FROM client
-- WHERE clien_id = 1;


