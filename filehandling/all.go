package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "sample.txt"

	// 1. Create a file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("❌ Error creating file:", err)
		return
	}
	_, err = file.WriteString("Hello, this is the first line.\n")
	if err != nil {
		fmt.Println("❌ Error writing to file:", err)
		file.Close()
		return
	}
	file.Close()
	fmt.Println("✅ File created and written to.")

	// 2. Append to the file
	file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("❌ Error opening file for appending:", err)
		return
	}
	_, err = file.WriteString("This line is appended.\n")
	if err != nil {
		fmt.Println("❌ Error appending to file:", err)
		file.Close()
		return
	}
	file.Close()
	fmt.Println("✅ Line appended to file.")

	// 3. Read entire file using os.ReadFile
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("❌ Error reading file:", err)
		return
	}
	fmt.Println("\n📖 File content using os.ReadFile:")
	fmt.Println(string(data))

	// 4. Read line-by-line using bufio
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("❌ Error opening file for line-by-line reading:", err)
		return
	}
	fmt.Println("\n📖 File content line-by-line:")
	scanner := bufio.NewScanner(readFile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("❌ Error scanning file:", err)
	}
	readFile.Close() // Close before deleting the file

	// 5. Delete the file
	err = os.Remove(filename)
	if err != nil {
		fmt.Println("❌ Error deleting file:", err)
		return
	}
	fmt.Println("\n🗑️ File deleted successfully.")
}
