package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func main() {
    lines := readFile("input.txt")
    safeCount := 0

    for _, line := range lines {
        data := strings.Split(line, " ")
        prev, _ := strconv.Atoi(data[0])
        safeCount += check(data[1:], prev, 0)
    }
    fmt.Println(safeCount)

    // Part 2

    safeCount = 0
    for _, line := range lines {
        data := strings.Split(line, " ")
        fmt.Println(data)
        prev, _ := strconv.Atoi(data[0])
        if (check(data[1:], prev, 0) == 0) {
            // Check all subslices with one element removed
            for i := 0; i < len(data); i++ {
                splice := make([]string, 0, len(data) - 1)
                splice = append(splice, data[:i]...)
                splice = append(splice, data[i+1:]...)
                fmt.Println(splice)
                prev, _ = strconv.Atoi(splice[0])
                if check(splice[1:], prev, 0) == 1 {
                    safeCount++
                    break;
                }
            }
        } else {
            safeCount++
        }
    }
    fmt.Println(safeCount)

    
}

func check(data []string, prev, flag int) int {
    if (len(data) == 0) {
        return 1
    }
    num, _ := strconv.Atoi(data[0])
    diff := prev - num
    if diff < 0 {
        diff = -diff
    }
    if diff > 3 || diff == 0{
        fmt.Println("Diff")
        return 0
    }
    if flag > 0 && prev < num {
        return check(data[1:], num, flag,)
    } else if flag < 0 && prev > num{
        return check(data[1:], num, flag)
    } else {
        if flag != 0 {
            fmt.Println("Order swap")
            return 0
        } else {
            if prev > num {
                flag = -1
            } else {
            flag = 1
            }
        }
        return check(data[1:], num, flag)
    }
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
