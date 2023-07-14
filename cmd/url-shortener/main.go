package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	//DONE: Implement config: cleanenv 

	cfg := config.MustLoad()
	fmt.Printf("cfg: %v\n", cfg)
	//TODO: Implement logger: slog 

	//TODO: Implement storage: sqlite 

	//TODO: Implement router: chi, "chi render"  

	//TODO: To run server: 

}