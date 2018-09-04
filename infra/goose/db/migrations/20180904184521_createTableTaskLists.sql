
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE task_lists (
	id integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
	task_id integer NOT NULL,
	title varchar(100) NOT NULL,
	description varchar(2000),
	creted_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	INDEX idx_id(id),
	FOREIGN KEY(task_id) REFERENCES tasks(id)
) ENGINE=InnoDB;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE task_lists;
