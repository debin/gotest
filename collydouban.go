package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
	"sync"
	"time"
)


type DoubanTopMovie struct {
	Id           int     `db:"id" json:"id"`
	Rank         int     `db:"rank" json:"rank"`
	Title        string  `db:"title" json:"title"`
	TitleEn      string  `db:"title_en" json:"title_en"`
	TitleOther   string  `db:"title_other" json:"title_other"`
	Desc         string  `db:"desc" json:"desc"`
	RatingNum    float64 `db:"rating_num" json:"rating_num"`
	RatingPeople int     `db:"rating_people" json:"rating_people"`
	Quote        string  `db:"quote" json:"quote"`
	SubjectId    int     `db:"subject_id" json:"subject_id"`
	Subject      string  `db:"subject" json:"subject"`
	Img          string  `db:"img" json:"img"`
	CreateTime   string  `db:"create_time" json:"create_time"`
	UpdateTime   string  `db:"update_time" json:"update_time"`
}


const (
	//HOST     = "qq10086.zhidaohu.com"
	//PORT     = 3307
	//USER     = "debin"
	//PASSWORD = "debin10086"
	//DBNAME   = "test"
	//SSLMODE  = "disable"
)

func Query(db *sqlx.DB,title string) DoubanTopMovie{
	var movie  DoubanTopMovie
	err := db.Get(&movie, "select * from douban_top_movie where title=?",title)
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
	if err == sql.ErrNoRows {
		fmt.Printf("not found data\n")
		return movie
	}

	//for _, user := range mobie {
	//	fmt.Println(user.Title)
	//}
	return movie
}




func Insert(db *sqlx.DB, item *DoubanTopMovie) {
	fmt.Printf("item:%+v\n",item)
	result, err := db.Exec("insert into douban_top_movie(rank,title,title_en,title_other,`desc`,rating_num,rating_people,quote,subject_id,subject,img,create_time,update_time) values (?,?,?,?,?,?,?,?,?,?,?,?,?)", item.Rank,item.Title,item.TitleEn,item.TitleOther,
		item.Desc,item.RatingNum,item.RatingPeople,item.Quote,item.SubjectId,item.Subject,item.Img,time.Now(),time.Now())
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("last insert id:%d affect rows:%d\n", id, affected)
}



func doChn(chn chan colly.HTMLElement, wg *sync.WaitGroup)  {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",USER,PASSWORD,HOST,PORT,DBNAME)
	db, connErr := sqlx.Connect("mysql", dsn)
	if (connErr!=nil){
		fmt.Printf("connect err: %#v\n", connErr)
		panic(connErr)
	}
	defer db.Close()

	//waitGroup.Add(1)
	for  {
		item, _  := <-chn
		//fmt.Printf("item:%+v\n",item)

		rankSec := item.DOM.Find(".pic em")
		rank := rankSec.Text()

		linkSec := item.DOM.Find(".pic a")
		link,_ := linkSec.Attr("href")
		link = strings.Trim(link,"/")
		index := strings.LastIndex(link, "/")

		subIdStr := link[index+1:len(link)]
		subId, _ := strconv.Atoi(subIdStr)

		picSec := item.DOM.Find(".pic img")
		pic,_ := picSec.Attr("src")

		hdSec := item.DOM.Find(".info .hd a")
		titleSpan := hdSec.Children().Eq(0)
		titleEnSpan := hdSec.Children().Eq(1)
		titleOtherSpan := hdSec.Children().Eq(2)
		title := titleSpan.Text()
		titleEn := titleEnSpan.Text()
		titleEn = strings.TrimSpace(titleEn)
		titleOther := titleOtherSpan.Text()
		titleOther = strings.TrimSpace(titleOther)

		bdSec := item.DOM.Find(".info .bd")
		descSec := bdSec.Find("p").Eq(0)
		desc := descSec.Text()
		desc = strings.TrimSpace(desc)

		starNumSec := bdSec.Find(".star .rating_num")
		starNum := starNumSec.Text()
		starNumFloat,_:= strconv.ParseFloat(starNum,2)

		poopleSec := bdSec.Find(".star span").Last()
		poople := poopleSec.Text()
		poopleInt := strings.TrimRight(poople,"人评价")

		quoteSec := bdSec.Find(".quote .inq")
		quote := quoteSec.Text()


		fmt.Printf("rank:%+v\n",rank)
		fmt.Printf("link:%+v\n",link)
		fmt.Printf("pic:%+v\n",pic)
		fmt.Printf("subId:%+v\n",subId)
		fmt.Printf("title:%+v\n",title)
		fmt.Printf("titleEn:%+v\n",titleEn)
		fmt.Printf("titleOther:%+v\n",titleOther)
		fmt.Printf("desc:%+v\n",desc)
		fmt.Printf("starNumFloat:%+v\n",starNumFloat)
		fmt.Printf("poopleInt:%+v\n",poopleInt)
		fmt.Printf("quote:%+v\n",quote)
		fmt.Printf("\n")

		movie := Query(db, title)
		if movie.Id==0 {
			topMovie := new(DoubanTopMovie)
			topMovie.Rank,_ = strconv.Atoi(rank)
			topMovie.Title = title
			topMovie.TitleEn = titleEn
			topMovie.TitleOther = titleOther
			topMovie.TitleOther = titleOther
			topMovie.Desc = desc
			topMovie.RatingNum = starNumFloat
			topMovie.RatingPeople,_ = strconv.Atoi(poopleInt)
			topMovie.Quote = quote
			topMovie.SubjectId = subId
			topMovie.Subject = link
			topMovie.Img = pic
			Insert(db,topMovie)
		}

		wg.Done()
	}

}

func main() {


	chn := make(chan colly.HTMLElement,1000)
	var wg  sync.WaitGroup = sync.WaitGroup{}

	go doChn(chn, &wg)

	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	c.Limit(&colly.LimitRule{DomainGlob:  "*.douban.*", Parallelism: 2})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	//c.OnHTML(".hd", func(e *colly.HTMLElement) {
	//	fmt.Println(strings.Split(e.ChildAttr("a", "href"), "/")[4],
	//		strings.TrimSpace(e.DOM.Find("span.title").Eq(0).Text()))
	//})
	c.OnHTML(".item", func(e *colly.HTMLElement) {
		//fmt.Println("item:",*e)
		chn<-*e
		wg.Add(1)
	})

	c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.Visit("https://movie.douban.com/top250?start=0&filter=")
	c.Wait()

	wg.Wait()


}

