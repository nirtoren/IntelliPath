# IntelliPath

IntelliPath is an intelligent `cd` command for the terminal that allows you to navigate quickly to frequently visited paths without needing to type the exact name of the directory. This project, written in Go, leverages fuzzy search and Levenshtein distance algorithms to make directory navigation more intuitive and efficient.


## Features

- **Intelligent Path Navigation**: Quickly navigate to frequently visited paths using fuzzy search and Levenshtein distance algorithms.
- **Database Management**: Stores paths and their usage scores to optimize navigation.
- **Automatic Cleanup**: Periodically cleans up the database to remove paths that haven't been accessed in the last X days.

## Environmental variables usage 

IntelliPath uses **two**  environmental variables which should be added to your `.bashrc` file:

|                |Example                          |Usage                         |
|----------------|-------------------------------|-----------------------------|
|_INTELLIPATH_DIR|`/usr/loca/bin`            |This will be the installation folder            |
|_INTELLIPATH_DB_DTIMER          |`5`, `1`, etc..            |The maximum days for deleting un-touched paths            |

## How to start

1. installation:
	Download the **intellipath.tar.gz** with the **install.sh** file.
	Run: `sudo -E ./install.sh` 
2. Initialization:
	Run: `intellipath init`
3. Start using:
	`icd` command is now available for you :)
