package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Main function to display the menu and handle user input
func main() {
	taskCount := 1000
	reader := bufio.NewReader(os.Stdin)

	// taskList holds the task struct
	var tasksList []Task
	fmt.Println("\n*****************TODO LIST****************")

	for {
		fmt.Println("\n")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			fmt.Println("Invalid choice, please try again.")
			continue
		} else {
			fmt.Printf("\nChoice is : %d\n\n", choice)
			switch choice {
			case 1:
				tasksList = addTask(taskCount, tasksList)
				taskCount++
			case 2:
				listTask(tasksList)
			case 3:
				tasksList = markDone(tasksList)
			case 4:
				tasksList = deleteTask(tasksList)
			case 5:
				exit()
			}
		}

	}
}

// created a struct about the task.
type Task struct {
	taskId      int
	taskName    string
	description string
	isDone      bool
}

// addTask adds the task.
func addTask(taskCount int, tasksList []Task) []Task {
	reader := bufio.NewReader(os.Stdin)
	taskInstance := Task{}
	taskInstance.taskId = taskCount

	fmt.Println("Enter task name: ")
	input, _ := reader.ReadString('\n')
	taskInstance.taskName = input

	fmt.Println("Enter task description: ")
	input, _ = reader.ReadString('\n')
	taskInstance.description = input

	tasksList = append(tasksList, taskInstance)
    fmt.Println("\t\t\t\t******************************")
	fmt.Print("\t\t\t\t  Task added successfully...!\n")
    fmt.Println("\t\t\t\t******************************")


	return tasksList

}

// listTask lists all the task and its details.
func listTask(tasksList []Task) {
    fmt.Println("\t\t\t\t******************************")
	fmt.Print("\t\t\t\t\t\t Your Tasks\n")
    fmt.Println("\t\t\t\t******************************")
    if len(tasksList) == 0 {
        fmt.Print("\t\t\t\t\t\t Your task list is empty\n")
        return;
    }
	for _, value := range tasksList {
		completed := "Not completed"
		if value.isDone {
			completed = "Completed"
		}
		fmt.Print("-------------------------------------------------------\n")
		fmt.Printf("Task Id: %d\n\nTask name: %sTask Description: %sTask Completed: %s\n\n", value.taskId, value.taskName, value.description, completed)
		fmt.Println("-------------------------------------------------------\n\n")
	}
}

// markDone function updates the task status to completed
func markDone(taskList []Task) []Task {
	var taskId int
	var p Task
	ind := 0
	updated := false
	fmt.Println("Select the task from the list to be marked as done.")
	fmt.Println("**************************************************************")
	for _, value := range taskList {
		if !value.isDone {
			fmt.Printf("%d => %s", value.taskId, value.taskName)
		}
	}
	fmt.Println("**************************************************************")
	fmt.Println("Enter the task Id from the list: ")
	fmt.Scanf("%d", &taskId)

	for index, value := range taskList {
		if value.taskId == taskId {
			p = value
			ind = index
			updated = true
		}
	}
	p.isDone = true
	taskList[ind] = p
	if updated {
        fmt.Println("\t\t\t\t******************************")
		fmt.Println("\t\t\t\tStatus updated...!")
        fmt.Println("\t\t\t\t******************************")

	} else {
        fmt.Println("\t\t\t\t******************************")
		fmt.Println("\t\t\t\tInvalid task ID...!")
        fmt.Println("\t\t\t\t******************************")
	}

	return taskList
}

// Deletes the task from the list
func deleteTask(taskList []Task) []Task {
	var taskId int
    done := false;
	fmt.Println("Select the task from the list to be removed.")
	fmt.Println("**************************************************************")
	for _, value := range taskList {
		fmt.Printf("%d => %s", value.taskId, value.taskName)
	}
	fmt.Println("**************************************************************")
	fmt.Println("Enter the task Id from the list: ")
	fmt.Scanf("%d", &taskId)
	for index, value := range taskList {
		if value.taskId == taskId {
            done = true;
			taskList = append(taskList[:index], taskList[index+1:]...)
		}
	}
    if done {
        fmt.Println("\t\t\t\t******************************")
        fmt.Printf("\t\t\t\tTask (%d) removed...!\n", taskId)
        fmt.Println("\t\t\t\t******************************")
    } else {
        fmt.Printf("\n\t\t\t\tInvalid task Id");
    }
    
	return taskList
}

func exit() {
	fmt.Println("Are you sure? (y/n): ")
	input := ""
	fmt.Scanf("%s", &input)
	if input == "y" {
		os.Exit(0)
	} else if input == "n" {
		return
	} else {
		fmt.Println("\nInvalid input...!")
		return
	}
}
