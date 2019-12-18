/**
 * @Time : 2019-10-15 10:14
 * @Author : solacowa@gmail.com
 * @File : pet
 * @Software: GoLand
 */

package main

import (
	"fmt"
)

type Pet interface {
	GetPetType() string
}

//type Pet struct {
//	petType string
//}

type Dog struct {
	petType string
}

type Cat struct {
	petType string
}

func NewDog() Pet {
	return &Dog{petType: "dog"}
}

func NewCat() Pet {
	return &Cat{petType: "cat"}
}

func (c *Dog) GetPetType() string {
	return c.petType
}

func (c *Cat) GetPetType() string {
	return c.petType
}

type petQueue struct {
	catQueue []Pet
	dogQueue []Pet
	queue    []Pet
}

func (c *petQueue) Add(entry Pet) {
	switch entry.GetPetType() {
	case "dog":
		c.dogQueue = append(c.dogQueue, entry)
	case "cat":
		c.catQueue = append(c.catQueue, entry)
	}
	c.queue = append(c.queue, entry)
}

func (c *petQueue) PollAll() {

}

func (c *petQueue) PollDog() {
	for _, v := range c.queue {
		if v.GetPetType() == "dog" {
			fmt.Println(v.GetPetType())
		}
	}
}

func (c *petQueue) PollCat() {
	for _, v := range c.queue {
		if v.GetPetType() == "cat" {
			fmt.Println(v.GetPetType())
		}
	}
}

func (c *petQueue) IsEmpty(petType string) bool {
	for _, v := range c.queue {
		if v.GetPetType() == petType {
			return true
		}
	}
	return false
}

func main() {
	dog := NewDog()
	cat := NewCat()

	d := dogCatQueue{}
	d.Add(dog)
	d.Add(cat)
	d.Add(cat)
	d.Add(dog)
	d.Add(dog)
	d.Add(dog)
	d.Add(dog)
	d.Add(cat)
	d.Add(cat)
	d.Add(dog)
	d.Add(cat)

	fmt.Println(d.PollAll().GetPetType())

	fmt.Println(d.PollDog().GetPetType())

}

type petEnterQueue struct {
	pet   Pet
	count int64
}

func (c *petEnterQueue) getPet() Pet {
	return c.pet
}

func (c *petEnterQueue) getCount() int64 {
	return c.count
}

func (c *petEnterQueue) getEnterPetType() string {
	return c.pet.GetPetType()
}

type dogCatQueue struct {
	dogQueue []petEnterQueue
	catQueue []petEnterQueue
	count    int64
}

func (c *dogCatQueue) Add(entry Pet) {
	switch entry.GetPetType() {
	case "dog":
		c.count += 1
		c.dogQueue = append(c.dogQueue, petEnterQueue{pet: entry, count: c.count})
	case "cat":
		c.count += 1
		c.catQueue = append(c.catQueue, petEnterQueue{pet: entry, count: c.count})
	}
}

func (c *dogCatQueue) PollAll() Pet {
	dogLen := len(c.dogQueue)
	catLen := len(c.catQueue)
	if dogLen > 0 && catLen > 0 {
		if c.dogQueue[0].getCount() < c.catQueue[0].getCount() {
			return c.dogQueue[dogLen-1].getPet()
		} else {
			return c.catQueue[catLen-1].getPet()
		}
	} else if dogLen > 0 {
		return c.dogQueue[dogLen-1].getPet()
	} else if catLen > 0 {
		return c.catQueue[catLen-1].getPet()
	}
	return nil
}

func (c *dogCatQueue) PollDog() Pet {
	if !c.isDogQueueEmpty() {
		dogLen := len(c.dogQueue)
		return c.dogQueue[dogLen].getPet()
	}
	return nil
}

func (c *dogCatQueue) PollCat() Pet {
	if !c.isCatQueueEmpty() {
		catLen := len(c.catQueue)
		pet := c.catQueue[catLen].getPet()
		c.catQueue = c.catQueue[:catLen-2]
		return pet
	}
	return nil
}

func (c *dogCatQueue) isEmpty() bool {
	if len(c.dogQueue) == 0 && len(c.catQueue) == 0 {
		return true
	}
	return false
}

func (c *dogCatQueue) isCatQueueEmpty() bool {
	if len(c.catQueue) > 0 {
		return false
	}
	return true
}

func (c *dogCatQueue) isDogQueueEmpty() bool {
	if len(c.dogQueue) > 0 {
		return false
	}
	return true
}
