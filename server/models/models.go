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

	if !DatabseExists() {
		CreateDatabase()
		AddDefaultData()
	}
}

func DatabseExists() bool {
	result := Database.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='{actions}'")

	var name string
	result.Scan(&name)

	return name == "actions"
}

func CreateDatabase() {

	Database.Exec("CREATE TABLE IF NOT EXISTS actions (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, method TEXT, location INTEGER, description TEXT)")
	Database.Exec("CREATE TABLE IF NOT EXISTS beliefs (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, value INTEGER, description TEXT)")
	Database.Exec("CREATE TABLE IF NOT EXISTS agents (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, location INTEGER, description TEXT)")
	Database.Exec("CREATE TABLE IF NOT EXISTS goals (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, priority INTEGER, description TEXT)")
	Database.Exec("CREATE TABLE IF NOT EXISTS locations (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, description TEXT)")

	Database.Exec("CREATE TABLE IF NOT EXISTS agent_beliefs (agent_id INTEGER REFERENCES agents(id) ON DELETE CASCADE, belief_id INTEGER REFERENCES beliefs(id) ON DELETE CASCADE, PRIMARY KEY (agent_id, belief_id))")
	Database.Exec("CREATE TABLE IF NOT EXISTS agent_actions (agent_id INTEGER REFERENCES agents(id) ON DELETE CASCADE, action_id INTEGER REFERENCES actions(id) ON DELETE CASCADE, PRIMARY KEY (agent_id, action_id))")
	Database.Exec("CREATE TABLE IF NOT EXISTS agent_goals (agent_id INTEGER REFERENCES agents(id) ON DELETE CASCADE, goal_id INTEGER REFERENCES goals(id) ON DELETE CASCADE, PRIMARY KEY (agent_id, goal_id))")

	Database.Exec("CREATE TABLE IF NOT EXISTS action_prerequisites (action_id INTEGER REFERENCES actions(id) ON DELETE CASCADE, belief_id INTEGER REFERENCES beliefs(id) ON DELETE CASCADE, PRIMARY KEY (action_id, belief_id))")
	Database.Exec("CREATE TABLE IF NOT EXISTS action_outcomes (action_id INTEGER REFERENCES actions(id) ON DELETE CASCADE, belief_id INTEGER REFERENCES beliefs(id) ON DELETE CASCADE, PRIMARY KEY (action_id, belief_id))")

	Database.Exec("CREATE TABLE IF NOT EXISTS goal_beliefs (goal_id INTEGER REFERENCES goals(id) ON DELETE CASCADE, belief_id INTEGER REFERENCES beliefs(id) ON DELETE CASCADE, PRIMARY KEY (goal_id, belief_id))")
}

