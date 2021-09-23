package controllers

import (
	"context"
	"exam/core"
	"exam/models"
	"exam/models/req"
	"exam/utils/json"
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
)

type PagerController struct {
	core.BaseController
}

func (c PagerController) List() {
	var papers []models.Paper
	core.GetOrm().Raw("select * from paper order by gmt_created desc").QueryRows(&papers)
	c.Success(papers)
}

func (c PagerController) Create() {
	paperReq := c.ParseFromJsonParam(req.PaperCreateReq{}).(req.PaperCreateReq)
	marshal, _ := json.ToStr(paperReq)
	fmt.Println("创建问卷参数 -> ", marshal)
	err := core.GetOrm().DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		// step1:新增试卷
		paper := models.Paper{
			Title: paperReq.Title,
		}
		var err error
		err = createPaper(err, txOrm, paper)
		if err != nil {
			return err
		}

		// step2:新增试卷题目
		for _, question := range paperReq.Questions {
			var insertQuestion, err = createQuestions(paper, question, err, txOrm)
			if err != nil {
				return err
			}

			// step3:新增试卷题目的选项
			err = createQuestionOptions(question, insertQuestion.Id, err, txOrm)
			if err != nil {
				return err
			}

		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	c.Success()
}

func createQuestionOptions(question req.Questions, insertQuestionId uint64, err error, txOrm orm.TxOrmer) error {
	for index, option := range question.Options {
		insertOption := models.QuestionOption{
			QuestionId: insertQuestionId,
			Title:      option.Title,
			Seq:        index + 1,
		}
		_, err = txOrm.Insert(&insertOption)
		if err != nil {
			logs.Error("新增试卷题目选项失败 -> %s", err)
			return err
		}
	}
	return nil
}

func createQuestions(paper models.Paper, question req.Questions, err error, txOrm orm.TxOrmer) (models.Question, error) {
	insertQuestion := models.Question{
		PaperId: paper.Id,
		Title:   question.Title,
		Type:    question.Type,
		Score:   5,
		Answer:  strconv.Itoa(question.Answer),
	}
	_, err = txOrm.Insert(&insertQuestion)
	if err != nil {
		logs.Error("新增试卷题目失败 -> %s", err)
		return models.Question{}, err
	}
	return insertQuestion, nil
}

func createPaper(err error, txOrm orm.TxOrmer, paper models.Paper) error {
	_, err = txOrm.Insert(&paper)
	if err != nil {
		logs.Error("新增试卷失败 -> %s", err)
		return err
	}
	return nil
}
