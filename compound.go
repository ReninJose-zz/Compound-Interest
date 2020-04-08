
/*
The following Program Calculates Compound interest of multiple items and
stores it in a "Calculated.txt" file (Will be generated automatically).
NOTE: For the program to run, the user will have to input the duration of
inflation, and make a new file manually, which will be named "items.txt"
where the intial content will be stored.
Allowed Format: 'item_name' : 'float/decimal'
*/

package main

import ("fmt"
        "os"
        "bufio"
        "strings"
        "strconv"
        "math"
       )

// Counts the number of lines of the file and parses it
func Read_file(fl_name *os.File) ([]string, []float64, int, error) {
  var (item_name []string
       price []float64
       name string
       s []string
       )

  counter := 0
  reading_lines := bufio.NewScanner(fl_name)

  for reading_lines.Scan() {
      line := reading_lines.Text()
      s = strings.Split(line,":")
      name = strings.TrimSpace(s[0])
      prc, err := strconv.ParseFloat(strings.TrimSpace(s[1]), 64)
      if err != nil {
        return nil, nil, 1, fmt.Errorf("Couldn't convert float64 to string")
      }
      item_name = append(item_name, name)
      price = append(price, prc)
      counter++
  }

  return item_name, price, counter, nil
}

//Calculate the price based on Compound Interest formula
func calculate(result_cst []float64, cst []float64, count int, yrs int) ([]float64) {

  for i:=0 ; i<count ; i++ {
    result_cst[i] = cst[i]*(math.Pow(1+(0.02/1), 1*float64(yrs)))
  }
  return result_cst
}

//Using the new file, Write in new values
func update_file(up_fil *os.File, itm []string, calc_price []float64, cnt int) (fl *os.File, err error) {

  for i:=0 ; i<cnt ; i++ {
    bytes, err := fmt.Fprintf(up_fil, "%s : %f\n", itm[i], calc_price[i])
    if err != nil {
        return nil, err
    }
    fmt.Printf("Written bytes: %d\n", bytes)
  }
  return up_fil, nil
}

func main() {
  var years int
  var result_cost []float64
  var calculated_price []float64

  file_name, err := os.Open("items.txt")
  if err != nil {
    fmt.Println("Error: Can not open file")
    os.Exit(1)
  }

  //Function to read the file items.txt line by line
  item, cost, count_lines, err := Read_file(file_name)
  if err != nil {
    fmt.Println("Error: %v", err)
    os.Exit(1)
  }

  fmt.Printf("Enter the number of years to calculate the Compound Interest: \n")
  fmt.Scanf("%d", &years)

  //Dynamically allocating space for the inflated cost
  result_cost = make([]float64, count_lines)

  //Returning calculated array
  for i:=0 ; i<years ; i++ {
    calculated_price = calculate(result_cost, cost, count_lines, years)
  }

  //Creating a new file which will contain the calculated data
  new_file, err := os.Create("Calculated.txt")
  if err != nil {
    fmt.Println("Error: %v", err)
    os.Exit(1)
  }

  //Returning updated value
  updated_file, err := update_file(new_file, item, calculated_price, count_lines)
  if err != nil {
    fmt.Printf("Error: Could not Write the file succesfully: %v", err)
    os.Exit(1)
  }

  updated_file.Close()
  file_name.Close()
}