func AddDefaultData() {
	Database.Exec("INSERT INTO locations (name, description) VALUES ('Home', 'Home')")
	Database.Exec("INSERT INTO locations (name, description) VALUES ('Entrance', 'Entrance')")
	Database.Exec("INSERT INTO locations (name, description) VALUES ('Reception', 'Reception')")
	Database.Exec("INSERT INTO locations (name, description) VALUES ('WaitingRoom', 'Waiting Room')")
	Database.Exec("INSERT INTO locations (name, description) VALUES ('Cubicle', 'Cubicle')")
	Database.Exec("INSERT INTO locations (name, description) VALUES ('BreakRoom', 'Break Room')")

	Database.Exec("INSERT INTO agents (name, location, description) VALUES ('Nurse1', 6, 'Nurse 1')")
	Database.Exec("INSERT INTO agents (name, location, description) VALUES ('Nurse2', 6, 'Nurse 2')")

	Database.Exec("INSERT INTO agents (name, location, description) VALUES ('Patient1', 1, 'Patient 1')")
	Database.Exec("INSERT INTO agents (name, location, description) VALUES ('Patient2', 1, 'Patient 2')")
	Database.Exec("INSERT INTO agents (name, location, description) VALUES ('Patient3', 1, 'Patient 3')")
	Database.Exec("INSERT INTO agents (name, location, description) VALUES ('Patient4', 1, 'Patient 4')")
	Database.Exec("INSERT INTO agents (name, location, description) VALUES ('Patient5', 1, 'Patient 5')")

	Database.Exec("INSERT INTO beliefs (name, value, description) VALUES ('AtHome', 1, 'At home')")
	Database.Exec("INSERT INTO beliefs (name, value, description) VALUES ('AtEntrance', 1, 'At entrance')")
	Database.Exec("INSERT INTO beliefs (name, value, description) VALUES ('AtReception', 1, 'At reception')")
	Database.Exec("INSERT INTO beliefs (name, value, description) VALUES ('AtWaitingRoom', 1, 'At waiting room')")
	Database.Exec("INSERT INTO beliefs (name, value, description) VALUES ('AtCubicle', 1, 'At cubicle')")
	Database.Exec("INSERT INTO beliefs (name, value, description) VALUES ('AtBreakRoom', 1, 'At break room')")

	Database.Exec("INSERT INTO actions (name, method, location, description) VALUES ('GoToHospital', '', 2, 'Go to the hospital')")
	Database.Exec("INSERT INTO actions (name, method, location, description) VALUES ('GoToReception', 'Register', 3, 'Go to reception')")
	Database.Exec("INSERT INTO actions (name, method, location, description) VALUES ('GoToWaitingRoom', '', 4, 'Go to waiting room')")
	Database.Exec("INSERT INTO actions (name, method, location, description) VALUES ('GoToCubicle', '', 5, 'Go to cubicle')")
	Database.Exec("INSERT INTO actions (name, method, location, description) VALUES ('GoToHome', '', 1, 'Go home')")
	Database.Exec("INSERT INTO actions (name, method, location, description) VALUES ('GoToBreakRoom', 'Rest', 6, 'Go to break room')")
	Database.Exec("INSERT INTO actions (name, method, location, description) VALUES ('TreatPatient', 'Heal', 6, 'Treat a patient')")

	Database.Exec("INSERT INTO goals (name, priority, description) VALUES ('AwaitingTreatment', 1, 'At the hospital waiting room')")
	Database.Exec("INSERT INTO goals (name, priority, description) VALUES ('BeingTreated', 1, 'Being treated at the hospital')")
	Database.Exec("INSERT INTO goals (name, priority, description) VALUES ('GoneHome', 1, 'Back home')")
	Database.Exec("INSERT INTO goals (name, priority, description) VALUES ('TreatPatient', 1, 'Treat a sick patient')")
	Database.Exec("INSERT INTO goals (name, priority, description) VALUES ('TakeABreak', 1, 'Take a break in the break room')")

	Database.Exec("INSERT INTO action_prerequisites (action_id, belief_id) VALUES (0, 0)")
	Database.Exec("INSERT INTO action_prerequisites (action_id, belief_id) VALUES (1, 1)")
	Database.Exec("INSERT INTO action_prerequisites (action_id, belief_id) VALUES (2, 2)")
	Database.Exec("INSERT INTO action_prerequisites (action_id, belief_id) VALUES (3, 3)")
	Database.Exec("INSERT INTO action_prerequisites (action_id, belief_id) VALUES (4, 4)")
	Database.Exec("INSERT INTO action_prerequisites (action_id, belief_id) VALUES (5, 4)")

	Database.Exec("INSERT INTO action_outcomes (action_id, belief_id) VALUES (0, 1)")
	Database.Exec("INSERT INTO action_outcomes (action_id, belief_id) VALUES (1, 2)")
	Database.Exec("INSERT INTO action_outcomes (action_id, belief_id) VALUES (2, 3)")
	Database.Exec("INSERT INTO action_outcomes (action_id, belief_id) VALUES (3, 4)")
	Database.Exec("INSERT INTO action_outcomes (action_id, belief_id) VALUES (4, 0)")
	Database.Exec("INSERT INTO action_outcomes (action_id, belief_id) VALUES (5, 5)")

	Database.Exec("INSERT INTO goal_beliefs (goal_id, belief_id) VALUES (0, 3)")
	Database.Exec("INSERT INTO goal_beliefs (goal_id, belief_id) VALUES (1, 4)")
	Database.Exec("INSERT INTO goal_beliefs (goal_id, belief_id) VALUES (2, 0)")
	Database.Exec("INSERT INTO goal_beliefs (goal_id, belief_id) VALUES (3, 4)")
	Database.Exec("INSERT INTO goal_beliefs (goal_id, belief_id) VALUES (4, 5)")

	Database.Exec("INSERT INTO agent_beliefs (agent_id, belief_id) VALUES (1, 6)")
	Database.Exec("INSERT INTO agent_beliefs (agent_id, belief_id) VALUES (2, 6)")

	Database.Exec("INSERT INTO agent_beliefs (agent_id, belief_id) VALUES (3, 1)")
	Database.Exec("INSERT INTO agent_beliefs (agent_id, belief_id) VALUES (4, 1)")
	Database.Exec("INSERT INTO agent_beliefs (agent_id, belief_id) VALUES (5, 1)")
	Database.Exec("INSERT INTO agent_beliefs (agent_id, belief_id) VALUES (6, 1)")
	Database.Exec("INSERT INTO agent_beliefs (agent_id, belief_id) VALUES (7, 1)")

	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (1, 6)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (1, 7)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (2, 6)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (2, 7)")

	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (3, 1)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (3, 2)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (3, 3)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (3, 4)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (3, 5)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (4, 1)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (4, 2)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (4, 3)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (4, 4)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (4, 5)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (5, 1)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (5, 2)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (5, 3)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (5, 4)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (5, 5)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (6, 1)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (6, 2)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (6, 3)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (6, 4)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (6, 5)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (7, 1)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (7, 2)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (7, 3)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (7, 4)")
	Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (7, 5)")

	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (1, 4)")
	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (1, 5)")
	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (2, 4)")
	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (2, 5)")

	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (3, 1)")
	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (3, 2)")
	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (3, 3)")
	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (4, 1)")
	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (4, 2)")
	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (4, 3)")
	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (5, 1)")
	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (5, 2)")
	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (5, 3)")
	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (6, 1)")
	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (6, 2)")
	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (6, 3)")
	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (7, 1)")
	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (7, 2)")
	Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (7, 3)")
}
