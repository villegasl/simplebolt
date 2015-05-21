package main

import (
	"fmt"
	"log"

	"github.com/xyproto/simplebolt"
)

func main() {
	db := simplebolt.New("bolt.db")
	defer db.Close()

	kv := simplebolt.NewKeyValue(db, "fruit")
	if err := kv.Set("banana", "yes"); err != nil {
		log.Println("Could not set a key+value.")
		log.Fatalln(err)
	}

	val, err := kv.Get("banana")
	if err != nil {
		log.Println("Could not get value.")
		log.Fatalln(err)
	}
	fmt.Println("Got it:", val)

	kv.Set("banana", "2")
	kv.Inc("banana")
	three, err := kv.Get("banana")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Three:", three)

	kv.Inc("fnu")
	fnu, err := kv.Get("fnu")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("fnu", fnu)

	val, err = kv.Get("doesnotexist")
	//fmt.Println("does not exist", val, err)

	kv.Remove()

	l := simplebolt.NewList(db, "fruit")

	l.Add("kiwi")
	l.Add("banana")
	l.Add("pear")
	l.Add("apple")

	if results, err := l.GetAll(); err == nil {
		for _, fruit := range results {
			fmt.Println(fruit, "is a fruit")
		}
	}

	last, err := l.GetLast()
	if err == nil {
		fmt.Println("The last one is:", last)
	}

	lastN, err := l.GetLastN(3)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("The last 3 are:", lastN)

	l.Remove()

	s := simplebolt.NewSet(db, "numbers")
	s.Add("9")
	s.Add("7")
	s.Add("2")
	s.Add("2")
	s.Add("2")
	s.Add("7")
	s.Add("8")
	v, err := s.GetAll()
	if err == nil {
		fmt.Println(v)
	}
	s.Remove()

	val, err = kv.Inc("counter")
	if (val != "1") || (err != nil) {
		log.Println("counter should be 1 but is", val)
	}
	kv.Remove()

	h := simplebolt.NewHashMap(db, "counter")
	h.Set("bob", "password", "hunter1")
	h.Set("bob", "email", "bob@zombo.com")
	fmt.Println(h.GetAll())

	ok, err := h.Has("bob", "password")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("bob ok?", ok)

	ok, err = h.Exists("bob")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("really?", ok)

	h.Remove()

	ok, err = h.Has("bob", "password")
	if err == nil {
		log.Fatalln("not supposed to exist")
	}
	fmt.Println("ok")

	ok, err = h.Exists("bob")
	if err == nil {
		log.Fatalln("not supposed to exist")
	}
	fmt.Println("ok")

}
