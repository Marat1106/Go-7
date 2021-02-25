package main
import "fmt"
type Request struct {
	RequestType    string
	RequestContent string
	Number         int
}

func (r*Request)String()string  {
	return fmt.Sprintf("some rest")
}
func NewRequest(RequestType,RequestContent string,Number int )*Request{
	return &Request{RequestType: RequestType,RequestContent: RequestContent,Number: Number}
}
type Manager interface {
	SetNext(next Manager)
	RequestHandler(request Request)
}

type CommonManager struct {
	Manager
	Name string
}

func (cm *CommonManager) SetNext(next Manager) {
	cm.Manager = next
}

func (cm *CommonManager) RequestHandler(request Request) {
	if request.RequestType == "Ask for leave" && request.Number <= 2 {
		fmt.Printf("%s: %s number %d Approved\n", cm.Name, request.RequestContent, request.Number)
	} else {
		if cm.Manager != nil {
			cm.Manager.RequestHandler(request)
		}

	}
}

type MajorManager struct {
	Manager
	Name string
}

func (mm *MajorManager) SetNext(next Manager) {
	mm.Manager = next
}

func (mm *MajorManager) RequestHandler(request Request) {
	if request.RequestType == "Ask for leave" && request.Number <= 5 {
		fmt.Printf("%s: %s number %d Approved\n", mm.Name, request.RequestContent, request.Number)
	} else {
		if mm.Manager != nil {
			mm.Manager.RequestHandler(request)
		}
	}
}

type GeneralManager struct {
	Manager
	Name string
}

func (gm *GeneralManager) SetNext(next Manager) {
	gm.Manager = next
}

func (gm *GeneralManager) RequestHandler(request Request) {
	if request.RequestType == "Ask for leave" {
		fmt.Printf("%s: %s number %d Approved\n", gm.Name, request.RequestContent, request.Number)
	} else if request.RequestType == "Raise salary" && request.Number <= 500 {
		fmt.Printf("%s: %s number %d Approved\n", gm.Name, request.RequestContent, request.Number)
	} else if request.RequestType == "Raise salary" && request.Number > 500 {
		fmt.Printf("%s: %s number %d ask again\n", gm.Name, request.RequestContent, request.Number)
	}
}
func main()  {
	John:=&CommonManager{Name: "John"}
	James:=&MajorManager{Name: "James"}
    Ella:=&GeneralManager{Name: "Ella"}
    John.SetNext(James)
    James.SetNext(Ella)
    request:=NewRequest("Ask for leave","take some rest",3)
    John.RequestHandler(*request)
    request=NewRequest("Ask for leave","get sick",2)
	James.RequestHandler(*request)
    request=NewRequest("Raise salary","Work so hard",499)
    Ella.RequestHandler(*request)
    Mason:=&GeneralManager{Name: "Mason"}
    request=NewRequest("Raise salary","lack of money",501)
    Mason.RequestHandler(*request)
}