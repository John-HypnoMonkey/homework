package main

import (
    "fmt"
    "os"
    "bufio"
    "net/http"
    "io/ioutil"
    "regexp"
    //"time"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    total := 0
    //тут задаем максимальное количество гоурутин
    max_goroutines := 5
    goroutines := make(chan struct{}, max_goroutines)
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
            goroutines <- struct{}{}
            wg.Add(1)
            data := scanner.Text()

        go func(data string, goroutines <-chan struct{}, wg *sync.WaitGroup, total *int) {
            
            response, err := http.Get(data)
            if err != nil {
                fmt.Println(err)
            }  
            body, errb := ioutil.ReadAll(response.Body)
            if errb != nil {
                fmt.Println(errb)
            }  
            validID := regexp.MustCompile(`go`)
            matches := validID.FindAllStringIndex(string(body), -1)
            fmt.Println("Count for", data, len(matches))
            //время для наглядности
            //time.Sleep(1 * time.Second)    
            *total = *total + len(matches)
            <-goroutines
            wg.Done()  
        }(data, goroutines, &wg, &total)

        
    }
    
    
    wg.Wait()
    close(goroutines)
    fmt.Println("Total", total)

}
