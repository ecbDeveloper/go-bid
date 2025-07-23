package jsonutils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ecbDeveloper/go-bid/internal/validator"
)

// Essa função pega os dados que nós queremos retornar indepente do tipo
// faz um marshall e retorna pro usuario.
func EncodeJson[T any](w http.ResponseWriter, statusCode int, data T) error {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("failed to encode json %w", err)
	}

	return nil
}

// Essa função recebe o corpo da requisição faz o unmarshall dos dados para a struct
// que satisfaz a interface validator, valida os dados recebidos
// e nos retorna a struct pra lidarmos com ela, por exemplo inserir no banco.
func DecodeValidJson[T validator.Validator](r *http.Request) (T, map[string]string, error) {
	var data T
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return data, nil, fmt.Errorf("failed to decode json %w", err)
	}

	problems := data.Valid(r.Context())
	if len(problems) > 0 {
		return data, problems, fmt.Errorf("invalid %T: %d problems", data, len(problems))
	}

	return data, nil, nil
}

// Essa função recebe o corpo da requisição e faz o unmarshall dentro de uma struct
// que não precisa de uma validação, e nos devolve a struct para lidar com ela por
// exemplo inserir no banco.
func DecodeJson[T any](r *http.Request) (T, error) {
	var data T
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return data, fmt.Errorf("failed to decode json %w", err)
	}

	return data, nil
}
