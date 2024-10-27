## TODO CLI Written in [Cobra](https://github.com/spf13/cobra)

Use this tool to add, list, manage and track your day-to-day priority tasks without leaving the CLI.

### Install

#### For Mac Users

Clone this repository, and in the main folder, run this command:

```
make all
```

It will prompt you to enter your password as `make` executes `sudo` to move the binary file to your device's `bin` folder.

#### For Other Users

```
make build
make run
```

Make sure to move the `bin/todo` and `bin/td` [Optional] file to your device's `bin` folder.

### Autocompletion

Command:

```
todo completion
```

Use:

```
Generate the autocompletion script for todo for the specified shell.
See each sub-command's help for details on how to use the generated script.

Usage:
  todo completion [command]

Available Commands:
  bash        Generate the autocompletion script for bash
  fish        Generate the autocompletion script for fish
  powershell  Generate the autocompletion script for powershell
  zsh         Generate the autocompletion script for zsh
```

### Format

```
todo [COMMAND] [ARG] [--FLAG]
```

or using Alias

```
td [COMMAND] [ARG] [--FLAG]
```

### Commands

#### TODO

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

#### Add

Command:

```
todo add
```

Command Details:

```
Just Type the command "todo add" and it will prompt you for the required fields.

Usage:
  todo add [flags]

Flags:
  -h, --help            help for add
  -s, --status string   Mark the status of a task
                        Status can be a String in List ("Not Started","In Progress","Completed") - Make sure to Add the Double Quotes ""
                        or you can use Short Hands for each possible status value (ns,ip,c)
                        By Default, It will be marked as Completed (default "Not Started")
```

Sample Prompts:

```
Enter Title (Required) :
Go to the gym
Enter Description ([Default] Empty) :
Workout on bicep with side focus on abs.
New Task Added
```

#### List

Command:

```
todo list
```

Sample Use:

```
List all tasks

Usage:
  todo list [flags]

Flags:
  -h, --help            help for list
  -s, --status string   Only List Tasks with this status
```

Sample Output:

```
ID     Task
6     Work on the k8s setup
```

#### Edit

Command:

```
todo edit <TASK_ID>
```

Sample Use:

```
Current Title : Work on the k8s setup
Enter Title (Used Current If Left Empty) :
Go to gym
Current Description : For the freelancing users
Enter Description ([Default] Empty) :
I will do bicep workout with focus on abs.
Status should a [String] in List (Not Started,In Progress,Completed) or Short Hands (ns,ip,c) :
ns
```

#### Mark

Command:

```
todo mark <TODO_ID> -s <STATUS>
```

Sample Use:

```
todo mark 6 -s ip
```

This command marks task with id `6` in state `In Progress`.

#### Delete

Sample Command:

```
todo delete <TASK_ID>
```

Sample Output:

```
Task Deleted Successfully
```

#### Show

Normal `todo list` shows `ID` and `Title` of all tasks. `todo show` prints more details about a task.

Command:

```
todo show <TASK_ID>
```

Sample output of `todo show <ANY_TASK_ID>`:

```
Title : Go to the gym
Description : Workout on bicep with side focus on abs.
Status : Not Started
```

#### Clear

Command:

```
todo clear -s <STATUS>
```

### Data Format Details

- ID `[INT]`
- Title `[String]`
- Description `[String]`
- Status `[String] in List ("Not Started","In Progress","Completed")` or `Short Hands (ns,ip,c)`
