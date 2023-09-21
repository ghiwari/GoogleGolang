package _5_misc

type distance *func()

// ERROR : distance is not a type
// ERROR : if you change distance to type then we wil get
//"Invalid receiver type 'distance' ('distance' is a pointer type)"

//func (d *distance) me1() { //This will wgive error
//	fmt.Println("Hello World")
//}
