package main

type User struct {
	Membership
	Name string
}

type Membership struct{
	Type string
	MessageCharLimit int
}
func newUser(name string, membershipType string) User {
	if membershipType=="premium"{
		return  User{
			Name: 	name,
			Membership : Membership{Type : 	membershipType,
									MessageCharLimit : 1000	,},
		}	
	}else {
		return User{
			Name: name,
			Membership : Membership{Type : membershipType,
			MessageCharLimit : 100,},
	}
}
}
