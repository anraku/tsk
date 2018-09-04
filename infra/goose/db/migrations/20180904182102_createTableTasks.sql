
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tasks (
	id integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
	title varchar(100) NOT NULL,
	description varchar(2000),
	deadline datetime,
	priority integer,
	status integer NOT NULL DEFAULT 0,
	execute_time datetime,
	creted_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	INDEX idx_id(id)
) ENGINE=InnoDB;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE tasks;
