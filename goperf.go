package main

import (
    "fmt"
    "flag"
    "io/ioutil"
    "net/http"
    "time"
)

// avg_load_time returns the average response time of the specified uri
// expressed in seconds. The uri is fetched n="times" times and the average
// load time is computed
func avg_load_time(uri string, times int) float64 {

    counter := times
    wait_chan := make(chan time.Duration)

    for ; counter > 0; counter-- {

        go time_uri(uri, wait_chan)
    }

    var total int64
    var res_time time.Duration
    for ; counter < times; counter++ {

        res_time = <-wait_chan
        total += res_time.Nanoseconds()
    }

    return float64(total)/float64(times)/1000000000.0
}

// time_uri fecth the specified uri and send the time.Duration of the operation
// to the specified res chan
func time_uri(uri string, res chan time.Duration) {

    start_time := time.Now()
    get_uri(uri)

    res <- time.Since(start_time)
}

// get_uri tryes to fetch the specified uri and returs its content as bytes
// and/or the error if one has occurred
func get_uri(uri string) (body []byte, err error) {

    resp, err := http.Get(uri)

    if err != nil {

        return nil, err
    }

    defer resp.Body.Close()

    body, err = ioutil.ReadAll(resp.Body)

    return
}

// Usage of the goperf tool
var Usage = func() {
    usage := "goperf --uri http://uri.to.test/ [--requests 10]"
    fmt.Fprintf(os.Stderr, "Usage of %s:\n%s\n", os.Args[0], usage)
    PrintDefaults()
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
