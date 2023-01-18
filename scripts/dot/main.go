package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(`MONGO_URL="` + os.Getenv("MONGO_URL") + `"`)
	fmt.Println(`NEXT_PUBLIC_FULL=` + os.Getenv("NEXT_PUBLIC_FULL"))
	fmt.Println(`EMAIL_FROM="` + os.Getenv("EMAIL_FROM") + `"`)
	fmt.Println(`SENDGRID_API_KEY="` + os.Getenv("SENDGRID_API_KEY") + `"`)
	fmt.Println(`NEXT_PUBLIC_BW_SECRET_KEY="` + os.Getenv("NEXT_PUBLIC_BW_SECRET_KEY") + `"`)
}
