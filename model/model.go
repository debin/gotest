package model

type Douban_top_movie struct {
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

func (*Douban_top_movie) TableName() string {
	return "douban_top_movie"
}
