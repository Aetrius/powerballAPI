package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func powerballRoll(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome home!")

	sum := 0
	var jsonOut string

	jsonOut += fmt.Sprintf("{\"status\": \"ok\", \"rolls\": [")

	for i := 1; i < 20; i++ {
		sum += i

		ball_roll1 := roll(0, 0, 0, 0, 0)
		ball_roll2 := roll(ball_roll1, 0, 0, 0, 0)
		ball_roll3 := roll(ball_roll1, ball_roll2, 0, 0, 0)
		ball_roll4 := roll(ball_roll1, ball_roll2, ball_roll3, 0, 0)
		ball_roll5 := roll(ball_roll1, ball_roll2, ball_roll3, ball_roll4, 0)
		sortRolls := []int{ball_roll1, ball_roll2, ball_roll3, ball_roll4, ball_roll5}

		sort.Ints(sortRolls)
		//fmt.Println(sortRolls)

		ball_roll1 = sortRolls[0]
		ball_roll2 = sortRolls[1]
		ball_roll3 = sortRolls[2]
		ball_roll4 = sortRolls[3]
		ball_roll5 = sortRolls[4]

		numbers := &PowerBall{
			ball1: ball_roll1,
			ball2: ball_roll2,
			ball3: ball_roll3,
			ball4: ball_roll4,
			ball5: ball_roll5,
			pb:    (rand.Intn(26-1) + 1),
		}

		if i == 1 {
			jsonOut += fmt.Sprintf("{\"ballOne\": \"%d\", \"ballTwo\": \"%d\", \"ballThree\": \"%d\", \"ballFour\": \"%d\", \"ballFive\": \"%d\", \"pb\": \"%d\"}", numbers.ball1, numbers.ball2, numbers.ball3, numbers.ball4, numbers.ball5, numbers.pb)
		} else {
			jsonOut += fmt.Sprintf(",{\"ballOne\": \"%d\", \"ballTwo\": \"%d\", \"ballThree\": \"%d\", \"ballFour\": \"%d\", \"ballFive\": \"%d\", \"pb\": \"%d\"}", numbers.ball1, numbers.ball2, numbers.ball3, numbers.ball4, numbers.ball5, numbers.pb)
		}

		//Write latest value for each metric in the prometheus metric channel.
		//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
		//m1 := prometheus.MustNewConstMetric(collector.powerballMetric, prometheus.GaugeValue, float64(i), strconv.Itoa(numbers.ball1), strconv.Itoa(numbers.ball2), strconv.Itoa(numbers.ball3), strconv.Itoa(numbers.ball4), strconv.Itoa(numbers.ball5), strconv.Itoa(numbers.pb))
		//ch <- m1
		//fmt.Println(numbers)

	}
	jsonOut += fmt.Sprintf("]}")
	fmt.Println(jsonOut)

	res, err := PrettyString(jsonOut)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, res)
}

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", powerballRoll)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func duplicateRollCheck(rollIn int, roll1 int, roll2 int, roll3 int, roll4 int, roll5 int) bool {
	if rollIn == roll1 || rollIn == roll2 || rollIn == roll3 || rollIn == roll4 || rollIn == roll5 {
		return true
	}
	return false
}

func roll(roll1 int, roll2 int, roll3 int, roll4 int, roll5 int) int {

	time.Sleep(time.Millisecond * 100) // sleeping adds randomness to seed
	rand.Seed(time.Now().UnixNano())

	min := 1
	max := 69
	roll := rand.Intn(max-min) + min

	//Roll with loop
	if roll1 == 0 {
		//roll ball 1
		roll = rand.Intn(max-min) + min
	} else if roll2 == 0 {
		//roll ball 2
		for {
			roll = rand.Intn(max-min) + min
			if duplicateRollCheck(roll, roll1, roll2, roll3, roll4, roll5) == false {
				return roll
			} else {
				log.Warn("Duplicate found!, rerolling...")
			}
		}

	} else if roll3 == 0 {
		//roll ball 3
		for {
			roll = rand.Intn(max-min) + min
			if duplicateRollCheck(roll, roll1, roll2, roll3, roll4, roll5) == false {
				return roll
			} else {
				log.Warn("Duplicate found!, rerolling...")
			}
		}
	} else if roll4 == 0 {
		//roll ball 4
		for {
			roll = rand.Intn(max-min) + min
			if duplicateRollCheck(roll, roll1, roll2, roll3, roll4, roll5) == false {
				return roll
			} else {
				log.Warn("Duplicate found!, rerolling...")
			}
		}
	} else {
		//rolling ball 5
		for {
			roll = rand.Intn(max-min) + min
			if duplicateRollCheck(roll, roll1, roll2, roll3, roll4, roll5) == false {
				return roll
			} else {
				fmt.Println("Duplicate found! Rerolling...")
			}
		}
	}

	return roll
}

type PowerBall struct {
	ball1 int
	ball2 int
	ball3 int
	ball4 int
	ball5 int
	pb    int
}
