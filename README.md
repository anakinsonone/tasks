# Tasks

## A cli app for managing tasks in the terminal.

## Available Commands

### 1. Add

Add a new task to the database.

```bash
$ tasks add <task description>
# optional flags
# add deadline in minutes
$ tasks add <task description> -m/--minutes <number>
# add deadline in hours
$ tasks add <task description> -r/--hours <number>
# add deadline in days
$ tasks add <task description> -d/--days <number>
```

### 2. List

Fetch all incomplete tasks from the database.

```bash
$ tasks list
# or
$ tasks ls
# optional flag to fetch all tasks(complete ones as well).
$ tasks list -a
```

### 3. Complete

Mark a task as `completed`.

```bash
$ tasks complete <task_id>
# or
$ tasks c <task_id>
```

### 4. Delete

Delete a particular task from the database.

```bash
$ tasks delete <task_id>
# or
$ tasks d <task_id>
```

## Notable packages used

- `strconv` - for turning types into strings and vice versa.
- `text/tabwriter` - for writing out tab aligned output.
- `github.com/spf13/cobra` - for the command line interface.
- `github.com/mattn/go-sqlite3` - for providing sqlite drivers.
- `github.com/mergestat/timediff` - for displaying relative, friendly time differences.
