package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func InitialiseDatabase() {
	db, error := sql.Open("sqlite3", "./database")

	if error != nil {
		panic(error)
	}

	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	database = db

	createDatabase()
}

func createDatabase() {
	database.Exec("CREATE DATABASE IF NOT EXISTS database")

	database.Exec("CREATE TABLE IF NOT EXISTS actions (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, function TEXT, location INTEGER)")
	database.Exec("CREATE TABLE IF NOT EXISTS beliefs (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, amount INT)")
	database.Exec("CREATE TABLE IF NOT EXISTS agents (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL)")
	database.Exec("CREATE TABLE IF NOT EXISTS goals (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL)")
	database.Exec("CREATE TABLE IF NOT EXISTS locations (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL)")

	database.Exec("CREATE TABLE IF NOT EXISTS agent_beliefs (agent_id INT REFERENCES agents(id) ON DELETE CASCADE, belief_id INT REFERENCES beliefs(id) ON DELETE CASCADE), PRIMARY KEY (agents_id, beliefs_id)")
	database.Exec("CREATE TABLE IF NOT EXISTS agent_actions (agent_id INT REFERENCES agents(id) ON DELETE CASCADE, action_id INT REFERENCES actions(id) ON DELETE CASCADE), PRIMARY KEY (agents_id, actions_id)")
	database.Exec("CREATE TABLE IF NOT EXISTS agent_goals (agent_id INT REFERENCES agents(id) ON DELETE CASCADE, goal_id INT REFERENCES goals(id) ON DELETE CASCADE), PRIMARY KEY (agents_id, goals_id)")
	database.Exec("CREATE TABLE IF NOT EXISTS agent_locations (agent_id INT REFERENCES agents(id) ON DELETE CASCADE, location_id INT REFERENCES locations(id) ON DELETE CASCADE), PRIMARY KEY (agents_id, locations_id)")

	database.Exec("CREATE TABLE IF NOT EXISTS action_prerequites (action_id INT REFERENCES actions(id) ON DELETE CASCADE, beliefs_id INT REFERENCES beliefs(id) ON DELETE CASCADE), PRIMARY KEY (actions_id, beliefs_id)")
	database.Exec("CREATE TABLE IF NOT EXISTS action_outcomes (action_id INT REFERENCES actions(id) ON DELETE CASCADE, beliefs_id INT REFERENCES beliefs(id) ON DELETE CASCADE), PRIMARY KEY (actions_id, beliefs_id)")

	database.Exec("CREATE TABLE IF NOT EXISTS goal_beliefs (goal_id INT REFERENCES goals(id) ON DELETE CASCADE, belief_id INT REFERENCES beliefs(id) ON DELETE CASCADE), PRIMARY KEY (quest_id, beliefs_id)")
}
