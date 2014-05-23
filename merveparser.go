// (C)2014 by Johannes Amorosa

// Apache License
// Version 2.0, January 2004
// http://www.apache.org/licenses/ 

package main

import (
    "os"
    "fmt"
    "flag"
    "io/ioutil"
    "log"
    "path/filepath"
    "path"
    "strings"
    "strconv"
)

const (
    ProgName    = "merveparser"
    ProgVersion = "0.1"
)

var (
    versionFlag         = flag.Bool("v", false, "Display version number and exit")
    importPathFlag      = flag.String("ip", ".", "Path to textfiles")
    NCPU                = flag.Int("threads", 1, "Max CPU cores")             
)

type Page struct {
   Filename string
   Content string
   Filecount int
   Scanimage byte
}

type Metadata struct {
    Title []string
    Authors []string
    Year int
    Publisher string
    Coverjpeg byte
    Booknumber int

}

type Book struct {
    Items []Page
}

func main() {
    flag.Parse()

    if *versionFlag {
        log.Printf(ProgName, ProgVersion)
        os.Exit(0)
    }

    // Get list of files
    textFileArray := dir(*importPathFlag +"/text")
    imageFileArray := dir(*importPathFlag +"/book")

    // Parse all textfiles into a Book struct that contains Page structs
    MyBook := ReadPages(textFileArray, imageFileArray, NCPU)

    // Cleaning the Pages 
    ParseMetadata(MyBook, NCPU)


    //TODO add metadata
    //TODO Output formats
}

func ParseMetadata(MyBook Book, NCPU *int) {
    return
}

func ReadPages(textFileArray []string, imageFileArray []string, NCPU *int) Book{
    //fmt.Println(fileArray)
    items := []Page{}
    book := Book{items}

    book := gatherText(textFileArray, book, NCPU)
    
    return book
}

func gatherText(textFileArray []string, book Book, NCPU *int) Book{
    for _, f := range textFileArray {
        fmt.Println(path.Base(f))
        content, err := ioutil.ReadFile(f)
        if err != nil {
            //Do something
        }
        // parsing the Filecount from the filename "9783883961507-61.txt" --> 61 (int)
        filecount, _ := strconv.Atoi(strings.SplitAfter(f[0:len(f)-len(filepath.Ext(f))], "-")[1])

        // building Page struct 
        singlepage := Page {Filename: path.Base(f), Content: string(content), Filecount: filecount }
        
        // adding a Page to Book struct
        book.AddItem(singlepage)

    }
    return book
}

// Append a Page to a book
func (book *Book) AddItem(item Page) []Page {
    book.Items = append(book.Items, item)
    return book.Items
}

func dir(thepath string) []string {
    //TODO filter
    var files []string

    filepath.Walk(thepath, func(path string, _ os.FileInfo, _ error) error {
        //fmt.Println(path)
        files = append(files, path)
        return nil
    })

  return files
}
