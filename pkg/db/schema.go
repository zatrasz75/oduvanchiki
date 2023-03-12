package schema

import "time"

type Quiestions struct {
	Id       int    `gorm:"primaryKey"`
	Question string `gorm:"not null;size:255"`
}

type Answers struct {
	Id          int        `gorm:"primaryKey"`
	Answer1     string     `gorm:"not null;size:255"`
	Answer2     string     `gorm:"not null;size:255"`
	Answer3     string     `gorm:"not null;size:255"`
	Answer4     string     `gorm:"not null;size:255"`
	Quiestionid int        `gorm:"not null"`
	Quiestions  Quiestions `gorm:"foreignKey:Quiestionid"`
}

type Correctanswers struct {
	Id            int        `gorm:"primaryKey"`
	Questionid    int        `gorm:"not null"`
	Answercorrect string     `gorm:"not null;size:255"`
	Correct       bool       `gorm:"not null;default:true"`
	Quiestions    Quiestions `gorm:"foreignKey:Questionid"`
}

type Clientusers struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"not null;size:255"`
	Ip       string `gorm:"not null;size:17"`
	Browser  string `gorm:"not null;size:20"`
	Platform string `gorm:"size:20"`
}

type Quizes struct {
	Id          int         `gorm:"primaryKey"`
	Userid      int         `gorm:"not null"`
	Started     time.Time   `gorm:"not null"`
	Clientusers Clientusers `gorm:"foreignKey:Userid"`
}

type Results struct {
	Id         int        `gorm:"primaryKey"`
	Questionid int        `gorm:"not null"`
	Answerid   int        `gorm:"not null"`
	Quizid     int        `gorm:"not null"`
	Answered   time.Time  `gorm:"not null;size:6"`
	Point      int        `gorm:"not null;default:0"`
	Answers    Answers    `gorm:"foreignKey:Answerid"`
	Quiestions Quiestions `gorm:"foreignKey:Questionid"`
	Quizes     Quizes     `gorm:"foreignKey:Quizid"`
}
