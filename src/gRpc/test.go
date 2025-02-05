package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"reflect"
)

type Arith struct {
}

type ArithRequest struct {
	A int
	B int
}

type ArithResponse struct {
	Pro int
	Quo int
	Rem int
}

func (*Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

func (*Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

func main() {
	v := reflect.ValueOf(3) // a reflect.Value
	fmt.Println(v)          // "3"
	fmt.Printf("%v\n", v)   // "3"
	fmt.Println(v.String()) // NOTE: "<int Value>"
	t := v.Type()           // a reflect.Type
	fmt.Println(t.String()) // "int"
	rpc.Register(new(Arith))
	//rpc.HandleHTTP() //http协议
	lis, err := net.Listen("tcp", "127.0.0.1:8095")
	if err != nil {
		log.Fatalln("fatal error :", err)
	}
	fmt.Fprintf(os.Stdout, "%s", "start connection")
	//http.Serve(lis, nil)
	for {
		conn, err := lis.Accept() //接收请求
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			fmt.Fprintf(os.Stdout, "%s", "new client in coming\n")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}

type List struct {
	Val  int
	Next *List
}

func hasCycle(head *List) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow.Next = slow
		if slow == fast {
			return true
		}
	}
	return false
}

func findCycleStart(head *List) *List {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow.Next = slow
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func getIntersectionNode(headA *List, headB *List) *List {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p1 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
