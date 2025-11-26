package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    fmt.Println("Reading Time Estimator")
    fmt.Println("=====================")
    fmt.Println("Paste your text below (end with Enter):")

    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.TrimSpace(text)

    words := countWords(text)

    // Average adult reading speed in English ~200-250 wpm (words per minute)
    const avgWPM = 200

    minutes := float64(words) / float64(avgWPM)
    seconds := minutes * 60

    fmt.Printf("\nWord count: %d\n", words)
    fmt.Printf("Estimated reading time: %s min (%.0f seconds) at %d wpm\n", formatMinutes(minutes), seconds, avgWPM)

    fmt.Println("\nWant to try a different reading speed? Type a new wpm (or just press Enter to quit):")
    fmt.Print("> ")
    wpmString, _ := reader.ReadString('\n')
    wpmString = strings.TrimSpace(wpmString)
    if wpmString != "" {
        customWPM, err := strconv.Atoi(wpmString)
        if err == nil && customWPM > 0 {
            customMin := float64(words) / float64(customWPM)
            customSec := customMin * 60
            fmt.Printf("At %d wpm: %s min (%.0f seconds)\n", customWPM, formatMinutes(customMin), customSec)
        } else {
            fmt.Println("Invalid input, quitting.")
        }
    }
}

func countWords(text string) int {
    words := strings.Fields(text)
    return len(words)
}

// Gives minutes as X min Y sec (rounded)
func formatMinutes(min float64) string {
    minInt := int(min)
    sec := int((min - float64(minInt)) * 60)
    if minInt == 0 {
        return fmt.Sprintf("%d sec", sec)
    }
    if sec == 0 {
        return fmt.Sprintf("%d min", minInt)
    }
    return fmt.Sprintf("%d min %d sec", minInt, sec)
}
