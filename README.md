# Advanced Caesar Cypher
This repository is for programming languages project.

The goal of this project is to implement the advanced Ceasar Cypher by this research paper: https://ieeexplore.ieee.org/document/7749010/.

This implementation uses the basis of ceasar cypher of shifting the different letters in the alphabet one for the other, normally your encription key would tell you how many places to shift for all characters in the text, however ACC iterates the key for each character changing due to the last character encrypted making each character be shifted a different amount.

# How to run

In order to run this code the user needs Golang installed on their computer or container. Golang is a programming language developed by google, it can be downloaded and installed from here https://golang.org/.

After installing Go and having the repository in your files, simply go to that location in cmd and use this command.
```
  go run "file.go"
```

Especifically for this repository there are two go files, encrypt.go and decrypt.go, you need to use encrypt.go before using decrypt.go.

### Arguments

Both programs receive the same arguments, first the number of files to concurrently encrypt/decrypt and then a set of 3 arguments for each file to encrypt, source file, destination file, and encryption key. The encryption key is created by selecting two private keys and a public one, all of them can be and is better if they are non-prime numbers as it creates more possible factorizations of the encryption key.

### Sample Command
```
  go run encrypt.go 3 source1.txt encrypted1.txt 759 source2.txt encrypted2.txt 2240 source3.txt encrypted3.txt 320
  go run decrypt.go 3 encrypted1.txt decrypted1.txt 759 encrypted2.txt decrypted2.txt 2240 encrypted3.txt decrypted3.txt 320
```
It is important that the decryption key is the same as the one used for encrytion.

### Output

Both program generate an output file with the transformed text, from normal to encrypted for encrypt.go and vice versa for decrypt.go

# Dependencies

Golang uses libraries to implement functions into your source code, the difference with other programming languages is that Golang uses repositories to store its libraries, so there is no need to install the libraries by yourself, once needed the will be downloaded to be able to use.

### fmt
Package fmt implements formatted I/O with functions analogous to C's printf and scanf. 
### io/ioutil
Package ioutil implements some I/O utility functions. (Read and write files)
### math
Package math provides basic constants and mathematical functions. 
### os
Package os provides a platform-independent interface to operating system functionality. (Use command line arguments)
### strconv
Package strconv implements conversions to and from string representations of basic data types. 
