package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(`NEXT_PUBLIC_BW_SECRET_KEY="` + os.Getenv("NEXT_PUBLIC_BW_SECRET_KEY") + `"`)
	fmt.Println(`NEXT_PUBLIC_SUPABASE_URL="` + os.Getenv("NEXT_PUBLIC_SUPABASE_URL") + `"`)
	fmt.Println(`NEXT_PUBLIC_SUPABASE_ANON_KEY="` + os.Getenv("NEXT_PUBLIC_SUPABASE_ANON_KEY") + `"`)
}
