package symphony

type Problem struct {
	category int
	pro      string
	ansA     string
	ansB     string
	answer   string
}
type ProblemSet struct {
	problems []Problem
}

type Result struct {
	res1 string
	res2 string
	res3 string
	res4 string
}

func getProblems() *ProblemSet {
	problems := make([]Problem, 0)
	// 1
	problem := Problem{
		pro:      "当你要外出一整天，你会：",
		ansA:     "计划要做什么和什么时候做",
		ansB:     "说去就去",
		category: 4,
	}
	problems = append(problems, problem)

	//2
	problem = Problem{
		pro:      "你是否：",
		ansA:     "容易让人了解",
		ansB:     "难以让人了解",
		category: 1,
	}
	problems = append(problems, problem)

	//3
	problem = Problem{
		pro:      "你认为自己是一个：",
		ansA:     "比较随兴所至的人",
		ansB:     "比较有条理的人",
		category: 4,
	}
	problems = append(problems, problem)
	//4
	problem = Problem{
		pro:      "假如你是教师你愿意教：",
		ansA:     "以事实为主的课程",
		ansB:     "涉及理论的课程",
		category: 4,
	}
	problems = append(problems, problem)

	//5
	problem = Problem{
		pro:      "处理许多事情时，你喜欢：",
		ansA:     "随兴所至行事",
		ansB:     "按照计划行事",
		category: 4,
	}
	problems = append(problems, problem)
	//6
	problem = Problem{
		pro:      "下面哪个词语更合你心意：",
		ansA:     "仁慈慷慨的",
		ansB:     "意志坚定地",
		category: 4,
	}
	problems = append(problems, problem)
	//7
	problem = Problem{
		pro:      "按照程序表做事：",
		ansA:     "合你心意",
		ansB:     "令你感到束缚",
		category: 4,
	}
	problems = append(problems, problem)
	//8
	problem = Problem{
		pro:      "你做事多数是：",
		ansA:     "按当天心情去做",
		ansB:     "照拟好的程序表做",
		category: 4,
	}
	problems = append(problems, problem)

	//9
	problem = Problem{
		pro:      "你倾向：",
		ansA:     "重视情感多于逻辑",
		ansB:     "难以让人了解",
		category: 4,
	}
	problems = append(problems, problem)

	//10
	problem = Problem{
		pro:      "你做事多数是：",
		ansA:     "令你活力倍增",
		ansB:     "难以让人了解",
		category: 4,
	}
	problems = append(problems, problem)
	//11
	problem = Problem{
		pro:      "你做事多数是：",
		ansA:     "按当天心情去做",
		ansB:     "难以让人了解",
		category: 4,
	}
	problems = append(problems, problem)

	//12
	problem = Problem{
		pro:      "在大多数情况下，你会选择：",
		ansA:     "顺其自然",
		ansB:     "程序表做事",
		category: 4,
	}
	problems = append(problems, problem)
	//13
	problem = Problem{
		pro:      "你通常：",
		ansA:     "与人容易混熟",
		ansB:     "比较沉默和矜持",
		category: 4,
	}
	problems = append(problems, problem)

	//14
	problem = Problem{
		pro:      "哪些人会更吸引你：",
		ansA:     "思维敏捷、聪颖的人",
		ansB:     "实事求是、知识丰富的人",
		category: 4,
	}
	problems = append(problems, problem)
	//15
	problem = Problem{
		pro:      "你做事多数是：",
		ansA:     "按当天心情去做",
		ansB:     "难以让人了解",
		category: 4,
	}
	problems = append(problems, problem)

	//16
	problem = Problem{
		pro:      "你做事多数是：",
		ansA:     "按当天心情去做",
		ansB:     "难以让人了解",
		category: 4,
	}
	problems = append(problems, problem)
	//17
	problem = Problem{
		pro:      "你做事多数是：",
		ansA:     "按当天心情去做",
		ansB:     "难以让人了解",
		category: 4,
	}
	problems = append(problems, problem)

	//18
	problem = Problem{
		pro:      "你做事多数是：",
		ansA:     "按当天心情去做",
		ansB:     "难以让人了解",
		category: 4,
	}
	problems = append(problems, problem)
	//19
	problem = Problem{
		pro:      "你做事多数是：",
		ansA:     "按当天心情去做",
		ansB:     "难以让人了解",
		category: 4,
	}
	problems = append(problems, problem)

	//20
	problem = Problem{
		pro:      "你做事多数是：",
		ansA:     "按当天心情去做",
		ansB:     "难以让人了解",
		category: 4,
	}
	problems = append(problems, problem)
	//21
	problem = Problem{
		pro:      "你做事多数是：",
		ansA:     "按当天心情去做",
		ansB:     "难以让人了解",
		category: 4,
	}
	problems = append(problems, problem)

	//22
	problem = Problem{
		pro:      "你做事多数是：",
		ansA:     "按当天心情去做",
		ansB:     "难以让人了解",
		category: 4,
	}
	problems = append(problems, problem)

	//23
	problem = Problem{
		pro:      "你做事多数是：",
		ansA:     "按当天心情去做",
		ansB:     "难以让人了解",
		category: 4,
	}
	problems = append(problems, problem)

	//24
	problem = Problem{
		pro:      "你做事多数是：",
		ansA:     "按当天心情去做",
		ansB:     "难以让人了解",
		category: 4,
	}
	problems = append(problems, problem)

	//25
	problem = Problem{
		pro:      "你做事多数是：",
		ansA:     "按当天心情去做",
		ansB:     "难以让人了解",
		category: 4,
	}
	problems = append(problems, problem)

	//26
	problem = Problem{
		pro:      "你做事多数是：",
		ansA:     "按当天心情去做",
		ansB:     "难以让人了解",
		category: 4,
	}
	problems = append(problems, problem)
	//27
	problem = Problem{
		pro:      "你做事多数是：",
		ansA:     "按当天心情去做",
		ansB:     "难以让人了解",
		category: 4,
	}
	problems = append(problems, problem)

	//28
	problem = Problem{
		pro:      "你做事多数是：",
		ansA:     "按当天心情去做",
		ansB:     "难以让人了解",
		category: 4,
	}
	problems = append(problems, problem)

	return &ProblemSet{problems: problems}
}
func calculateMBTI(problemSet ProblemSet) *Result {
	var res1 string
	var res2 string
	var res3 string
	var res4 string

	var r1 int32
	var r2 int32
	var r3 int32
	var r4 int32

	for i, problem := range problemSet.problems {
		if i <= 6 && i >= 0 {
			if problem.answer == "A" {
				r1++
			} else {
				r1--
			}
		}
		if r1 < 0 {
			res1 = "I"
		} else {
			res1 = "E"
		}

		if i <= 13 && i >= 7 {
			if problem.answer == "A" {
				r2++
			} else {
				r2--
			}
		}
		if r2 < 0 {
			res2 = "S"
		} else {
			res2 = "N"
		}

		if i <= 20 && i >= 14 {
			if problem.answer == "A" {
				r3++
			} else {
				r3--
			}
		}

		if r3 < 0 {
			res3 = "T"
		} else {
			res3 = "F"
		}

		if i <= 27 && i >= 21 {
			if problem.answer == "A" {
				r4++
			} else {
				r4--
			}
		}

		if r4 < 0 {
			res4 = "J"
		} else {
			res4 = "P"
		}
	}
	return &Result{
		res1: res1,
		res2: res2,
		res3: res3,
		res4: res4,
	}
}
