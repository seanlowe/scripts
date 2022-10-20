package cmd

import (
  "fmt"
  "os"
  "time"

  "github.com/spf13/cobra"
)


func init() {
  rootCmd.AddCommand(noteCmd)
}

var noteCmd = &cobra.Command{
  Use:   "note",
  Short: "add a new note",
  Long:  "can create or append to a note",
  Run: func(cmd *cobra.Command, args []string) {
    message := args[0]

    if (!checkDir(topLevelDir)) {
      createDir(topLevelDir)
    }

    startJournalEntry(message)
  },
}

const topLevelDir string = "/home/slowe/Documents/notes"

func buildDateFileName(year int, month int, day int) string {
  return fmt.Sprintf("%d-%d-%d", year, month, day)
}

func startJournalEntry(message string) {
  // check if folder with this week's sunday exists
  date := time.Now()
  currentYear, currentMonth, currentDay := date.Date()
  start := currentDay - int(date.Weekday())

  // build filename for today's date
  filenameToday := buildDateFileName(currentYear, int(currentMonth), currentDay)

  if (start == 0) {
    // if start is 0, revert to previous month, last day
    date = date.AddDate(0, 0, -1)
  }

  if (start < 0) {
    // if start is negative, revert to previous month and
    // increase value of number by 1 to account for 1-based calendars
    // ex: start is -5 -> revert to previous month and subtract from end by 6
    start--
    date = date.AddDate(0, 0, start)
  }

  year, month, day := date.Date()

  // otherwise start is positive and we're in the middle of the month
  if (start > 0) {
    day = start
  }
  
  filenameWeek := buildDateFileName(year, int(month), day)
  
  weekDir := fmt.Sprintf("%s/%s", topLevelDir, filenameWeek)
  weekFolderExists := checkDir(weekDir)

  // if week folder does not exist, create it
  if (!weekFolderExists) {
    createDir(weekDir)
  }

  // check for day entry
  dayEntry := fmt.Sprintf("%s/%s.txt", weekDir, filenameToday)
  dayEntryExists := checkEntry(dayEntry)
  if (dayEntryExists == nil) {
    // if day entry does not exist, create it with message
    createEntry(dayEntry, message)
  } else {
    // if day entry exists, append message to it
    editEntry(dayEntry, message)
  }
}

// create the notes directory according to the path provided
func createDir(path string) {
  err := os.Mkdir(path, 0777)

  if (os.IsExist(err)) {
    fmt.Println("there was an error creating the dir at:", path)

    return
  }

  fmt.Println("dir created at:", path)
}

// check only for the directory in question
func checkDir(path string) bool {
  _, err := os.ReadDir(path)

  return !os.IsNotExist(err)
}

// create the file and add the entry to it
func createEntry(path string, message string) {
  err := os.WriteFile(path, []byte(message), 0666)

  if (err != nil) {
    fmt.Println("there was an error creating and writing the file")
  }
}

// open the file, append to it, then close it
func editEntry(path string, message string) {
  file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0666)

  paddedMessage := fmt.Sprintf("\n%s", message)
  writeable := []byte(paddedMessage)

  if (err != nil) {
    fmt.Println("there was an error opening the file")
  }

  _, err = file.Write(writeable)
  if (err != nil) {
    fmt.Println("there was an error writing the file")
  }

  err = file.Close()
  if (err != nil) {
    fmt.Println("there was an error closing the file")
  }
}

// look for an entry at the path and return the contents
func checkEntry(path string) []byte {
  contents, err := os.ReadFile(path)

  if (os.IsNotExist(err)) {
    return nil
  }

  return contents
}
