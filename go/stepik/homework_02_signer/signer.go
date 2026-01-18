package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var SingleHash job = func(in, out chan interface{}) {
	fmt.Printf("SingleHash!!!\n");
	counter := 0
	InputElemsCount := len(in)
	for data := range in {
		fmt.Printf("SingleHash get data!\n");
		dataString, ok := data.(string)
		if !ok {
			dataInt, ok := data.(int)
			if !ok {
				fmt.Printf("SingleHash not ok!\n");
				continue
			}
			dataString = strconv.Itoa(dataInt)
		}
		fmt.Printf("SingleHash: data = %v \n", dataString);
		out <- DataSignerCrc32(dataString) + "~" + DataSignerCrc32(DataSignerMd5(dataString))
		counter++
		if counter == InputElemsCount {
			break
		}
	}
}

var MultiHash job = func(in, out chan interface{}) {
	fmt.Printf("MultiHash!!!\n");
	preResult := make([]string, 6)
	counter := 0
	InputElemsCount := len(in)
	for data := range in {
		fmt.Printf("MultiHash get data!\n")
		dataString, ok := data.(string)
		if !ok {
			fmt.Printf("MultiHash data not ok!\n");
			continue
		}
		for th := 0; th <= 5; th++ {
			preResult[th] = DataSignerCrc32(strconv.Itoa(th) + dataString)
		}
		preResult := strings.Join(preResult, "")
		fmt.Printf("MultiHash: data = %v \n", preResult);
		out <- preResult
		counter++
		if counter == InputElemsCount {
			break
		}
	}
}

var CombineResults job = func(in, out chan interface{}) {
	fmt.Printf("CombineResults!!!\n");
	var allData []string = make([]string, 0, 200)
	counter := 0
	InputElemsCount := len(in)
	for data := range in {
		fmt.Printf("CombineResults get data!\n");
		data, ok := data.(string)
		if !ok {
			fmt.Printf("CombineResults data not ok!\n");
			continue
		}
		allData = append(allData, data)
		counter++
		if counter == InputElemsCount {
			break
		}
	}
	sort.Slice(allData, func(i, j int) bool {
		return allData[i] < allData[j]
	})
	str := strings.Join(allData, "_") 
	fmt.Printf("CombineResults = %v\n", str);
	out <- str
}

func ExecutePipeline(jobs ...job) {
	fmt.Printf("ExecutePipeline!!!\n");
	channels := make([](chan interface{}), len(jobs) + 1)
	for i := 0; i < len(jobs) + 1; i++ {
		channels[i] = make(chan interface{}, 200)
	}

	fmt.Printf("%v\n", jobs)
	for i, job := range(jobs) {
		job(channels[i], channels[i + 1])
	}

	for i := 0; i < len(jobs) + 1; i++ {
		close(channels[i])
	}
}

func main() {
	inputData := []int{0, 1, 1, 2, 3, 5, 8}
	hashSignJobs := []job{
		job(func(in, out chan interface{}) {
			for _, fibNum := range inputData {
				out <- fibNum
			}
		}),
		job(SingleHash),
		job(MultiHash),
		job(CombineResults),
		job(func(in, out chan interface{}) {
			dataRaw := <-in
			data, ok := dataRaw.(string)
			if !ok {
				fmt.Printf("Pipeline data not valid!\n");
			} else {
				fmt.Printf("CombineResults result data = %v\n", data);
			}
		}),
	}

	ExecutePipeline(hashSignJobs...)
}