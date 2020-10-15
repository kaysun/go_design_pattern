package responsibility_chain

import "testing"

func Test(t *testing.T) {
	t.Run("responsibility_chain_200: ", ChainOfResponsibility200)
	t.Run("responsibility_chain_3000: ", ChainOfResponsibility3000)
	t.Run("responsibility_chain_6000: ", ChainOfResponsibility6000)
}

// ChainOfResponsibility200 200元报销请求
func ChainOfResponsibility200(t *testing.T) {
	var leader, director, cfo Manager
	leader = &Leader{}
	director = &Director{}
	cfo = &CFO{}

	handlerChain := HandlerChain{}
	handlerChain.AddHandler(leader)
	handlerChain.AddHandler(director)
	handlerChain.AddHandler(cfo)

	request := Request{
		Name:   "kay",
		Amount: 200,
	}
	handlerChain.HandleRequest(request)
}

// ChainOfResponsibility3000 3000元报销请求
func ChainOfResponsibility3000(t *testing.T) {
	var leader, director, cfo Manager
	leader = &Leader{}
	director = &Director{}
	cfo = &CFO{}

	handlerChain := HandlerChain{}
	handlerChain.AddHandlers(leader, director, cfo)

	request := Request{
		Name:   "kay",
		Amount: 3000,
	}
	handlerChain.HandleRequest(request)
}

// ChainOfResponsibility6000 6000元报销请求
func ChainOfResponsibility6000(t *testing.T) {
	var leader, director, cfo Manager
	leader = &Leader{}
	director = &Director{}
	cfo = &CFO{}

	handlerChain := HandlerChain{}
	handlerChain.AddHandlers(leader, director, cfo)

	request := Request{
		Name:   "kay",
		Amount: 6000,
	}
	handlerChain.HandleRequest(request)
}