package main
import(
	"fmt"
	"os"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)
type deck []string
func (d deck)print(){
	for i,card:=range d{
		fmt.Println(i,card)
	}
	
}
func newDeck() deck{
	cards:=deck{}
	cardSuits:=[]string{"Spades","Hearts","Diamonds","Club"}
	cardValues:=[]string{"Ace","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","King","Queen"}
	for _,suit:=range cardSuits{
		for _,value:=range cardValues{
			cards=append(cards,value+" of "+suit)
		}
	}
	return cards
}
func deal(d deck,handSize int)(deck,deck){
	return d[:handSize],d[:handSize]
}
func (d deck)toString() string{
	return strings.Join([]string(d),",")
}
func (d deck) saveToFile(fileName string)error{
	return ioutil.WriteFile(fileName,[]byte(d.toString()),0666)
}
func newDeckFromFile(fileName string)deck{
	bs,err:=ioutil.ReadFile(fileName)
	if err!=nil{
		fmt.Println("Error")
	}
	os.Exit(1)
	S:=strings.Split(string(bs),",")
	return deck(S)
}
func (d deck)shufflw(){
	source:=rand.NewSource(time.Now().UnixNano())
	r:=rand.New(source)
	
	for i:=range d{
		newPosition:=r.Intn(len(d)-1)
		d[i],d[newPosition]=d[newPosition],d[i]
	}
}
