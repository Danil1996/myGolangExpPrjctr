package main

import (
	"fmt"

	"github.com/myGolangExpPrjctr/arrays_slice"
)

// TODO: just run it
func main() {
	// region FIRST TASK
	fmt.Println("START OF FIRST TASK\n")
	myTexts := arrays_slice.UserText{}
	myTexts.AppendText("Hello")
	myTexts.AppendText("Hello, guys")
	myTexts.AppendText("Kon'nichiwa, guys")
	myTexts.AppendText("Kon'nichiwa, girls")
	myTexts.AppendText("Hello, girls")
	myTexts.AppendText("Zdorovenki buli, guys")
	myTexts.AppendText("Zdorovenki buli, girls")
	myTexts.ShowMyInputs()
	myTexts.FindSubstringInInputs("girls")
	fmt.Println("\nEND OF FIRST TASK\n")
	// endregion FIRST TASK

	// region SECOND TASK
	fmt.Println("START OF SECOND TASK\n")
	tatyanaIvanovnaJournal := arrays_slice.TeachersJournalWithParallelAssessment{
		GradesOf9V:        []float32{9.6, 11.5, 3.5, 8.7, 7.7, 10.9},
		GradesOf9B:        []float32{8.6, 10.5, 4.5, 9.7, 4.7, 10.8},
		GradesOf9A:        []float32{9.1, 11.9, 3.0, 2.3, 7.9, 12.0},
		GpaBetweenClasses: 0,
	}
	tatyanaIvanovnaJournal.GpaOf9V = arrays_slice.CountAverageGradeInClass(tatyanaIvanovnaJournal.GradesOf9V)
	fmt.Printf("GpaOf9V %v\n", tatyanaIvanovnaJournal.GpaOf9V)

	tatyanaIvanovnaJournal.GpaOf9B = arrays_slice.CountAverageGradeInClass(tatyanaIvanovnaJournal.GradesOf9B)
	fmt.Printf("GpaOf9B %v\n", tatyanaIvanovnaJournal.GpaOf9B)

	tatyanaIvanovnaJournal.GpaOf9A = arrays_slice.CountAverageGradeInClass(tatyanaIvanovnaJournal.GradesOf9A)
	fmt.Printf("GpaOf9A %v\n", tatyanaIvanovnaJournal.GpaOf9A)

	gpaOfClasses := [3]float32{
		tatyanaIvanovnaJournal.GpaOf9V,
		tatyanaIvanovnaJournal.GpaOf9B,
		tatyanaIvanovnaJournal.GpaOf9A,
	}
	tatyanaIvanovnaJournal.GpaBetweenClasses = arrays_slice.CountAverageGradeBetweenParallelClasses(gpaOfClasses)
	fmt.Printf("GpaBetweenClasses %v\n", tatyanaIvanovnaJournal.GpaBetweenClasses)
	fmt.Println("\nEND OF SECOND TASK")
	// endregion SECOND TASK

}
