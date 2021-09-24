package controllers

import (
	"context"
	json2 "encoding/json"
	"exam/core"
	paperDao "exam/dao/paper"
	questionDao "exam/dao/question"
	questionOptionDao "exam/dao/questionoption"
	"exam/models"
	"exam/models/req"
	"exam/models/resp"
	jsonUtils "exam/utils/json"
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
)

type PagerController struct {
	core.BaseController
}

func (c PagerController) Info() {
	jsonParam := c.GetJsonParam()
	paperId, _ := jsonParam["id"].(json2.Number).Int64()
	paperInfoResp := resp.PaperInfoResp{}

	// step1:查询试卷
	paper := paperDao.SelectById(paperId)
	paperInfoResp.Id = paper.Id
	paperInfoResp.Title = paper.Title

	// step2:查询试卷题目
	var questions = questionDao.SelectListByPaperId(paperId)

	// step3:查询试卷题和选项
	questionIds := extractQuestionIds(questions)
	var questionOptions = questionOptionDao.SelectListByQuestionIds(questionIds)
	questionOptionMap := questionOptionsGroupByQuestionId(questionIds, questionOptions)

	// step4:组装响应体
	paperInfoResp.Questions = assembleQuestionResp(questions, questionOptionMap)
	c.Success(paperInfoResp)
}

func extractQuestionIds(questions []models.Question) []int64 {
	var questionIds []int64
	for _, item := range questions {
		questionIds = append(questionIds, item.Id)
	}
	return questionIds
}

func assembleQuestionResp(questions []models.Question, questionOptionMap map[int64][]resp.QuestionOption) []resp.Questions {
	var questionResp []resp.Questions
	for _, item := range questions {
		questionResp = append(questionResp, resp.Questions{
			Id:      item.Id,
			Answer:  item.Answer,
			Options: questionOptionMap[item.Id],
			Title:   item.Title,
			Type:    item.Type,
		})
	}
	return questionResp
}

func questionOptionsGroupByQuestionId(questionIds []int64, questionOptions []models.QuestionOption) map[int64][]resp.QuestionOption {
	questionOptionMap := make(map[int64][]resp.QuestionOption, len(questionIds))
	for _, item := range questionOptions {
		_, ok := questionOptionMap[item.QuestionId]
		if ok {
			questionOptionMap[item.QuestionId] = append(questionOptionMap[item.QuestionId], resp.QuestionOption{
				Id:            item.Id,
				Title:         item.Title,
				Seq:           item.Seq,
				CheckedValues: nil,
			})
		} else {
			questionOptionMap[item.QuestionId] = []resp.QuestionOption{}
			questionOptionMap[item.QuestionId] = append(questionOptionMap[item.QuestionId], resp.QuestionOption{
				Id:            item.Id,
				Title:         item.Title,
				Seq:           item.Seq,
				CheckedValues: nil,
			})
		}
	}
	return questionOptionMap
}

func (c PagerController) List() {
	value := context.Background().Value("curr_user")
	fmt.Println(value)
	var papers []models.Paper
	core.GetOrm().Raw("select * from paper order by gmt_created desc").QueryRows(&papers)
	c.Success(papers)
}

func (c PagerController) Create() {
	paperReq := c.ParseFromJsonParam(req.PaperCreateReq{}).(req.PaperCreateReq)
	marshal := jsonUtils.ToJSONStr(paperReq)
	fmt.Println("创建问卷参数 -> ", marshal)
	err := core.GetOrm().DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		// step1:新增试卷
		paper := models.Paper{
			Title: paperReq.Title,
		}
		var err error
		err = createPaper(txOrm, &paper)
		if err != nil {
			return err
		}

		// step2:新增试卷题目
		for _, question := range paperReq.Questions {
			var insertQuestion, err = createQuestions(paper.Id, question, txOrm)
			if err != nil {
				return err
			}

			// step3:新增试卷题目的选项
			for index, option := range question.Options {
				err = createQuestionOptions(index, option, insertQuestion.Id, txOrm)
				if err != nil {
					return err
				}
			}

		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	c.Success()
}

func createQuestionOptions(index int, option req.QuestionOption, insertQuestionId int64, txOrm orm.TxOrmer) error {
	insertOption := models.QuestionOption{
		QuestionId: insertQuestionId,
		Title:      option.Title,
		Seq:        index + 1,
	}
	_, err := questionOptionDao.Insert(&insertOption, txOrm)
	if err != nil {
		logs.Error("新增试卷题目选项失败 -> %s", err)
		return err
	}
	return nil
}

func createQuestions(paperId int64, question req.Questions, txOrm orm.TxOrmer) (models.Question, error) {
	insertQuestion := models.Question{
		PaperId: paperId,
		Title:   question.Title,
		Type:    question.Type,
		Score:   5,
		Answer:  strconv.Itoa(question.Answer),
	}
	_, err := questionDao.Insert(insertQuestion, txOrm)
	if err != nil {
		logs.Error("新增试卷题目失败 -> %s", err)
		return models.Question{}, err
	}
	return insertQuestion, nil
}

func createPaper(txOrm orm.TxOrmer, paper *models.Paper) error {
	_, err := paperDao.Insert(paper, txOrm)
	if err != nil {
		logs.Error("新增试卷失败 -> %s", err)
		return err
	}
	return nil
}
