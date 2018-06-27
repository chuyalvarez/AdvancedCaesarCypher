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
      go decrypt(os.Args[(i*3)+2],os.Args[(i*3)+3],float64(key),buffer)
  }
  for j := 0; j < n; j++ {
      <-buffer
  }
  fmt.Println("done decrypting")
}

func decrypt(input string,output string,key float64,buffer chan struct{}){
    text, err := ioutil.ReadFile(input) // just pass the file name
    if err != nil {
        fmt.Print(err)
    }
    fmt.Println("opening",input,"for decryption")

    var b []byte = text[:len(text)-2]
    var shift float64  = 1
    var mod float64 = 0

    for i,_ := range b{
        if float64(b[i])<32{

        }else{
          if shift+33==(float64(b[i])){
              b[i] = 32
          }else{
              shift = math.Mod(key,94)
              if shift==0{
                  key += 1
                  shift = math.Mod(key,94)
              }
              key = key+math.Floor(key/shift)
              mod = math.Mod(float64(b[i])-33-shift,94)

              if mod<=0 {
                  b[i]=byte(mod+127)
              }else{
                  b[i]=byte(mod+33)
              }
            }
        }
    }

    ioutil.WriteFile(output, text, 0644)
    fmt.Println("file",output,"decrypted")
    buffer <- struct{}{}
}
