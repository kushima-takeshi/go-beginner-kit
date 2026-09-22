package practice

func Level(minutes int) int {
    // ここに実装
    if minutes >= 60 {
        return 2
    } else if minutes  <= 0 {
        return 0
    } else {
        return 1
    }
}