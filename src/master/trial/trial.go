package trial

import (
    "time"
)

func CPURuntime(f func(string), ags string) int64 {
    start_time := time.Now().UnixNano()
    f(ags)
    end_time := time.Now().UnixNano()

    return (end_time - start_time)
}

