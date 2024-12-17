package main
import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "regexp"
)

func main() {

    input := readFile("input.txt")
    m1 := regexp.MustCompile(`mul\(\d+,\d+\)`)
    total := 0
    lines := m1.FindAllString(input, -1)

    m2 := regexp.MustCompile(`[^\d,]`)
    for _, line := range lines {
        a := m2.ReplaceAllString(line, "")
        b := strings.Split(a, ",")
        total += atoiIgnoreErr(b[0]) * atoiIgnoreErr(b[1])
    }
    fmt.Println(total)

    // Part Two

    total = 0
    m3 := regexp.MustCompile(`(mul\(\d+,\d+\))|(do(n't)?\(\))`)
    lines = m3.FindAllString(input, -1)
    do := true
    for _, line := range lines {
        if line == "do()" {
            do = true
        } else if line == "don't()"{
            do = false
        } else if do {
            a := m2.ReplaceAllString(line, "")
            b := strings.Split(a, ",")
            total += atoiIgnoreErr(b[0]) * atoiIgnoreErr(b[1])
        }
    }
    fmt.Println(total)
}

func readFile(filename string) string {
    f, _ := os.ReadFile(filename)
    return string(f)
}

func atoiIgnoreErr(s string) int {
    n, _ := strconv.Atoi(s)
    return n
}
