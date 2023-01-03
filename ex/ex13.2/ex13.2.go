package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPuser struct {
	UserInfo User
	VIPLever int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VIPuser{
		User{"화랑", "hwarang", 40},
		3,
		250,
	}

	fmt.Printf("유저 : %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d VIP 가격: %d만 원 \n",
		vip.UserInfo.Name,
		vip.UserInfo.ID,
		vip.UserInfo.Age,
		vip.VIPLever,
		vip.Price,
	)
}
