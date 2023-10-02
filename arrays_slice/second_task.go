package arrays_slice

type TeachersJournalWithParallelAssessment struct {
	GradesOf9V        []float32
	GpaOf9V           float32
	GradesOf9B        []float32
	GpaOf9B           float32
	GradesOf9A        []float32
	GpaOf9A           float32
	GpaBetweenClasses float32
}

func CountAverageGradeInClass(classGrades []float32) float32 {
	var totalSum float32
	var totalAmountOfGrades int = len(classGrades)
	for _, grades := range classGrades {
		totalSum += grades
	}
	var GPA float32 = totalSum / float32(totalAmountOfGrades)
	return GPA
}

func CountAverageGradeBetweenParallelClasses(averageGradesOfClasses [3]float32) float32 {
	return (averageGradesOfClasses[0] + averageGradesOfClasses[1] + averageGradesOfClasses[2]) / 3
}
