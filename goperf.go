package main

import (
    "fmt"
    "flag"
    "io/ioutil"
    "net/http"
    "time"
)

func avg_load_time(uri string, times int) float64 {

    counter := times
    wait_chan := make(chan time.Duration)

    for ; counter > 0; counter-- {

        go go_get_uri(uri, wait_chan)
    }

    var total int64
    var res_time time.Duration
    for ; counter < times; counter++ {

        res_time = <-wait_chan
        total += res_time.Nanoseconds()
    }

    return float64(total)/float64(times)/1000000000.0
}

func go_get_uri(uri string, res chan time.Duration) {

    start_time := time.Now()
    get_uri(uri)

    res <- time.Since(start_time)
}

func get_uri(uri string) (body []byte, err error) {

    resp, err := http.Get(uri)

    if err != nil {

        return nil, err
    }

    defer resp.Body.Close()

    body, err = ioutil.ReadAll(resp.Body)

    return
}

func main(){

    var uri string
    var requests int
    flag.StringVar(&uri, "uri", "", "specify URI to test")
    flag.IntVar(&requests, "requests", 100, "specify the number of requests")
    flag.Parse()

    fmt.Printf("starting to test %s with %d requests\n", uri, requests)
    avg := avg_load_time(uri, requests)
    fmt.Printf("average load time for uri %s : %f\n", uri, avg)
}
