package main
import (
    "fmt"
    "os"
    "strings"
    "regexp"
    "bufio"
)

func main() {
    input := readFile("input.txt")
    
    xm := regexp.MustCompile(`XMAS`)
    mx := regexp.MustCompile(`SAMX`)
    total := 0
    mastotal := 0
    for i, line := range input {
        fw := xm.FindAllString(line, -1) // All xmas forwards
        bw := mx.FindAllString(line, -1) // All xmas backwards (samx)
        total += len(fw) + len(bw)
        for j, _ := range line {
            word := ""
            total += checkXmas(input, word, i, j, -1) + checkXmas(input, word, i, j, 1) // Diagonals
            total += checkXmas(input, word, i, j, 0) // Verticals
            // Part Two
            mastotal += checkMas(input, word, i, j, 1)
        }
    }
    fmt.Println(total)
    fmt.Println(mastotal)
    
}

func checkMas(input []string, word string, i, j, direction int) int {
    if !(strings.HasPrefix("MAS", word) || strings.HasPrefix("SAM", word))  {
        return 0
    } else if word == "SAM" || word == "MAS" {
        if direction > 0 {
            return checkMas(input, "", i-3, j-1, -direction)
        } else {
            return 1
        }
    } else if len(word) > 3 || i == len(input) || j < 0 || j == len(input[i]) {
        return 0
    }
    word += string(input[i][j])
    if direction > 0 {
        return checkMas(input, word, i+1, j+1, direction)
    } else {
        return checkMas(input, word, i+1, j-1, direction)
    }
    return 0
}

func checkXmas(input []string, word string, i, j, direction int) int {
    if !(strings.HasPrefix("XMAS", word) || strings.HasPrefix("SAMX", word)) {
        return 0
    } else if word == "XMAS" || word == "SAMX" {
        return 1
    } else if len(word) > 4 {
        return 0
    } else if i == len(input) || j < 0 || j == len(input[i]) {
        return 0
    }

    word += string(input[i][j])
    if direction > 0 {
        return checkXmas(input, word, i+1, j+1, direction) // South East
    } else if direction < 0 {
        return checkXmas(input, word, i+1, j-1, direction)      // South West
    } else {
        return checkXmas(input, word, i+1, j, direction)
    }
    return 0
}

func readFile(filename string) []string {
    f, _ := os.Open(filename)
    defer f.Close()

    var lines = []string{}
    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines
}
