package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Req struct {
	CharterSectionId int    `json:"charterSectionId"`
	VideoId          int    `json:"videoId"`
	VideoType        int    `json:"videoType"`
	StudyTime        int    `json:"studyTime"`
	UploadType       int    `json:"uploadType"`
	LocalCreateTime  int64  `json:"localCreateTime"`
	AppType          int    `json:"appType"`
	LastStudyTime    int    `json:"lastStudyTime"`
	ApiType          string `json:"apiType"`
}

func main() {
	answer()

}

func answer() {
	answerStr := `[
  {
    "id": 59318140,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982799,
    "sequence": 1,
    "answer": "B",
    "status": 2,
    "score": 2,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "B",
    "topic": "判断理论的核心观点是(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490140,
        "questionId": 2982799,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "人幸福和痛苦是由其特质或者认知方式决定的",
        "containFormula": false
      },
      {
        "id": 8490141,
        "questionId": 2982799,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "幸福感是在“比较”中产生的",
        "containFormula": false
      },
      {
        "id": 8490142,
        "questionId": 2982799,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "态度包括认知成分、情感成分和行为成分",
        "containFormula": false
      },
      {
        "id": 8490143,
        "questionId": 2982799,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "社会会告诉个体，你应该幸福或者你应该痛苦，个体就会贴上快乐和痛苦的标签",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982799
      }
    ],
    "collected": false
  },
  {
    "id": 59318141,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982800,
    "sequence": 2,
    "answer": "B",
    "status": 3,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "C",
    "topic": "中国人的人际交往模式是(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490144,
        "questionId": 2982800,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "人际吸引",
        "containFormula": false
      },
      {
        "id": 8490145,
        "questionId": 2982800,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "交往程度",
        "containFormula": false
      },
      {
        "id": 8490146,
        "questionId": 2982800,
        "sequence": 3,
        "isRight": true,
        "answerKey": "C",
        "answerText": "差序格局",
        "containFormula": false
      },
      {
        "id": 8490147,
        "questionId": 2982800,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "契约式交往",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982800
      }
    ],
    "collected": false
  },
  {
    "id": 59318142,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982801,
    "sequence": 3,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "C",
    "topic": "下面不属于大学生群体重要的四类人际关系的是(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490148,
        "questionId": 2982801,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "家人关系",
        "containFormula": false
      },
      {
        "id": 8490149,
        "questionId": 2982801,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "恋人关系",
        "containFormula": false
      },
      {
        "id": 8490150,
        "questionId": 2982801,
        "sequence": 3,
        "isRight": true,
        "answerKey": "C",
        "answerText": "同事关系",
        "containFormula": false
      },
      {
        "id": 8490151,
        "questionId": 2982801,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "师生关系",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982801
      }
    ],
    "collected": false
  },
  {
    "id": 59318143,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982802,
    "sequence": 4,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "D",
    "topic": "‏心理学家根据宽恕主体和指向的不同，有三种宽恕行为。下列哪一个选项不属于宽恕行为。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490152,
        "questionId": 2982802,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "宽恕他人",
        "containFormula": false
      },
      {
        "id": 8490153,
        "questionId": 2982802,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "自我宽恕",
        "containFormula": false
      },
      {
        "id": 8490154,
        "questionId": 2982802,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "寻求宽恕",
        "containFormula": false
      },
      {
        "id": 8490155,
        "questionId": 2982802,
        "sequence": 4,
        "isRight": true,
        "answerKey": "D",
        "answerText": "报复行为",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982802
      }
    ],
    "collected": false
  },
  {
    "id": 59318144,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982803,
    "sequence": 5,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "B",
    "topic": "自尊概念最早是由哪位学者提出(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490156,
        "questionId": 2982803,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "罗森伯格",
        "containFormula": false
      },
      {
        "id": 8490157,
        "questionId": 2982803,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "詹姆斯",
        "containFormula": false
      },
      {
        "id": 8490158,
        "questionId": 2982803,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "维纳",
        "containFormula": false
      },
      {
        "id": 8490159,
        "questionId": 2982803,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "马斯洛",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982803
      }
    ],
    "collected": false
  },
  {
    "id": 59318145,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982804,
    "sequence": 6,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "B",
    "topic": "心理健康教育最早产生于美国，是在本世纪初(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490160,
        "questionId": 2982804,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "1890年",
        "containFormula": false
      },
      {
        "id": 8490161,
        "questionId": 2982804,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "1894年",
        "containFormula": false
      },
      {
        "id": 8490162,
        "questionId": 2982804,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "1898年",
        "containFormula": false
      },
      {
        "id": 8490163,
        "questionId": 2982804,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "1900年",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982804
      }
    ],
    "collected": false
  },
  {
    "id": 59318146,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982805,
    "sequence": 7,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "C",
    "topic": "心理健康教育课程活动的主要特点不包括(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490164,
        "questionId": 2982805,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "兴趣性",
        "containFormula": false
      },
      {
        "id": 8490165,
        "questionId": 2982805,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "启发性",
        "containFormula": false
      },
      {
        "id": 8490166,
        "questionId": 2982805,
        "sequence": 3,
        "isRight": true,
        "answerKey": "C",
        "answerText": "灵活性",
        "containFormula": false
      },
      {
        "id": 8490167,
        "questionId": 2982805,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "社会性",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982805
      }
    ],
    "collected": false
  },
  {
    "id": 59318147,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982806,
    "sequence": 8,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "A",
    "topic": "下面不属于积极情绪能给我们带来的益处的是(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490168,
        "questionId": 2982806,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "增加的压力感受",
        "containFormula": false
      },
      {
        "id": 8490169,
        "questionId": 2982806,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "更愿意正面地看待困境",
        "containFormula": false
      },
      {
        "id": 8490170,
        "questionId": 2982806,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "拓宽注意的范围",
        "containFormula": false
      },
      {
        "id": 8490171,
        "questionId": 2982806,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "增强免疫力",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982806
      }
    ],
    "collected": false
  },
  {
    "id": 59318148,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982807,
    "sequence": 9,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "C",
    "topic": "“一个人是否感到幸福,取决于他日常生活中幸福事件的多寡”是(   )的观点。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490172,
        "questionId": 2982807,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "认知决定论",
        "containFormula": false
      },
      {
        "id": 8490173,
        "questionId": 2982807,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "判断理论",
        "containFormula": false
      },
      {
        "id": 8490174,
        "questionId": 2982807,
        "sequence": 3,
        "isRight": true,
        "answerKey": "C",
        "answerText": "状态理论",
        "containFormula": false
      },
      {
        "id": 8490175,
        "questionId": 2982807,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "目标理论",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982807
      }
    ],
    "collected": false
  },
  {
    "id": 59318149,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982808,
    "sequence": 10,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "D",
    "topic": "美国心理学家(   )首先提出了自我接纳概念，认为它是自我意识的一个重要组成部分。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490176,
        "questionId": 2982808,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "皮亚杰",
        "containFormula": false
      },
      {
        "id": 8490177,
        "questionId": 2982808,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "罗杰斯",
        "containFormula": false
      },
      {
        "id": 8490178,
        "questionId": 2982808,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "埃里克森",
        "containFormula": false
      },
      {
        "id": 8490179,
        "questionId": 2982808,
        "sequence": 4,
        "isRight": true,
        "answerKey": "D",
        "answerText": "奥尔波特",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982808
      }
    ],
    "collected": false
  },
  {
    "id": 59318150,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982809,
    "sequence": 11,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "B",
    "topic": "“我必须成功”是哪种不合理的观念？",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490180,
        "questionId": 2982809,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "过分概括化",
        "containFormula": false
      },
      {
        "id": 8490181,
        "questionId": 2982809,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "绝对化的要求",
        "containFormula": false
      },
      {
        "id": 8490182,
        "questionId": 2982809,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "以偏概全",
        "containFormula": false
      },
      {
        "id": 8490183,
        "questionId": 2982809,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "糟糕至极",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982809
      }
    ],
    "collected": false
  },
  {
    "id": 59318151,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982810,
    "sequence": 12,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "D",
    "topic": "压力可以分为哪些类型(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490184,
        "questionId": 2982810,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "正性压力",
        "containFormula": false
      },
      {
        "id": 8490185,
        "questionId": 2982810,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "中性压力",
        "containFormula": false
      },
      {
        "id": 8490186,
        "questionId": 2982810,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "负性压力",
        "containFormula": false
      },
      {
        "id": 8490187,
        "questionId": 2982810,
        "sequence": 4,
        "isRight": true,
        "answerKey": "D",
        "answerText": "以上都是",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982810
      }
    ],
    "collected": false
  },
  {
    "id": 59318152,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982811,
    "sequence": 13,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "B",
    "topic": "大学生有哪些情绪特点(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490188,
        "questionId": 2982811,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "情绪体验表现得往往不具有社会性",
        "containFormula": false
      },
      {
        "id": 8490189,
        "questionId": 2982811,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "情绪具有冲动性与爆发性",
        "containFormula": false
      },
      {
        "id": 8490190,
        "questionId": 2982811,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "比较容易控制自己的情绪",
        "containFormula": false
      },
      {
        "id": 8490191,
        "questionId": 2982811,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "情绪比较稳定",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982811
      }
    ],
    "collected": false
  },
  {
    "id": 59318153,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982812,
    "sequence": 14,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "A",
    "topic": "学校心理健康教育的根本目的在于(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490192,
        "questionId": 2982812,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "培养学生良好的心理素质",
        "containFormula": false
      },
      {
        "id": 8490193,
        "questionId": 2982812,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "医治和矫正学生的心理异常",
        "containFormula": false
      },
      {
        "id": 8490194,
        "questionId": 2982812,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "解决人的心理发展问题",
        "containFormula": false
      },
      {
        "id": 8490195,
        "questionId": 2982812,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "实现学生身心和谐发展",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982812
      }
    ],
    "collected": false
  },
  {
    "id": 59318154,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982813,
    "sequence": 15,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "B",
    "topic": "‎心理健康的实质是对个体生命活动的有效管理。这种管理的有效性能否实现以及能在多大程度上实现，主要由个体(   )的发挥水平决定。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490196,
        "questionId": 2982813,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "生活性",
        "containFormula": false
      },
      {
        "id": 8490197,
        "questionId": 2982813,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "主体性",
        "containFormula": false
      },
      {
        "id": 8490198,
        "questionId": 2982813,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "积极性",
        "containFormula": false
      },
      {
        "id": 8490199,
        "questionId": 2982813,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "全面性",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982813
      }
    ],
    "collected": false
  },
  {
    "id": 59318155,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982814,
    "sequence": 16,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "B",
    "topic": "个体由于自身人际关系紧张而产生焦虑不安的情绪，这属于(   )",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490200,
        "questionId": 2982814,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "自我认识",
        "containFormula": false
      },
      {
        "id": 8490201,
        "questionId": 2982814,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "自我体验",
        "containFormula": false
      },
      {
        "id": 8490202,
        "questionId": 2982814,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "自我控制",
        "containFormula": false
      },
      {
        "id": 8490203,
        "questionId": 2982814,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "自我觉知",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982814
      }
    ],
    "collected": false
  },
  {
    "id": 59318156,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982815,
    "sequence": 17,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "C",
    "topic": "下关于自我概念的说法错误的是(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490204,
        "questionId": 2982815,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "米德认为自我分为主体我和客体我",
        "containFormula": false
      },
      {
        "id": 8490205,
        "questionId": 2982815,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "对自我的概括性评价",
        "containFormula": false
      },
      {
        "id": 8490206,
        "questionId": 2982815,
        "sequence": 3,
        "isRight": true,
        "answerKey": "C",
        "answerText": "库利认为他人评价与自我的形成无关",
        "containFormula": false
      },
      {
        "id": 8490207,
        "questionId": 2982815,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "弗洛伊德认为自我包括本我、自我和超我",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982815
      }
    ],
    "collected": false
  },
  {
    "id": 59318157,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982816,
    "sequence": 18,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "C",
    "topic": "列表现与自我同一性建立的人不相符的是(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490208,
        "questionId": 2982816,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "清楚地知道我是谁",
        "containFormula": false
      },
      {
        "id": 8490209,
        "questionId": 2982816,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "有比较明确的目标",
        "containFormula": false
      },
      {
        "id": 8490210,
        "questionId": 2982816,
        "sequence": 3,
        "isRight": true,
        "answerKey": "C",
        "answerText": "对自己是什么样的人感到迷惑",
        "containFormula": false
      },
      {
        "id": 8490211,
        "questionId": 2982816,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "了解自己追求什么样的生活",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982816
      }
    ],
    "collected": false
  },
  {
    "id": 59318158,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982817,
    "sequence": 19,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "D",
    "topic": "以下关于自我同一性建立的优点的描述不正确的是(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490212,
        "questionId": 2982817,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "有利于与父母建立良好的关系",
        "containFormula": false
      },
      {
        "id": 8490213,
        "questionId": 2982817,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "有利于个体自身自尊的提高",
        "containFormula": false
      },
      {
        "id": 8490214,
        "questionId": 2982817,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "让我们对大学的生活更满意",
        "containFormula": false
      },
      {
        "id": 8490215,
        "questionId": 2982817,
        "sequence": 4,
        "isRight": true,
        "answerKey": "D",
        "answerText": "对提升我们在新环境中的适应性没有太大帮助",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982817
      }
    ],
    "collected": false
  },
  {
    "id": 59318159,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982818,
    "sequence": 20,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "C",
    "topic": "以下不属于自恋型人格障碍的人的特点的是(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490216,
        "questionId": 2982818,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "无法共情，经常会撒谎",
        "containFormula": false
      },
      {
        "id": 8490217,
        "questionId": 2982818,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "夸大自己",
        "containFormula": false
      },
      {
        "id": 8490218,
        "questionId": 2982818,
        "sequence": 3,
        "isRight": true,
        "answerKey": "C",
        "answerText": "适当考虑他人的需求",
        "containFormula": false
      },
      {
        "id": 8490219,
        "questionId": 2982818,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "自负、傲慢",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982818
      }
    ],
    "collected": false
  },
  {
    "id": 59318160,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982819,
    "sequence": 21,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "A",
    "topic": "下列不属于自恋形成的主要心理层面的原因的是(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490220,
        "questionId": 2982819,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "自身的经济外貌条件十分优越",
        "containFormula": false
      },
      {
        "id": 8490221,
        "questionId": 2982819,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "害怕自己一无是处，自卑",
        "containFormula": false
      },
      {
        "id": 8490222,
        "questionId": 2982819,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "容易对比自己优秀的人产生嫉妒感",
        "containFormula": false
      },
      {
        "id": 8490223,
        "questionId": 2982819,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "缺乏安全感",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982819
      }
    ],
    "collected": false
  },
  {
    "id": 59318161,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982820,
    "sequence": 22,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "A",
    "topic": "以下关于自尊的描述正确的是(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490224,
        "questionId": 2982820,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "个体遇到挫折时自尊会降低",
        "containFormula": false
      },
      {
        "id": 8490225,
        "questionId": 2982820,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "被他人接纳和积极评价会使自尊先上升再下降",
        "containFormula": false
      },
      {
        "id": 8490226,
        "questionId": 2982820,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "在高焦虑状态下，外部给与的负面评价对个体自尊的影响较小",
        "containFormula": false
      },
      {
        "id": 8490227,
        "questionId": 2982820,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "外部的负面评价会使自尊先降低再提升",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982820
      }
    ],
    "collected": false
  },
  {
    "id": 59318162,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982821,
    "sequence": 23,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "B",
    "topic": "以下关于拖延描述错误的是(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490228,
        "questionId": 2982821,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "从心理层面上接纳拖延有利于缓解拖延",
        "containFormula": false
      },
      {
        "id": 8490229,
        "questionId": 2982821,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "拖延是一种疾病",
        "containFormula": false
      },
      {
        "id": 8490230,
        "questionId": 2982821,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "战胜拖延是为了让我们更好地生活，生活不是为了战胜拖延",
        "containFormula": false
      },
      {
        "id": 8490231,
        "questionId": 2982821,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "拖延描述的是行为而不是人",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982821
      }
    ],
    "collected": false
  },
  {
    "id": 59318163,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982822,
    "sequence": 24,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "C",
    "topic": "下列关于人际交往的说法错误的是(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490232,
        "questionId": 2982822,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "大学的人际交往能为毕业以后参加工作做好准备",
        "containFormula": false
      },
      {
        "id": 8490233,
        "questionId": 2982822,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "人际交往能够提供情感支持",
        "containFormula": false
      },
      {
        "id": 8490234,
        "questionId": 2982822,
        "sequence": 3,
        "isRight": true,
        "answerKey": "C",
        "answerText": "人际交往会激发竞争，阻碍合作",
        "containFormula": false
      },
      {
        "id": 8490235,
        "questionId": 2982822,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "人际交往有利于自我完善",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982822
      }
    ],
    "collected": false
  },
  {
    "id": 59318164,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982823,
    "sequence": 25,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "A",
    "topic": "‎有同理心的人，不具有下面什么特征",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490236,
        "questionId": 2982823,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "比较内向",
        "containFormula": false
      },
      {
        "id": 8490237,
        "questionId": 2982823,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "善于倾听",
        "containFormula": false
      },
      {
        "id": 8490238,
        "questionId": 2982823,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "能够更好做到换位思考",
        "containFormula": false
      },
      {
        "id": 8490239,
        "questionId": 2982823,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "尊重他人",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982823
      }
    ],
    "collected": false
  },
  {
    "id": 59318165,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982824,
    "sequence": 26,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "A",
    "topic": "‏要具有一颗同理心，下面做法是错误的是(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490240,
        "questionId": 2982824,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "放弃自己的价值体系，尽量迎合他人",
        "containFormula": false
      },
      {
        "id": 8490241,
        "questionId": 2982824,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "尽量真诚待人",
        "containFormula": false
      },
      {
        "id": 8490242,
        "questionId": 2982824,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "尊重他人",
        "containFormula": false
      },
      {
        "id": 8490243,
        "questionId": 2982824,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "具有一颗好奇心",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982824
      }
    ],
    "collected": false
  },
  {
    "id": 59318166,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982825,
    "sequence": 27,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "B",
    "topic": "在意向上表现为追求智慧、能力的发展和追求理想、信仰的自我意识属于(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490244,
        "questionId": 2982825,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "生理自我",
        "containFormula": false
      },
      {
        "id": 8490245,
        "questionId": 2982825,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "心理自我",
        "containFormula": false
      },
      {
        "id": 8490246,
        "questionId": 2982825,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "物质自我",
        "containFormula": false
      },
      {
        "id": 8490247,
        "questionId": 2982825,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "社会自我",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982825
      }
    ],
    "collected": false
  },
  {
    "id": 59318167,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982826,
    "sequence": 28,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "C",
    "topic": "心理学家鲁夫特与英格汉提出“周哈里窗（JohariWindow）”模式，“窗”是指一个人的心就像一扇窗，周哈里窗展示了关于自我认知、行为举止和他人对自己的认知之间在有意识或无意识的前提下形成的差异。其中“你知我不知”的是指(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490248,
        "questionId": 2982826,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "开放我",
        "containFormula": false
      },
      {
        "id": 8490249,
        "questionId": 2982826,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "隐藏我",
        "containFormula": false
      },
      {
        "id": 8490250,
        "questionId": 2982826,
        "sequence": 3,
        "isRight": true,
        "answerKey": "C",
        "answerText": "背脊我",
        "containFormula": false
      },
      {
        "id": 8490251,
        "questionId": 2982826,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "未知我",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982826
      }
    ],
    "collected": false
  },
  {
    "id": 59318168,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982827,
    "sequence": 29,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "A",
    "topic": "(   )以最近3年的大型心理健康调查为基础，探讨我国不同人群的心理健康现状，发展趋势及影响因素。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490252,
        "questionId": 2982827,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "《中国国民心理健康发展报告（2017~2018）》",
        "containFormula": false
      },
      {
        "id": 8490253,
        "questionId": 2982827,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "《心理健康白皮书》",
        "containFormula": false
      },
      {
        "id": 8490254,
        "questionId": 2982827,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "《中国心理健康状况调查与分析报告（2017~2018）》",
        "containFormula": false
      },
      {
        "id": 8490255,
        "questionId": 2982827,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "《中国心理健康报告》",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982827
      }
    ],
    "collected": false
  },
  {
    "id": 59318169,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211276,
    "questionId": 2982828,
    "sequence": 30,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "单选题",
    "rightAnswer": "D",
    "topic": "面对可能出现情绪或精神方面的“精神感冒”，学校心理教育工作应注重(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490256,
        "questionId": 2982828,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "疏导",
        "containFormula": false
      },
      {
        "id": 8490257,
        "questionId": 2982828,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "指导",
        "containFormula": false
      },
      {
        "id": 8490258,
        "questionId": 2982828,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "治疗",
        "containFormula": false
      },
      {
        "id": 8490259,
        "questionId": 2982828,
        "sequence": 4,
        "isRight": true,
        "answerKey": "D",
        "answerText": "预防",
        "containFormula": false
      }
    ],
    "questionType": 1,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982828
      }
    ],
    "collected": false
  },
  {
    "id": 59318170,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211277,
    "questionId": 2982829,
    "sequence": 31,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "多选题",
    "rightAnswer": "AD",
    "topic": "下列选项中，回答错误的是(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490260,
        "questionId": 2982829,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "在我的二十个自我描述中，如果消极的自我描述多，积极的自我描述少，说明我这个人缺点比优点多",
        "containFormula": false
      },
      {
        "id": 8490261,
        "questionId": 2982829,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "客观认识自己对健康成长有重要意义",
        "containFormula": false
      },
      {
        "id": 8490262,
        "questionId": 2982829,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "我们可以通过与别人比较、与自己比较、心理测试等方法来客观认识自己",
        "containFormula": false
      },
      {
        "id": 8490263,
        "questionId": 2982829,
        "sequence": 4,
        "isRight": true,
        "answerKey": "D",
        "answerText": "人的价值由他的学习成绩、家庭经济地位以及衣着相貌所决定",
        "containFormula": false
      }
    ],
    "questionType": 2,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982829
      }
    ],
    "collected": false
  },
  {
    "id": 59318171,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211277,
    "questionId": 2982830,
    "sequence": 32,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "多选题",
    "rightAnswer": "AD",
    "topic": "下列选项中，回答正确的是(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490264,
        "questionId": 2982830,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "要想让梦想变得触手可及，你可以设定实现梦想的目标，如长期目标、中期目标和短期目标",
        "containFormula": false
      },
      {
        "id": 8490265,
        "questionId": 2982830,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "如果你在网上做了一个关于人际关系的心理测试，发现自己的人际关系得分较低，由此可以推断你患上了人际关系恐怖症",
        "containFormula": false
      },
      {
        "id": 8490266,
        "questionId": 2982830,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "为了提升学习效率，制订计划只需要将学习时间安排紧凑即可，睡眠、运动等其他活动不需要考虑",
        "containFormula": false
      },
      {
        "id": 8490267,
        "questionId": 2982830,
        "sequence": 4,
        "isRight": true,
        "answerKey": "D",
        "answerText": "要根据自己的实际情况来制订计划",
        "containFormula": false
      }
    ],
    "questionType": 2,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982830
      }
    ],
    "collected": false
  },
  {
    "id": 59318172,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211277,
    "questionId": 2982831,
    "sequence": 33,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "多选题",
    "rightAnswer": "ABD",
    "topic": "在生活中我们会体验到挫折感，是因为自己有一些期待没有得到满足。关于如何满足自己的期待，下列说法正确的是(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490268,
        "questionId": 2982831,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "通过努力去满足期待。如，小娟通过努力，提升自己酒店服务的能力",
        "containFormula": false
      },
      {
        "id": 8490269,
        "questionId": 2982831,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "调整努力的方向，找到替代之法",
        "containFormula": false
      },
      {
        "id": 8490270,
        "questionId": 2982831,
        "sequence": 3,
        "isRight": false,
        "answerKey": "C",
        "answerText": "以上说法都不正确",
        "containFormula": false
      },
      {
        "id": 8490271,
        "questionId": 2982831,
        "sequence": 4,
        "isRight": true,
        "answerKey": "D",
        "answerText": "对于一些无法改变的现实，我们可以选择接纳，放下期待。如，小凯奶奶已经去世，他无法实现带奶奶旅游的承诺，就要面对现实",
        "containFormula": false
      }
    ],
    "questionType": 2,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982831
      }
    ],
    "collected": false
  },
  {
    "id": 59318173,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211277,
    "questionId": 2982832,
    "sequence": 34,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "多选题",
    "rightAnswer": "AC",
    "topic": "自我控制体现了人的自我控制和人的能动性，包括了以下(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490272,
        "questionId": 2982832,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "自律",
        "containFormula": false
      },
      {
        "id": 8490273,
        "questionId": 2982832,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "自信",
        "containFormula": false
      },
      {
        "id": 8490274,
        "questionId": 2982832,
        "sequence": 3,
        "isRight": true,
        "answerKey": "C",
        "answerText": "最求名誉地位，获得他人好感。",
        "containFormula": false
      },
      {
        "id": 8490275,
        "questionId": 2982832,
        "sequence": 4,
        "isRight": false,
        "answerKey": "D",
        "answerText": "自卑",
        "containFormula": false
      }
    ],
    "questionType": 2,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982832
      }
    ],
    "collected": false
  },
  {
    "id": 59318174,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211277,
    "questionId": 2982833,
    "sequence": 35,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "多选题",
    "rightAnswer": "ABCDE",
    "topic": "下列选项中属于积极体验对人生意义的有(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490276,
        "questionId": 2982833,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "积极体验使我们对新思想和新活动保持开放的心态",
        "containFormula": false
      },
      {
        "id": 8490277,
        "questionId": 2982833,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "积极体验有助于纠正消极体验形成的消极认识",
        "containFormula": false
      },
      {
        "id": 8490278,
        "questionId": 2982833,
        "sequence": 3,
        "isRight": true,
        "answerKey": "C",
        "answerText": "积极体验能够提高主观幸福感",
        "containFormula": false
      },
      {
        "id": 8490279,
        "questionId": 2982833,
        "sequence": 4,
        "isRight": true,
        "answerKey": "D",
        "answerText": "积极体验对学习、健康等有积极的促进作用",
        "containFormula": false
      },
      {
        "id": 8490280,
        "questionId": 2982833,
        "sequence": 5,
        "isRight": true,
        "answerKey": "E",
        "answerText": "积极体验对健康有促进作用",
        "containFormula": false
      }
    ],
    "questionType": 2,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982833
      }
    ],
    "collected": false
  },
  {
    "id": 59318175,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211277,
    "questionId": 2982834,
    "sequence": 36,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "多选题",
    "rightAnswer": "ACD",
    "topic": "影响自尊水平的因素有哪些(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490281,
        "questionId": 2982834,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "调节不良状态",
        "containFormula": false
      },
      {
        "id": 8490282,
        "questionId": 2982834,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "控制锻炼时间",
        "containFormula": false
      },
      {
        "id": 8490283,
        "questionId": 2982834,
        "sequence": 3,
        "isRight": true,
        "answerKey": "C",
        "answerText": "控制锻炼强度",
        "containFormula": false
      },
      {
        "id": 8490284,
        "questionId": 2982834,
        "sequence": 4,
        "isRight": true,
        "answerKey": "D",
        "answerText": "获得成功体验",
        "containFormula": false
      }
    ],
    "questionType": 2,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982834
      }
    ],
    "collected": false
  },
  {
    "id": 59318176,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211277,
    "questionId": 2982835,
    "sequence": 37,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "多选题",
    "rightAnswer": "ABCD",
    "topic": "心理健康教育课程活动的主要形式有(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490285,
        "questionId": 2982835,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "角色活动",
        "containFormula": false
      },
      {
        "id": 8490286,
        "questionId": 2982835,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "表演活动",
        "containFormula": false
      },
      {
        "id": 8490287,
        "questionId": 2982835,
        "sequence": 3,
        "isRight": true,
        "answerKey": "C",
        "answerText": "游戏活动",
        "containFormula": false
      },
      {
        "id": 8490288,
        "questionId": 2982835,
        "sequence": 4,
        "isRight": true,
        "answerKey": "D",
        "answerText": "体育活动",
        "containFormula": false
      }
    ],
    "questionType": 2,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982835
      }
    ],
    "collected": false
  },
  {
    "id": 59318177,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211277,
    "questionId": 2982836,
    "sequence": 38,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "多选题",
    "rightAnswer": "ABCD",
    "topic": "心理健康的重要意义有(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490289,
        "questionId": 2982836,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "有利于帮助与促进大学生善爱我，做最好的自己",
        "containFormula": false
      },
      {
        "id": 8490290,
        "questionId": 2982836,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "有利于帮助与促进大学生更好的进行情绪管控",
        "containFormula": false
      },
      {
        "id": 8490291,
        "questionId": 2982836,
        "sequence": 3,
        "isRight": true,
        "answerKey": "C",
        "answerText": "有利于帮助与促进大学生更好地经营人际关系",
        "containFormula": false
      },
      {
        "id": 8490292,
        "questionId": 2982836,
        "sequence": 4,
        "isRight": true,
        "answerKey": "D",
        "answerText": "有利于帮助与促进大学生培育健康的人格",
        "containFormula": false
      }
    ],
    "questionType": 2,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982836
      }
    ],
    "collected": false
  },
  {
    "id": 59318178,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211277,
    "questionId": 2982837,
    "sequence": 39,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "多选题",
    "rightAnswer": "BCD",
    "topic": "詹姆斯(W.James)将自我意识划分为三种形式，它们是(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490293,
        "questionId": 2982837,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "现实自我",
        "containFormula": false
      },
      {
        "id": 8490294,
        "questionId": 2982837,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "物质自我",
        "containFormula": false
      },
      {
        "id": 8490295,
        "questionId": 2982837,
        "sequence": 3,
        "isRight": true,
        "answerKey": "C",
        "answerText": "心理自我",
        "containFormula": false
      },
      {
        "id": 8490296,
        "questionId": 2982837,
        "sequence": 4,
        "isRight": true,
        "answerKey": "D",
        "answerText": "社会自我",
        "containFormula": false
      }
    ],
    "questionType": 2,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982837
      }
    ],
    "collected": false
  },
  {
    "id": 59318179,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211277,
    "questionId": 2982838,
    "sequence": 40,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "多选题",
    "rightAnswer": "ABCDE",
    "topic": "关于幸福的理论有(   )。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490297,
        "questionId": 2982838,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "认知决定论",
        "containFormula": false
      },
      {
        "id": 8490298,
        "questionId": 2982838,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "判断理论",
        "containFormula": false
      },
      {
        "id": 8490299,
        "questionId": 2982838,
        "sequence": 3,
        "isRight": true,
        "answerKey": "C",
        "answerText": "认知协调论",
        "containFormula": false
      },
      {
        "id": 8490300,
        "questionId": 2982838,
        "sequence": 4,
        "isRight": true,
        "answerKey": "D",
        "answerText": "目标理论",
        "containFormula": false
      },
      {
        "id": 8490301,
        "questionId": 2982838,
        "sequence": 5,
        "isRight": true,
        "answerKey": "E",
        "answerText": "活动理论",
        "containFormula": false
      }
    ],
    "questionType": 2,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982838
      }
    ],
    "collected": false
  },
  {
    "id": 59318180,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211278,
    "questionId": 2982839,
    "sequence": 41,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "判断题",
    "rightAnswer": "B",
    "topic": "积极情绪就是正性情绪；消极情绪就是负性情绪。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490302,
        "questionId": 2982839,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "对",
        "containFormula": false
      },
      {
        "id": 8490303,
        "questionId": 2982839,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "错",
        "containFormula": false
      }
    ],
    "questionType": 4,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982839
      }
    ],
    "collected": false
  },
  {
    "id": 59318181,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211278,
    "questionId": 2982840,
    "sequence": 42,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "判断题",
    "rightAnswer": "A",
    "topic": "情绪的生理唤醒有共同模式，也有个体差异。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490304,
        "questionId": 2982840,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "对",
        "containFormula": false
      },
      {
        "id": 8490305,
        "questionId": 2982840,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "错",
        "containFormula": false
      }
    ],
    "questionType": 4,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982840
      }
    ],
    "collected": false
  },
  {
    "id": 59318182,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211278,
    "questionId": 2982841,
    "sequence": 43,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "判断题",
    "rightAnswer": "B",
    "topic": "弗里杰达曾提出情绪是由六个相关联的成分组成。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490306,
        "questionId": 2982841,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "对",
        "containFormula": false
      },
      {
        "id": 8490307,
        "questionId": 2982841,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "错",
        "containFormula": false
      }
    ],
    "questionType": 4,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982841
      }
    ],
    "collected": false
  },
  {
    "id": 59318183,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211278,
    "questionId": 2982842,
    "sequence": 44,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "判断题",
    "rightAnswer": "A",
    "topic": "接触大自然，获得美的感受与心灵的愉悦有助于增进积极体验。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490308,
        "questionId": 2982842,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "对",
        "containFormula": false
      },
      {
        "id": 8490309,
        "questionId": 2982842,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "错",
        "containFormula": false
      }
    ],
    "questionType": 4,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982842
      }
    ],
    "collected": false
  },
  {
    "id": 59318184,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211278,
    "questionId": 2982843,
    "sequence": 45,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "判断题",
    "rightAnswer": "B",
    "topic": "中国人的人际交往重感情，讲血缘。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490310,
        "questionId": 2982843,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "对",
        "containFormula": false
      },
      {
        "id": 8490311,
        "questionId": 2982843,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "错",
        "containFormula": false
      }
    ],
    "questionType": 4,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982843
      }
    ],
    "collected": false
  },
  {
    "id": 59318185,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211278,
    "questionId": 2982844,
    "sequence": 46,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "判断题",
    "rightAnswer": "A",
    "topic": "人际交往及其关系反映了一个人的社会性。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490312,
        "questionId": 2982844,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "对",
        "containFormula": false
      },
      {
        "id": 8490313,
        "questionId": 2982844,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "错",
        "containFormula": false
      }
    ],
    "questionType": 4,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982844
      }
    ],
    "collected": false
  },
  {
    "id": 59318186,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211278,
    "questionId": 2982845,
    "sequence": 47,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "判断题",
    "rightAnswer": "B",
    "topic": "独立性自尊是指依赖他人肯定和表扬而产生的自尊。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490314,
        "questionId": 2982845,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "对",
        "containFormula": false
      },
      {
        "id": 8490315,
        "questionId": 2982845,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "错",
        "containFormula": false
      }
    ],
    "questionType": 4,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982845
      }
    ],
    "collected": false
  },
  {
    "id": 59318187,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211278,
    "questionId": 2982846,
    "sequence": 48,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "判断题",
    "rightAnswer": "B",
    "topic": "“自尊=成功/抱负水平”是由罗森伯格提出的。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490316,
        "questionId": 2982846,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "对",
        "containFormula": false
      },
      {
        "id": 8490317,
        "questionId": 2982846,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "错",
        "containFormula": false
      }
    ],
    "questionType": 4,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982846
      }
    ],
    "collected": false
  },
  {
    "id": 59318188,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211278,
    "questionId": 2982847,
    "sequence": 49,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "判断题",
    "rightAnswer": "B",
    "topic": "所有的研究都表明体育锻炼能够积极改善自尊水平。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490318,
        "questionId": 2982847,
        "sequence": 1,
        "isRight": false,
        "answerKey": "A",
        "answerText": "对",
        "containFormula": false
      },
      {
        "id": 8490319,
        "questionId": 2982847,
        "sequence": 2,
        "isRight": true,
        "answerKey": "B",
        "answerText": "错",
        "containFormula": false
      }
    ],
    "questionType": 4,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982847
      }
    ],
    "collected": false
  },
  {
    "id": 59318189,
    "examId": 76706,
    "paperId": 66186,
    "groupId": 211278,
    "questionId": 2982848,
    "sequence": 50,
    "status": 1,
    "score": 0,
    "fullScore": 2,
    "groupName": "判断题",
    "rightAnswer": "A",
    "topic": "积极情绪就是一种带来愉悦体验的情绪。",
    "paperName": "课后练习",
    "answerItemList": [
      {
        "id": 8490320,
        "questionId": 2982848,
        "sequence": 1,
        "isRight": true,
        "answerKey": "A",
        "answerText": "对",
        "containFormula": false
      },
      {
        "id": 8490321,
        "questionId": 2982848,
        "sequence": 2,
        "isRight": false,
        "answerKey": "B",
        "answerText": "错",
        "containFormula": false
      }
    ],
    "questionType": 4,
    "containFormula": false,
    "subjectId": -1,
    "parentQuestionId": -1,
    "knowledgeSystemList": [
      {
        "id": 188975,
        "parentId": 0,
        "typeName": "课后练习",
        "idPath": "/188975/",
        "namePath": "/课后练习/",
        "typeSeq": 7647,
        "questionId": 2982848
      }
    ],
    "collected": false
  }
]`
	//checkVersionChange := map[string]interface{}{
	//	"studentTaskId": 1090833,
	//	"version":       7,
	//}
	//reqStr, _ := json.Marshal(checkVersionChange)
	//client := http.Client{}
	//req, err := http.NewRequest("POST", "https://www.kepeiol.com/student/web/student/task/checkVersionChange", bytes.NewReader(reqStr))
	//if err != nil {
	//	panic(err)
	//}
	//for key, value := range GetHeader() {
	//	req.Header.Set(key, value)
	//}
	//res, err := client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer res.Body.Close()
	//body, _ := io.ReadAll(res.Body)
	//fmt.Println("checkVersionChange:" + string(body))
	//
	//start := map[string]interface{}{
	//	"id":        "76666",
	//	"subjectId": "",
	//	"taskId":    "27344",
	//}
	//reqStr, _ = json.Marshal(start)
	//client = http.Client{}
	//req, err = http.NewRequest("POST", "https://www.kepeiol.com/student/web/examPaper/start", bytes.NewReader(reqStr))
	//if err != nil {
	//	panic(err)
	//}
	//for key, value := range GetHeader() {
	//	req.Header.Set(key, value)
	//}
	//res, err = client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer res.Body.Close()
	//body, _ = io.ReadAll(res.Body)
	//fmt.Println("start:" + string(body))

	answers := []Answer{}
	_ = json.Unmarshal([]byte(answerStr), &answers)
	answerReq := AnswerReq{
		QuestionAnswers: []QuestionAnswer{},
		TestUserId:      10575274,
		TaskId:          "27344",
		Version:         9,
	}
	for _, v := range answers {
		//time.Sleep(200 * time.Millisecond)
		//Update(v)
		if v.QuestionType == 13 {
			rightAnswer := strings.Split(v.RightAnswer, ".")
			answerReq.QuestionAnswers = append(answerReq.QuestionAnswers, QuestionAnswer{
				UserAnswerId: v.Id,
				QuestionId:   v.QuestionId,
				QuestionType: v.QuestionType,
				Answer:       []string{rightAnswer[1]},
				Sequence:     v.Sequence,
			})
			fmt.Println(v.Sequence, v.RightAnswer)
		} else if v.QuestionType == 2 {
			//多选
			answerReq.QuestionAnswers = append(answerReq.QuestionAnswers, QuestionAnswer{
				UserAnswerId: v.Id,
				QuestionId:   v.QuestionId,
				QuestionType: v.QuestionType,
				Answer:       strings.Split(v.RightAnswer, ""),
				Sequence:     v.Sequence,
			})
			fmt.Println(v.Sequence, v.RightAnswer)
		} else {
			//单选
			answerReq.QuestionAnswers = append(answerReq.QuestionAnswers, QuestionAnswer{
				UserAnswerId: v.Id,
				QuestionId:   v.QuestionId,
				QuestionType: v.QuestionType,
				Answer:       []string{v.RightAnswer},
				Sequence:     v.Sequence,
			})
			fmt.Println(v.Sequence, v.RightAnswer)
		}

	}
	//reqStr, _ := json.Marshal(answerReq)
	//
	//client := http.Client{}
	//req, err := http.NewRequest("POST", "https://www.kepeiol.com/student/web/examPaper/submit", bytes.NewReader(reqStr))
	//if err != nil {
	//	panic(err)
	//}
	//for key, value := range GetHeader() {
	//	req.Header.Set(key, value)
	//}
	//res, err := client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer res.Body.Close()
	//body, _ := io.ReadAll(res.Body)
	//fmt.Println("body:" + string(body))
	//
	//question := map[string]interface{}{
	//	"examUserId":   cast.ToString(answerReq.TestUserId),
	//	"questionType": nil,
	//	"status":       nil,
	//	"subjectId":    nil,
	//}
	//reqStr, _ = json.Marshal(question)
	//client = http.Client{}
	//req, err = http.NewRequest("POST", "https://www.kepeiol.com/student/web/examPaper/question", bytes.NewReader(reqStr))
	//if err != nil {
	//	panic(err)
	//}
	//for key, value := range GetHeader() {
	//	req.Header.Set(key, value)
	//}
	//res, err = client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer res.Body.Close()
	//body, _ = io.ReadAll(res.Body)
	//fmt.Println("body1:" + string(body))
	//
	//taskSubmitInfo := TaskSubmitInfo{
	//	ExamUserId: cast.ToString(answerReq.TestUserId),
	//	TaskId:     answerReq.TaskId,
	//}
	//reqStr, _ = json.Marshal(taskSubmitInfo)
	//client = http.Client{}
	//req, err = http.NewRequest("POST", "https://www.kepeiol.com/student/web/examPaper/taskSubmitInfo", bytes.NewReader(reqStr))
	//if err != nil {
	//	panic(err)
	//}
	//for key, value := range GetHeader() {
	//	req.Header.Set(key, value)
	//}
	//res, err = client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer res.Body.Close()
	//body, _ = io.ReadAll(res.Body)
	//fmt.Println("body1:" + string(body))
	//report := Report{
	//	Id: cast.ToString(answerReq.TestUserId),
	//}
	//reqStr, _ = json.Marshal(report)
	//client = http.Client{}
	//req, err = http.NewRequest("POST", "https://www.kepeiol.com/student/web/examPaper/report", bytes.NewReader(reqStr))
	//if err != nil {
	//	panic(err)
	//}
	//for key, value := range GetHeader() {
	//	req.Header.Set(key, value)
	//}
	//res, err = client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer res.Body.Close()
	//body, _ = io.ReadAll(res.Body)
	//fmt.Println("body2:" + string(body))
	//info := map[string]interface{}{
	//	"taskId": 27344,
	//}
	//reqStr, _ = json.Marshal(info)
	//client = http.Client{}
	//req, err = http.NewRequest("POST", "https://www.kepeiol.com/student/web/student/task/info", bytes.NewReader(reqStr))
	//if err != nil {
	//	panic(err)
	//}
	//for key, value := range GetHeader() {
	//	req.Header.Set(key, value)
	//}
	//res, err = client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer res.Body.Close()
	//body, _ = io.ReadAll(res.Body)
	//fmt.Println("body2:" + string(body))
}

