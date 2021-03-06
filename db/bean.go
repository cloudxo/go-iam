package db

import (
	"gopkg.in/mgo.v2/bson"
)

type AccountBean struct {
	AccountId   bson.ObjectId `bson:"_id"`
	AccountName string        `bson:"name"`
	Password    string        `bson:"password"`
	AccountType int           `bson:"type"`
	CreateDate  string        `bson:"create_date"`
}

type UserBean struct {
	UserId      bson.ObjectId `bson:"_id"`
	UserName    string        `bson:"name"`
	DisplayName string        `bson:"display_name"`
	Phone       string        `bson:"phone"`
	Email       string        `bson:"email"`
	Comments    string        `bson:"comments"`
	Password    string        `bson:"password"`
	Account     string        `bson:"account"`
	CreateDate  string        `bson:"create_date"`
}

type GroupBean struct {
	GroupId    bson.ObjectId `bson:"_id"`
	GroupName  string        `bson:"name"`
	Comments   string        `bson:"comments"`
	Account    string        `bson:"account"`
	CreateDate string        `bson:"create_date"`
}

type GroupUserBean struct {
	JoinId   bson.ObjectId `bson:"_id"`
	GroupId  string        `bson:"group"`
	UserId   string        `bson:"user"`
	JoinDate string        `bson:"join_date"`
}

type PolicyBean struct {
	PolicyId    bson.ObjectId `bson:"_id"`
	PolicyName  string        `bson:"name"`
	PolicyType  int           `bson:"type"`
	Document    string        `bson:"document"`
	Description string        `bson:"description"`
	Version     string        `bson:"version"`
	Account     string        `bson:"account"`
	CreateDate  string        `bson:"create_date"`
	UpdateDate  string        `bson:"update_date"`
}

type PolicyUserBean struct {
	AttachId   bson.ObjectId `bson:"_id"`
	PolicyId   string        `bson:"policy"`
	UserId     string        `bson:"user"`
	AttachDate string        `bson:"attach_date"`
}

type PolicyGroupBean struct {
	AttachId   bson.ObjectId `bson:"_id"`
	PolicyId   string        `bson:"policy"`
	GroupId    string        `bson:"group"`
	AttachDate string        `bson:"attach_date"`
}

type KeyBean struct {
	AccessKeyId     bson.ObjectId `bson:"_id"`
	AccessKeySecret string        `bson:"secret"`
	Status          int           `bson:"status"`
	Entity          string        `bson:"entity"`
	Entitype        int           `bson:"entitype"`
	CreateDate      string        `bson:"create_date"`
}
