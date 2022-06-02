package main



import (

   "testing"

)



func Testsample(t *testing.T){

     var sample =  20
     var datas  = 30


     if sample != datas {
         t.Errorf("value is wrong")

    }

}

