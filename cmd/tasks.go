package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mergestat/timediff"
)

const (
	file                    string = "tasks.db"
	CREATE_TABLE_QUERY      string = "CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, task TEXT NOT NULL, created TIMESTAMP DEFAULT CURRENT_TIMESTAMP, done BOOLEAN DEFAULT 0, due TIMESTAMP)"
	DELETE_QUERY            string = "DELETE FROM tasks WHERE id = ?"
	INSERT_QUERY            string = "INSERT INTO tasks (task, due) VALUES (?, ?)"
	SELECT_INCOMPLETE_QUERY string = "SELECT id, task, created, due FROM tasks WHERE done = false"
	SELECT_ALL_QUERY        string = "SELECT * FROM tasks"
	MARK_COMPLETE_QUERY     string = "UPDATE tasks SET done = true WHERE id = ?"
)

func getWriter() *tabwriter.Writer {
	return tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)
}

func printTasks(w *tabwriter.Writer, showCompletion bool) error {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return fmt.Errorf("Error opening database: %w\n", err)
	}
	defer db.Close()

	query := SELECT_INCOMPLETE_QUERY
	if showCompletion {
		query = SELECT_ALL_QUERY
	}

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("Error querying database: %w\n", err)
	}
	defer rows.Close()

	if showCompletion {
		fmt.Fprintf(w, "ID\tTask\tCreated\tDone\tDue\n")
	} else {
		fmt.Fprintf(w, "ID\tTask\tCreated\tDue\n")
	}

	for rows.Next() {
		var (
			id      int
			task    string
			created time.Time
			done    sql.NullBool
			due     *time.Time
		)

		if showCompletion {
			err = rows.Scan(&id, &task, &created, &done, &due)
		} else {
			err = rows.Scan(&id, &task, &created, &due)
		}
		if err != nil {
			return fmt.Errorf("Error scanning row: %w\n", err)
		}

		age := timediff.TimeDiff(created)
		var timeLeft string
		if due != nil {
			timeLeft = timediff.TimeDiff(*due)
		} else {
			timeLeft = "nil"
		}

		if showCompletion {
			completed := "No"
			if done.Valid && done.Bool {
				completed = "Yes"
			}
			fmt.Fprintf(w, "%d\t%s\t%v\t%s\t%s\n", id, task, age, completed, timeLeft)
		} else {
			fmt.Fprintf(w, "%d\t%s\t%v\t%s\n", id, task, age, timeLeft)
		}
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("Error iterating rows: %w\n", err)
	}

	return nil
}

func createTable() error {
	db, err := sql.Open("sqlite3", "tasks.db")
	if err != nil {
		return fmt.Errorf("Error opening db: %w\n", err)
	}
	defer db.Close()

	_, err = db.Exec(CREATE_TABLE_QUERY)
	if err != nil {
		return fmt.Errorf("Error creating db: %w\n", err)
	}

	return nil
}

func List(showCompletion bool) error {
	err := createTable()
	if err != nil {
		return err
	}

	w := getWriter()
	defer w.Flush()

	return printTasks(w, showCompletion)
}

func Add(task string, due *time.Time) error {
	err := createTable()
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return fmt.Errorf("Error opening database: %w\n", err)
	}
	defer db.Close()

	if due == nil {
		_, err = db.Exec(INSERT_QUERY, task, nil)
	} else {
		_, err = db.Exec(INSERT_QUERY, task, due)
	}

	if err != nil {
		return fmt.Errorf("Error inserting todo: %w\n", err)
	}

	w := getWriter()
	defer w.Flush()

	return printTasks(w, false)
}

func Complete(id int) error {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return fmt.Errorf("Error opening db: %w\n", err)
	}
	defer db.Close()

	result, err := db.Exec(MARK_COMPLETE_QUERY, id)
	if err != nil {
		return fmt.Errorf("Error updating row: %w\n", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error getting affected rows: %w\n", err)
	}

	if affectedRows == 0 {
		return fmt.Errorf("No task found with id %d\n", id)
	}

	fmt.Printf("Marked task %d as completed\n", id)
	return nil
}

func Delete(id int) error {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return fmt.Errorf("Error opening db: %w\n", err)
	}
	defer db.Close()

	result, err := db.Exec(DELETE_QUERY, id)
	if err != nil {
		return fmt.Errorf("Error deleting from db: %w\n", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error getting affected rows: %w\n", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No task found with id %d\n", id)
	}

	fmt.Printf("Deleted task %d\n", id)
	return nil
}
