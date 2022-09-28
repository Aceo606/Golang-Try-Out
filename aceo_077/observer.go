package main

import "fmt"

func main() {

	patternDeadline := newDeadline("Software design patterns")
	calcDeadline := newDeadline("Calculus 2 HomeWork")
	observerFirst := &Student{barcode: "211265"}
	observerSecond := &Student{barcode: "211223"}
	observerThird := &Student{barcode: "207207"}

	patternDeadline.register_student(observerFirst)  //добавили студента на подписку новости о дедлайнах
	patternDeadline.register_student(observerSecond) //добавили студента на подписку новости о дедлайнах
	patternDeadline.register_student(observerThird)  //добавили студента на подписку новости о дедлайнах
	calcDeadline.register_student(observerFirst)     //добавили студента на подписку новости о дедлайнах
	calcDeadline.register_student(observerSecond)    //добавили студента на подписку новости о дедлайнах
	calcDeadline.register_student(observerThird)     //добавили студента на подписку новости о дедлайнах
	patternDeadline.updateTaskAvailability()
	calcDeadline.updateTaskAvailability()
	patternDeadline.deregister_student(observerThird) //удалили студента из листа подписчиков новостей о дедлайнах
	patternDeadline.updateTaskAvailability()
}

type Student struct {
	barcode string
}

func (c *Student) update(deadlineName string) {
	fmt.Printf("Sending message to student %s for deadline %s\n", c.barcode, deadlineName)
}

func (c *Student) getBarcode() string {
	return c.barcode
}

type Deadline struct {
	observerList []Observer
	course_name  string
	onMoodle     bool
}

func newDeadline(name string) *Deadline {
	return &Deadline{
		course_name: name,
	}
}
func (i *Deadline) updateTaskAvailability() {
	fmt.Printf("Deadline %s is now on moodle\n", i.course_name)
	i.onMoodle = true
	i.notifyAll()
}

type Observer interface {
	update(string)
	getBarcode() string
}

func (i *Deadline) register_student(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Deadline) deregister_student(o Observer) {
	i.observerList = removeStudent(i.observerList, o)
}

func (i *Deadline) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.course_name)
	}
}

func removeStudent(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getBarcode() == observer.getBarcode() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
