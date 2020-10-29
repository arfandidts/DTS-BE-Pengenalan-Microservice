package handler

import (
	"net/http"
)

func AddMenu(w http.ResponseWriter, r *http.Request) {
	// response, err := json.Marshal(map[string]interface{}{
	// 	"success": true,
	// })

	// if err != nil {
	// 	fmt.Print("Failed to generate response")
	// 	return
	// }

	// w.Write(response)
	// Karena diatas terlalu panjang untuk setiap endpoint maka, dibuat
	// Wrapper mengatasi setiap response yang didapat
	utils.WrapAPISucces(w, r, "succes", http.StatusOK)
}
