package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) []student {
	var students []student
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for i, d := range data {
		if i > 0 {
			t1, _ := strconv.Atoi(d[3])
			t2, _ := strconv.Atoi(d[4])
			t3, _ := strconv.Atoi(d[5])
			t4, _ := strconv.Atoi(d[6])
			s := student{firstName: d[0], lastName: d[1], university: d[2], test1Score: t1, test2Score: t2, test3Score: t3, test4Score: t4}
			students = append(students, s)
		}
	}
	return students
}

func calculateGrade(students []student) []studentStat {
	var studentStats []studentStat
	var ss studentStat
	for _, s := range students {
		ss.student = s
		ss.finalScore = float32(s.test1Score+s.test2Score+s.test3Score+s.test4Score) / 4
		if ss.finalScore < 35 {
			ss.grade = F
		} else if ss.finalScore < 70 && ss.finalScore >= 50 {
			ss.grade = B
		} else if ss.finalScore < 50 && ss.finalScore >= 35 {
			ss.grade = C
		} else {
			ss.grade = A
		}
		studentStats = append(studentStats, ss)
	}

	return studentStats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	var topper studentStat
	var maxScore float32
	for _, gs := range gradedStudents {
		if gs.finalScore > maxScore {
			maxScore = gs.finalScore
			topper = gs
		}
	}
	return topper
}

func findTopperPerUniversity(students []studentStat) map[string]studentStat {
	universityTopper := make(map[string]studentStat)
	universityStudent := make(map[string][]studentStat)
	for _, ss := range students {
		universityStudent[ss.university] = append(universityStudent[ss.university], ss)
	}
	for u, s := range universityStudent {
		t := findOverallTopper(s)
		universityTopper[u] = t
	}
	return universityTopper
}

/*
1. getting student info and finalScore
2. Every university will be having a topper
3. evaluate the maximum finalScore for each university
4. As soon as a new university comes, store it as a new key and find the max score within that university
5. append that particular map value in the array
6. return that array


1. array for students of similar university
2. max finalScore within that array
3. map that final Score with the university
*/
