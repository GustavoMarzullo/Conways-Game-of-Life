//Page 155 - Get Programming with Go - Nathan Youngman, Roger Peppé (2018)
package main

import(
	"fmt"
	"math/rand"
	"time"
)

// Decide on the dimensions of the grid and define some constants

const(
	width = 80
	height = 15
	maxiter=100
	TimePerIteration = 1 //seconds
	LiveCell = "o"
	DeadCell = " "
	EscapeSequence = "\033[H\033[2J"
	)

/* Define a Universe type to hold a two-dimensional field of cells. With a Boolean type
each cell will be either dead ( false ) or alive ( true )*/

type Universe [][]bool

/* Write a NewUniverse function that uses make to allocate and return a Universe*/

func NewUniverse() Universe{
	u := make(Universe, height)
	for i:= range u{
		u[i] = make([]bool,width)
	}
	return u
} 

/*Write a method to print a universe to the screen using the fmt package. Be sure to move to a new line after
printing each row*/

func (u Universe) Show(){
	for h:=0;h<height;h++{
		for w:=0;w<width;w++{
			if u[h][w]{
				fmt.Printf(LiveCell)
			}else{
				fmt.Printf(DeadCell)
			}
		}
		fmt.Printf("\n")
	}
}

//Write a Seed method that randomly sets approximately 25% of the cells to alive (true)

func (u Universe) Seed(){
	var n int
	for h:=0;h<height;h++{
		for w:=0;w<width;w++{
			n = rand.Intn(100) + 1
			if n<=25{
				u[h][w]=true
			}else{
				u[h][w]=false
			}
		}
	}
}

/*It should be easy to determine whether a cell is dead or alive. Just look up a cell in the
Universe slice. If the Boolean is true , the cell is alive.
Write an Alive method on the Universe type with the following signature*/

func (u Universe) Alive(h,w int) bool{
	//making sure the universe wraps arround
	if w<0{
		w+=width
	}
	if w>=width{
		w %= width
	}
	if h<0{
		h+=height
	}
	if h>=height{
		h %= height
	}
	
	//checking if it's alive or not
	return u[h][w]
}

//Write a method to count the number of live neighbors for a given cell, from 0 to 8

func (u Universe) Neighbors(h, w int) int{
	neighbors := 0
	for H:=-1;H<=1;H++{
		for W:=-1;W<=1;W++{
			if H!=0 || W!=0{//making sure not to count the cell itself
				if u.Alive(h+H,w+W){
					neighbors++
				}
			}
		}
	}
	return neighbors
}

/*Now that you can determine whether a cell has two, three, or more neighbors, you can
implement the rules shown at the beginning of this section.
RULES:
	- A live cell with less than two live neighbors dies.
	- A live cell with two or three live neighbors lives on to the next generation.
	- A live cell with more than three live neighbors dies.
	- A dead cell with exactly three live neighbors becomes a live cell.
*/

func (u Universe) Next(h, w int) bool{ 
	n := u.Neighbors(h,w)
	if u[h][w] { //if a cell is alive
		if n==2 || n==3 {
			return true
		}else{
			return false
		}
	}else{//if a cell is dead
		if n==3{
			return true
		}else{
			return false
		}
	}
}

func run(){
	u:=NewUniverse()
	u.Seed()	
	dummy := NewUniverse()
	for iter:=0;iter<maxiter;iter++{
		for h:=0;h<height;h++{
			for w:=0;w<width;w++{
				dummy[h][w]=u.Next(h,w)
			}
		}
		u,dummy=dummy,u
		fmt.Print(EscapeSequence)
		u.Show()
		time.Sleep(TimePerIteration * time.Second)
	}
}


//main func
func main(){
	run()
}

