package main

import (
    "fmt"
    "log"
    "strings"
    "github.com/PuerkitoBio/goquery"
    "encoding/json"
    "flag"
)

func searchScrape() {
    //get string from argument
    keyWordStr := flag.String("keyword", "", "Text to parse")
    flag.Parse()
    parameter := *keyWordStr
    //replace space with '+' character from strings packages
    parameters := strings.Replace(parameter, " ", "+", -1)
    host := "https://www.blibli.com/search?s="

    url := host + parameters
    fmt.Println(url)
    doc, err := goquery.NewDocument(url)
    test := doc.Find("title").Contents().Text()
    fmt.Println(test)
    if err != nil {
        log.Fatal(err)
    }
    doc.Find("section .product-list .row .product-detail-wrapper").Each(func(index int, item *goquery.Selection) {

         title := item.Find(".product-title").Text()
         titlebersih := strings.TrimSpace(title)
         hargaNormal := strings.TrimSpace(item.Find(".product-price .old-price .old-price-text").Text())
         hargaDiskon := strings.TrimSpace(item.Find(".product-price .new-price .new-price-text").Text())
         //mappping string
         slcD := map[string]string{"title": titlebersih, "harga_normal": hargaNormal, "harga_diskon": hargaDiskon}
         slcB, _ := json.Marshal(slcD)

         fmt.Println(string(slcB))
         //fmt.Printf("Model - '%s' Harga Normal '%s' - Harga Diskon '%s' \n ", titlebersih, hargaNormal, hargaDiskon)
    })


}



func main(){
    searchScrape()
}
