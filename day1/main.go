package main
import ( 
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"
    "strconv"
    "sort"
)

func main() {
    
    // Open file
    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    // Remember to close file
    defer f.Close()

    // Read file line by line using scanner
    scanner := bufio.NewScanner(f)

    left := make([]int, 0)
    right := make([]int, 0)

    for scanner.Scan() {
        line := scanner.Text()
        values := strings.Split(line, "   ")
        lVal, _ := strconv.Atoi(values[0])
        rVal, _ := strconv.Atoi(values[1])
        left = append(left, lVal)
        right = append(right, rVal)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    sort.Ints(left)
    sort.Ints(right)

    totalOffset := 0
    for i := 0; i < len(left); i++ {
        offset := left[i] - right[i]
        if offset < 0 {
            offset = -offset
        }
        totalOffset += offset
    }
    fmt.Println(totalOffset)

    // Part 2
    
    simScore := 0
    for _, v := range left {
        count := 0
        for _, b := range right {
            if b == v {
                count++
            }
        }
        simScore += v * count
    }
    fmt.Println(simScore)

    
}

