package db

const (
	CreateUsersTable = `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(20) NOT NULL UNIQUE,
        password VARCHAR(60) NOT NULL
    );`

	CreateTasksTable = `CREATE  TABLE IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        title VARCHAR(30) NOT NULL UNIQUE,
        description VARCHAR(50) NOT NULL,
        completed BOOLEAN DEFAULT false,
        user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
        deadline INTEGER NOT NULL
    );`
	CreateGetExpiredTasksByUserFunc = `
    CREATE OR REPLACE FUNCTION get_expired_tasks_by_user(usr_id INT)
    RETURNS TABLE (task_id INT)
    LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY
        SELECT id FROM tasks WHERE user_id = usr_id AND
        created_at + deadline * interval '1 day' < NOW();
end;
$$;`
	CreateReassignTaskProcedure = `
    CREATE OR REPLACE PROCEDURE reassign_task(task_id INT, new_user_id INT)
    LANGUAGE plpgsql
    AS $$
    BEGIN
        UPDATE tasks
        SET user_id = new_user_id
        WHERE id = task_id;
    END;
    $$;`
)
