package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	water = 200
	milk  = 50
	beans = 15
)

type Ingredients struct {
	water, milk, beans int
}

type Coffee struct {
	coffeeType  CoffeeType
	ingredients Ingredients
}

type CoffeeMachine struct {
	Coffee
	disposableCups, money int
}

type CoffeeType int

const (
	espresso CoffeeType = iota + 1
	latte
	cappuccino
)

type Action string

const (
	BUY       Action = "buy"
	FILL      Action = "fill"
	TAKE      Action = "take"
	REMAINING Action = "remaining"
	EXIT      Action = "exit"
)

func main() {
	machine := CoffeeMachine{
		Coffee: Coffee{
			ingredients: Ingredients{
				water: 400,
				milk:  540,
				beans: 120,
			},
		},
		disposableCups: 9,
		money:          550,
	}

	for {
		fmt.Printf("Write action (buy, fill, take, remaining, exit):\n")
		process(&machine)
		fmt.Println()
	}
}

func getInput() int {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return 0
	}
	res, _ := strconv.Atoi(input)
	return res
}

// Calculate the required amount of ingredients and return them (`water`, `milk`, `beans`).
func calculate(cup int) (int, int, int) {
	return cup * water, cup * milk, cup * beans
}

// Prints how many from each ingredient is required to create X cup(s) of coffee.
func printAmount(answers []int) {
	w, m, b := calculate(answers[len(answers)-1])
	fmt.Printf(
		"For %d cups of coffee you will need:\n%d ml of water\n%d ml of milk\n%d g of coffee beans\n",
		answers[len(answers)-1],
		w,
		m,
		b,
	)
}

// Returns the `minimum` and `maximum` number of cups of coffee that can be created.
func numOfCups(w, m, b, cup int) (int, int) {
	wtr := w / water
	mlk := m / milk
	bns := b / beans
	min := math.Min(float64(bns), math.Min(float64(wtr), float64(mlk)))
	max := math.Abs(float64(cup) - min)
	return int(min), int(max)
}

// Print to stdout whether the given amount of coffee can be created or not.
func result(answers []int) {
	var canMake bool
	w, m, b := calculate(answers[len(answers)-1])
	min, max := numOfCups(answers[0], answers[1], answers[2], answers[3])

	if (w < answers[0] && m < answers[1] && b < answers[2]) || answers[len(answers)-1] == 0 {
		canMake = true
	}
	if canMake {
		if max != 0 {
			fmt.Printf(
				"Yes, I can make that amount of coffee (and even %d more than that)\n",
				max,
			)
		} else {
			fmt.Printf("Yes, I can make that amount of coffee\n")
		}
	} else {
		fmt.Printf(
			"No, I can make only %d cups of coffee",
			min,
		)
	}
}

// Takes an `Action` model the behavior the coffee machine.
func process(machine *CoffeeMachine) {
	var action string
	fmt.Scan(&action)
	strings.ToLower(action)
	switch Action(action) {
	case BUY:
		machine.buy()
	case FILL:
		machine.fill()
	case TAKE:
		machine.take()
	case REMAINING:
		machine.remaining()
	case EXIT:
		machine.exit()
	}
}

// Return the string representation of `Coffee Machine`.
func (cm *CoffeeMachine) String() string {
	return fmt.Sprintf(
		"The coffee machine has:\n"+
			"%d ml of water\n"+
			"%d ml of milk\n"+
			"%d g of coffee beans\n"+
			"%d disposable cups\n"+
			"$%d of money",
		cm.ingredients.water,
		cm.ingredients.milk,
		cm.ingredients.beans,
		cm.disposableCups,
		cm.money,
	)
}

// The `func checkResources` takes the `Ingredients` - water, milk, beans -,
// as argument, and decide if there are enough resources for serving the coffee.
func (cm *CoffeeMachine) checkResources(resource Ingredients) {
	if cm.disposableCups < 1 {
		fmt.Println("Sorry, not enough disposable cups!")
		return
	}

	if cm.ingredients.water-resource.water >= 0 {
		cm.ingredients.water -= resource.water
	} else {
		fmt.Println("Sorry, not enough water!")
		return
	}
	if cm.ingredients.milk-resource.milk >= 0 {
		cm.ingredients.milk -= resource.milk
	} else {
		fmt.Println("Sorry, not enough milk!")
		return
	}
	if cm.ingredients.beans-resource.beans >= 0 {
		cm.ingredients.beans -= resource.beans
	} else {
		fmt.Println("Sorry, not enough beans!")
		return
	}
	fmt.Printf("I have enough resources, making you a coffee!\n")
	cm.disposableCups--
}

// Buy control `Action` makes it possible to choose from the available `CoffeeType`s.
// Upon invalid input, we end up in the main loop.
func (cm *CoffeeMachine) buy() {
	fmt.Println()
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back to main menu:")
	input := getInput()
	switch CoffeeType(input) {
	// For one espresso, the coffee machine needs 250 ml of water and 16 g of coffee beans. It costs $4.
	case espresso:
		cm.checkResources(Ingredients{
			water: 250,
			milk:  0,
			beans: 16,
		})
		cm.money += 4
	// For a latte, the coffee machine needs 350 ml of water, 75 ml of milk, and 20 g of coffee beans.
	// It costs $7.
	case latte:
		cm.checkResources(Ingredients{
			water: 350,
			milk:  75,
			beans: 20,
		})
		cm.money += 7
	// And for a cappuccino, the coffee machine needs 200 ml of water, 100 ml of milk,
	// and 12 g of coffee beans. It costs $6.
	case cappuccino:
		cm.checkResources(Ingredients{
			water: 200,
			milk:  100,
			beans: 12,
		})
		cm.money += 6
	default:
		return
	}
}

// Fills the `CoffeeMachine`'s resources.
func (cm *CoffeeMachine) fill() {
	fmt.Println()
	questions := []string{
		"Write how many ml of water you want to add:",
		"Write how many ml of milk you want to add:",
		"Write how many grams of coffee beans you want to add:",
		"Write how many disposable cups you want to add:",
	}
	answers := make([]int, 0, 4)
	for _, question := range questions {
		fmt.Println(question)
		input := getInput()
		answers = append(answers, input)
	}
	cm.ingredients.water += answers[0]
	cm.ingredients.milk += answers[1]
	cm.ingredients.beans += answers[2]
	cm.disposableCups += answers[3]
}

// Zeroes/takes out the available `money`.
func (cm *CoffeeMachine) take() {
	fmt.Printf("I give you $%d\n", cm.money)
	cm.money -= cm.money
}

// Queries the `CoffeeMachine` resources.
func (cm *CoffeeMachine) remaining() {
	fmt.Printf("\n%v\n", cm)
}

// Shuts down the program.
func (cm *CoffeeMachine) exit() {
	os.Exit(0)
}
