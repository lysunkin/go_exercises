package hamming

import (
    "errors"
)

func Distance(a, b string) (int, error) {
    diff := 0

    if len(a) != len(b) {
        return -1, errors.New("different length")
    }

    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            diff++
        }
    }

    return diff, nil
}
