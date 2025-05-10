package models

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var Database *sql.DB

func InitialiseDatabase() {
	db, err := sql.Open("sqlite3", "./database.db")

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	Database = db

	CreateDatabase()
}

func CreateDatabase() {
	Database.Exec("CREATE DATABASE IF NOT EXISTS database")

	Database.Exec("CREATE TABLE IF NOT EXISTS actions (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, method TEXT, location INTEGER, description TEXT)")
	Database.Exec("CREATE TABLE IF NOT EXISTS beliefs (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, value INTEGER, description TEXT)")
	Database.Exec("CREATE TABLE IF NOT EXISTS agents (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, description TEXT)")
	Database.Exec("CREATE TABLE IF NOT EXISTS goals (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, priority INTEGER, description TEXT)")
	Database.Exec("CREATE TABLE IF NOT EXISTS locations (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, description TEXT)")

	Database.Exec("CREATE TABLE IF NOT EXISTS agent_beliefs (agent_id INTEGER REFERENCES agents(id) ON DELETE CASCADE, belief_id INTEGER REFERENCES beliefs(id) ON DELETE CASCADE, PRIMARY KEY (agent_id, belief_id))")
	Database.Exec("CREATE TABLE IF NOT EXISTS agent_actions (agent_id INTEGER REFERENCES agents(id) ON DELETE CASCADE, action_id INTEGER REFERENCES actions(id) ON DELETE CASCADE, PRIMARY KEY (agent_id, action_id))")
	Database.Exec("CREATE TABLE IF NOT EXISTS agent_goals (agent_id INTEGER REFERENCES agents(id) ON DELETE CASCADE, goal_id INTEGER REFERENCES goals(id) ON DELETE CASCADE, PRIMARY KEY (agent_id, goal_id))")
	Database.Exec("CREATE TABLE IF NOT EXISTS agent_locations (agent_id INTEGER REFERENCES agents(id) ON DELETE CASCADE, location_id INTEGER REFERENCES locations(id) ON DELETE CASCADE, PRIMARY KEY (agent_id, location_id))")

	Database.Exec("CREATE TABLE IF NOT EXISTS action_prerequisites (action_id INTEGER REFERENCES actions(id) ON DELETE CASCADE, belief_id INTEGER REFERENCES beliefs(id) ON DELETE CASCADE, PRIMARY KEY (action_id, belief_id))")
	Database.Exec("CREATE TABLE IF NOT EXISTS action_outcomes (action_id INTEGER REFERENCES actions(id) ON DELETE CASCADE, belief_id INTEGER REFERENCES beliefs(id) ON DELETE CASCADE, PRIMARY KEY (action_id, belief_id))")

	Database.Exec("CREATE TABLE IF NOT EXISTS goal_beliefs (goal_id INTEGER REFERENCES goals(id) ON DELETE CASCADE, belief_id INTEGER REFERENCES beliefs(id) ON DELETE CASCADE, PRIMARY KEY (goal_id, belief_id))")
}
