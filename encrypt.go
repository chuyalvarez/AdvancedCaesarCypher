package main

import (
    "fmt"
    "io/ioutil"
    "math"
    "os"
    "strconv"
)

func main() {
    n,_ := strconv.Atoi(os.Args[1])
    buffer := make(chan struct{})
    for i := 0; i < n; i++ {
        key,_:=strconv.Atoi(os.Args[(i*3)+4])
        go encrypt(os.Args[(i*3)+2],os.Args[(i*3)+3],float64(key),buffer)
    }
    for j := 0; j < n; j++ {
        <-buffer
    }
    fmt.Println("done encrypting")

}

func encrypt(input string,output string,key float64,buffer chan struct{}){

    text, err := ioutil.ReadFile(input)
    if err != nil {
        fmt.Print(err)
    }
    fmt.Println("opening",input,"for encryption")

    var b []byte = text[:len(text)-2]
    var shift float64  = 1

    for i,_ := range b{
        if float64(b[i])<32{

        }else if b[i]==32{
            b[i]= byte(33+shift)
        }else{
            shift = math.Mod(key,94)
            if shift==0{
                key += 1
                shift = math.Mod(key,94)
            }
            key = key+math.Floor(key/shift)
            b[i]=byte(math.Mod(float64(b[i])-33+shift,94)+33)
        }
    }

    ioutil.WriteFile(output, text, 0644)
    fmt.Println("file",output,"encrypted")
    buffer <- struct{}{}
}

/*

*/
