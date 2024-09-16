# Tasks - CLI Task Manager

## Overview

Tasks is a command-line interface(CLI) application for managing your to-do list directly from the terminal. Built with Go, it provides a simple and efficient way to add, list, complete and delete tasks, all while storing your data locally using SQLite.

## Features

- Add tasks with optional deadlines.
- List all tasks or only incomplete ones.
- Mark tasks as complete.
- Delete tasks.
- Data persistence using SQLite database.

## Installation

### Prerequisites

- Go 1.16 or higher.
- SQLite.

### Steps

1. Clone the repository.

```bash
$ git clone https://github.com/anakinsonone/tasks.git
```

2. Navigate to project directory.

```bash
$ cd tasks
```

3. Build the application.

```bash
$ go build
```

4. Move the binary to a directory in your PATH(e.g.,/usr/local/bin)

```bash
$ sudo mv tasks /usr/local/bin
```

## Usage

### Add a task

```bash
$ tasks add "Complete project proposal"
$ tasks add "Call client" -m 30 # Due in 30 minutes
$ tasks add "Submit report" -r 2 # Due in 2 hours
$ tasks add "Team meeting" -d 1 # Due in 1 day
```

### List tasks

```bash
$ tasks list  # List incomplete tasks
$ tasks ls    # Shorthand for list
$ tasks list -a  # List all tasks, including completed ones
```

### Complete a task

```bash
$ tasks complete 1  # Mark task with ID 1 as completed
$ tasks c 1         # Shorthand for complete
```

### Delete a task

```bash
$ tasks delete 1  # Delete task with ID 1
$ tasks d 1       # Shorthand for delete
```

## Dependencies

- github.com/spf13/cobra - For building the CLI interface
- github.com/mattn/go-sqlite3 - SQLite driver for Go
- github.com/mergestat/timediff - For human-readable time differences
