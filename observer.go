package main

import "fmt"

type instagram interface {
	followers(observer)
	unfollowers(observer)
	delete()
	publishAll()
}
type Instagram struct {
	fans []observer
	posts []string
	information []string
}

func (i *Instagram)followers(Observer observer)  {
	i.fans=append(i.fans,Observer)
}
func(i*Instagram) unfollwers(Observer observer){
	i.fans=DeletePerson(i.fans,Observer)
}
func (i*Instagram)publishAll()  {
	for _, observer := range i.fans {
		observer.update(i.posts)
	}
}
func (i*Instagram)delete()  {
	i.delete()
}
func (i*Instagram) addPostWithPicture(name string) {
	i.posts=append(i.posts,name)
	i.publishAll()
}
func (i*Instagram)deletePost(name string){
	i.deletePost("Today is gift")
	i.publishAll()
}
type observer interface {
	update([]string)
	getID() int
}
type Person struct {
	name string
	id int
	age int
}


func (p*Person)getID() int {
	return p.id
}
func (p*Person) update(posts []string)  {
	fmt.Println("Hello,",p.name)
	fmt.Println("There are some new posts")
	for i,post:=range posts{
		fmt.Println(i,post)
	}
}
func DeletePerson(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

func main()  {
	ins:=Instagram{
		fans:        []observer{},
		posts:       []string{},
		information: []string{},
	}
	ins.addPostWithPicture("Yesterday is history")
	Pop:=&Person{"Pop",1,18}
	ins.followers(Pop)
	ins.addPostWithPicture("Tomorrow is mystery")
	ins.addPostWithPicture("Today is gift,why we call present")
	ins.unfollwers(Pop)
	ins.addPostWithPicture("Nice day")
}
