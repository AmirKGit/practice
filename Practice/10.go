// Tez orada RoboContest tizimining qanday ishlayotganligini tekshirish uchun RoboticsLab direktori ishxonamizga tashrif buyuradi. Biz direktorni tantanali kutib olish maqsadida direktorning moshinadan tushgan joyidan ishxonamiz eshigigacha bo'lgan oraliqga gilam to'shashga qaror qildik. Bizga ma'lumki direktorning mashinasi to'xtatiladigan joydan ishxonamiz kirish eshigigacha bo'lgan masofa 
// �
// N metr, bozorda 1 metr gilamning narxi 
// �
// P so'm. Biz gilam sotib olish uchun jami qancha mablag' sarflashimizni aniqlang!.

package main

import (
    "fmt"
)

func main(){
    var N,P int
	fmt.Println("Enter metr")
	fmt.Scan(&N)
	fmt.Println("Som kirit")
	fmt.Scan(&P)

	fmt.Println(N*P)
}