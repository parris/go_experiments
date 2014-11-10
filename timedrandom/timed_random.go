package main

import (
    "fmt"
    "time"
    "math/rand"
    "flag"
)

var randomCount = flag.Int("randoms", 5, "Number of randoms to generate")

func timedRandom(randomCount int, timedRandomChannel chan uint32, quit chan uint32) {
    for i := 0; i < randomCount; i++ {
        time.Sleep(200 * time.Millisecond)
        timedRandomChannel <- rand.New(rand.NewSource(time.Now().UnixNano())).Uint32()
    }

    quit <- 0
}

func main() {
    flag.Parse()

    timedRandomChannel := make(chan uint32)
    quit := make(chan uint32)

    go timedRandom(*randomCount, timedRandomChannel, quit)

    main:
        for {
            var randomNumber uint32
            select {
                case randomNumber = <-timedRandomChannel:
                    fmt.Println(randomNumber)
                case <-quit:
                    break main
            }
        }

    fmt.Println("you have all the randoms woo!")
}
