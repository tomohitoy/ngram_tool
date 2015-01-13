package main
import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "strings"
    "errors"
)


type Dataset struct {
    ID int `json:"id"`
    Text string `json:"text"`
}

func main(){
    target_string := "徒然なるままに日暮し硯に向かいて心に移り行くよしなしごとをそこはかとなく書きつくれば怪しうこそ物狂おしけれ"
    fmt.Println("Raw text: ",target_string)
    unigrams, err := ngram(target_string,1)
    if err!=nil{
        fmt.Println(err)
    }
    fmt.Println("Unigrams: ",unigrams)
    bigrams, err := ngram(target_string,2)
    if err!=nil{
        fmt.Println(err)
    }
    fmt.Println("Bigrams: ",bigrams)
    trigrams, err := ngram(target_string,3)
    if err!=nil{
        fmt.Println(err)
    }
    fmt.Println("Trigrams: ",trigrams)
    fail, err := ngram(target_string,100)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Fail: ",fail)

    // Loading jsonfile
    file, err := ioutil.ReadFile("./input/test.json")
    fmt.Println(string(file))
    if err!=nil {
        fmt.Println("File error: ", err)
    }
    var datasets []Dataset
    json_err := json.Unmarshal(file, &datasets)
    if err!=nil{
        fmt.Println("Format Error: ", json_err)
    }
    fmt.Println(datasets)
    fmt.Println(datasets[0].Text)
    fmt.Println(datasets[2].ID)
}

func ngram(target_text string, n int) ([]string, error) {
    sep_text := strings.Split(target_text,"")
    var ngrams []string
    if len(sep_text)<n{
        err := errors.New("Error: Input string's length is less than n value")
        return nil, err
    }
    for i:=0; i<(len(sep_text)-n+1); i++{
        ngrams = append(ngrams,strings.Join(sep_text[i:i+n],""))
    }
    return ngrams, nil
}
