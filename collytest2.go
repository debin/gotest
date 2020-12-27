package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gocolly/colly"
	"os"
)


type IpResult struct {
	Code int `json:"code"`
	Data struct {
		City struct {
			GeonameID int `json:"geoname_id"`
			Names     struct {
				De    string `json:"de"`
				En    string `json:"en"`
				Es    string `json:"es"`
				Fr    string `json:"fr"`
				Ja    string `json:"ja"`
				Pt_BR string `json:"pt-BR"`
				Ru    string `json:"ru"`
				Zh_CN string `json:"zh-CN"`
			} `json:"names"`
		} `json:"city"`
		Continent struct {
			Code      string `json:"code"`
			GeonameID int    `json:"geoname_id"`
			Names     struct {
				De    string `json:"de"`
				En    string `json:"en"`
				Es    string `json:"es"`
				Fr    string `json:"fr"`
				Ja    string `json:"ja"`
				Pt_BR string `json:"pt-BR"`
				Ru    string `json:"ru"`
				Zh_CN string `json:"zh-CN"`
			} `json:"names"`
		} `json:"continent"`
		Country struct {
			GeonameID int    `json:"geoname_id"`
			IsoCode   string `json:"iso_code"`
			Names     struct {
				De    string `json:"de"`
				En    string `json:"en"`
				Es    string `json:"es"`
				Fr    string `json:"fr"`
				Ja    string `json:"ja"`
				Pt_BR string `json:"pt-BR"`
				Ru    string `json:"ru"`
				Zh_CN string `json:"zh-CN"`
			} `json:"names"`
		} `json:"country"`
		Location struct {
			AccuracyRadius int     `json:"accuracy_radius"`
			Latitude       float64 `json:"latitude"`
			Longitude      float64 `json:"longitude"`
			TimeZone       string  `json:"time_zone"`
		} `json:"location"`
		RegisteredCountry struct {
			GeonameID int    `json:"geoname_id"`
			IsoCode   string `json:"iso_code"`
			Names     struct {
				De    string `json:"de"`
				En    string `json:"en"`
				Es    string `json:"es"`
				Fr    string `json:"fr"`
				Ja    string `json:"ja"`
				Pt_BR string `json:"pt-BR"`
				Ru    string `json:"ru"`
				Zh_CN string `json:"zh-CN"`
			} `json:"names"`
		} `json:"registered_country"`
		Subdivisions []struct {
			GeonameID int    `json:"geoname_id"`
			IsoCode   string `json:"iso_code"`
			Names     struct {
				En    string `json:"en"`
				Fr    string `json:"fr"`
				Zh_CN string `json:"zh-CN"`
			} `json:"names"`
		} `json:"subdivisions"`
		Traits struct {
			IPAddress string `json:"ip_address"`
		} `json:"traits"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func writeExcle(result *IpResult)  {
	f := excelize.NewFile()

	//title
	title := []string{"ip","国家","省份","城市","经度","纬度"}
	for k, v := range title {
		tmp :=rune(int('A')+k)
		fmt.Println("tmp:",string(tmp)+"1",v)
		f.SetCellValue("Sheet1", string(tmp)+"1", v)
	}

	//内容
	f.SetCellValue("Sheet1", "A2", result.Data.Traits.IPAddress)
	f.SetCellValue("Sheet1", "B2", result.Data.Country.Names.Zh_CN)
	f.SetCellValue("Sheet1", "C2", result.Data.Subdivisions[0].Names.Zh_CN)
	f.SetCellValue("Sheet1", "D2", result.Data.City.Names.Zh_CN)
	f.SetCellValue("Sheet1", "E2", result.Data.Location.Longitude)
	f.SetCellValue("Sheet1", "F2", result.Data.Location.Latitude)


	// Save xlsx file by the given path.
	err := f.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}


}

func main() {
	s := os.Args[0]
	dir, _ := os.Getwd()
	fmt.Printf("path:%s\n",s)
	fmt.Printf("dir:%s\n",dir)

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		//colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})
	c.OnResponse(func(response *colly.Response) {
		//link := e.Attr("href")
		// Print link
		fmt.Printf("Link result: %s\n", response)
		fmt.Printf("Link result body: %s\n", response.Body)

		var myMap = &map[string]interface{}{}
		json.Unmarshal(response.Body,myMap)
		fmt.Printf("Link myMap: %s\n", myMap)

		//var testa = [] byte(`[123,456]`)
		//var myMap2 []interface{}
		//unmarshal := json.Unmarshal(testa, &myMap2)
		//fmt.Printf("Link err:%s,t:%T,myMap2: %s\n", unmarshal,myMap2,myMap2)
		//for k, v := range myMap2 {
		//	fmt.Printf("数组,k:%s,v: %s\n", strconv.Itoa(k),fmt.Sprintf("%v",v))
		//}
		var ipResult  = &IpResult{}
		json.Unmarshal(response.Body,ipResult)
		fmt.Printf("result struct,ipResult:%v \n",ipResult)
		fmt.Printf("result struct,ipResult#:%#v \n",ipResult)
		fmt.Printf("result struct,ipResul0t+:%+v \n",ipResult)

		writeExcle(ipResult)
	})



	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://s.zhidaohu.com/i")
}