type Report struct {
	Id string `json:"id"`
}

func Update(answer Answer) {
	updateReq := UpdateReq{
		UserAnswerId: answer.Id,
		Answer:       []string{answer.RightAnswer},
		Sequence:     answer.Sequence,
		QuestionId:   answer.QuestionId,
		QuestionType: answer.QuestionType,
		TaskId:       nil,
	}
	reqStr, _ := json.Marshal(updateReq)
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://www.kepeiol.com/student/web/examPaper/update", bytes.NewReader(reqStr))
	if err != nil {
		panic(err)
	}
	for key, value := range GetHeader() {
		req.Header.Set(key, value)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println("Update:" + string(body))
	return
}

type UpdateReq struct {
	UserAnswerId int         `json:"userAnswerId"`
	Answer       []string    `json:"answer"`
	Sequence     int         `json:"sequence"`
	QuestionId   int         `json:"questionId"`
	QuestionType int         `json:"questionType"`
	TaskId       interface{} `json:"taskId"`
}

type TaskSubmitInfo struct {
	ExamUserId string `json:"examUserId"`
	TaskId     string `json:"taskId"`
}

type Answer struct {
	Id           int    `json:"id"`
	QuestionId   int    `json:"questionId"`
	RightAnswer  string `json:"rightAnswer"`
	QuestionType int    `json:"questionType"`
	Sequence     int    `json:"sequence"`
}

type AnswerReq struct {
	QuestionAnswers []QuestionAnswer `json:"questionAnswers"`
	TestUserId      int              `json:"testUserId"`
	TaskId          string           `json:"taskId"`
	Version         int              `json:"version"`
}
type QuestionAnswer struct {
	UserAnswerId int      `json:"userAnswerId"`
	QuestionId   int      `json:"questionId"`
	QuestionType int      `json:"questionType"`
	Answer       []string `json:"answer"`
	Sequence     int      `json:"sequence"`
}

func video() {
	//videoListStr := ``

	//var videoList []Video
	//_ = json.Unmarshal([]byte(videoListStr), &videoList)
	//for _, v := range videoList {
	for i := 0; i < 300; i++ {
		addVideoTime(72975, 108392, 99)
		//fmt.Println("id:", 108243)
		time.Sleep(200 * time.Millisecond)

	}
	//}
}

func addVideoTime(charterSectionId int, videoId int, studyTime int) {
	bodyData := AddVideoTimeBodyData{
		//CharterSectionId: charterSectionId,
		VideoId: videoId,
		//VideoType:       1,
		StudyTime:       studyTime,
		UploadType:      1,
		LocalCreateTime: time.Now().UnixNano() / int64(time.Millisecond) * 1000,
		AppType:         3,
		LastStudyTime:   1000,
		ApiType:         "api_Connect-StudyVideo.js-1",
	}
	reqStr, _ := json.Marshal(bodyData)

	client := http.Client{}
	req, err := http.NewRequest("POST", "https://www.kepeiol.com/yunkai/admin/userstudyrecord/addVideoTime", bytes.NewReader(reqStr))
	if err != nil {
		panic(err)
	}
	for key, value := range GetHeader() {
		req.Header.Set(key, value)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println("body:" + string(body))
}

func GetHeader() map[string]string {
	headers := map[string]string{
		"Accept":             "application/json, text/plain, */*",
		"Accept-Encoding":    "gzip, deflate, br",
		"Accept-Language":    "zh-CN,zh;q=0.9",
		"Connection":         "keep-alive",
		"Content-Length":     "149",
		"Content-Type":       "application/json;charset=UTF-8",
		"Cookie":             "uuid=uuid-fdc3229d-e37c-92c0-398b-e7206016d49b",
		"Host":               "www.kepeiol.com",
		"Origin":             "https://www.kepeiol.com",
		"Referer":            "https://www.kepeiol.com/yunkaiExamPC/",
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "same-origin",
		"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
		"authorization":      "f653c6e86f7d4fb4a30df99d319ec051",
		"hostname":           "www.kepeiol.com",
		"sec-ch-ua":          "\"Chromium\";v=\"118\", \"Google Chrome\";v=\"118\", \"Not=A?Brand\";v=\"99\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"macOS\"",
		"version":            "v2.8.7",
	}

	return headers
}

type AddVideoTimeBodyData struct {
	CharterSectionId int    `json:"charterSectionId"`
	VideoId          int    `json:"videoId"`
	VideoType        int    `json:"videoType"`
	StudyTime        int    `json:"studyTime"`
	UploadType       int    `json:"uploadType"`
	LocalCreateTime  int64  `json:"localCreateTime"`
	AppType          int    `json:"appType"`
	LastStudyTime    int    `json:"lastStudyTime"`
	ApiType          string `json:"apiType"`
}

type Video struct {
	Id       int `json:"id"`
	Duration int `json:"duration"`
}
