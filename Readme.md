### TODO CLI Written in [Cobra](https://github.com/spf13/cobra)

Use this tool to add, list, manage and track your day-to-day priority tasks without leaving the CLI.

#### Install

##### For Mac Users

Clone this repository, and in the main folder, run this command:

```
make all
```

##### For Other Users

```
make build
make run
```

Make sure to move the `bin/todo` file to your device's `bin` folder.

#### Format

```
todo [COMMAND] [ARG] [--FLAG]
```

or using Alias

```
td [COMMAND] [ARG] [--FLAG]
```

#### Commands

##### TODO

Command:

```
todo
```

Sample Output:

```

  _____ ___  ____   ___
 |_   _/ _ \|  _ \ / _ \
   | || | | | | | | | | |
   | || |_| | |_| | |_| |
   |_| \___/|____/ \___/

Use this tool to add, list, manage and track your day-to-day priority tasks without leaving the CLI.

https://github.com/maniraja1122/todo-cli

Usage:
  todo [command]

Aliases:
  todo, td

Available Commands:
  add         Add a new task
  completion  Generate the autocompletion script for the specified shell
  delete      Delete a task
  edit        Edit a task
  help        Help about any command
  list        List all tasks
  mark        Mark the status of a task
  show        Show a task

Flags:
  -h, --help   help for todo

Use "todo [command] --help" for more information about a command.
```

##### Add

Command:

```

```

##### List

##### Edit

##### Mark

##### Delete

##### Show

#### Data Format Details

- ID `[INT]`
- Title `[String]`
- Description `[String]`
- Status `[String] in List ("Not Started","In Progress","Completed")` or `Short Hands (ns,ip,c)`
