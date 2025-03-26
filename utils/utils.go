package utils

import (
    "fmt"
    "strconv"
    "strings"
)

// Convert seconds to "min:sec" format
func FormatDuration(seconds int32) string {
    minutes := seconds / 60
    sec := seconds % 60
    return fmt.Sprintf("%d:%02d", minutes, sec)
}

// Convert bytes to MB with 2 decimal places
func FormatSize(bytes int64) string {
    sizeMB := float64(bytes) / (1024 * 1024)
    return fmt.Sprintf("%.2f MB", sizeMB)
}

func SplitAndFormat(input string) (int, int32) {
    parts := strings.Split(input, "/")
    firstPart, _ := strconv.Atoi(parts[0])
    secondPart, _ := strconv.Atoi(parts[1])
    formattedFirst := -1000000000000 - firstPart
    return formattedFirst, int32(secondPart)
}