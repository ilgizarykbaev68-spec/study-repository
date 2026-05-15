DROP TABLE users;

// This migration is irreversible because it drops the users table and all its data.
//migrate -path migrations "postgres://postgres:salamandra@localhost:5432/postgres" up or down, чтобы переходить между миграциями. В данном случае, при выполнении down, будет удалена таблица users и все данные в ней. 