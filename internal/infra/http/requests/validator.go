package requests

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var v = validator.New()

type requestType interface {
	ToDomainModel() (interface{}, error)
}

func Bind[reqType requestType, domain interface{}](r *http.Request, req reqType, targetType domain) (domain, error) {
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Print(err)
		return targetType, err
	}

	if err := v.Struct(req); err != nil {
		log.Print(err)
		return targetType, err
	}

	d, err := req.ToDomainModel()
	if err != nil {
		log.Print(err)
		return targetType, err
	}

	return d.(domain), nil
}

func BindToMap[reqType requestType](r *http.Request, req reqType) (map[string]interface{}, error) {
	// Розпарсити тіло JSON у структуру запиту
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Print(err)
		return nil, err
	}

	// Провалідувати структуру
	if err := v.Struct(req); err != nil {
		log.Print(err)
		return nil, err
	}

	// Отримати доменну модель
	domainModel, err := req.ToDomainModel()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	// Перетворити у мапу з JSON-тегами
	resultMap, err := structToMapLowercaseKeys(domainModel)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return resultMap, nil
}

func structToMapLowercaseKeys(data interface{}) (map[string]interface{}, error) {
	// Marshal структуру у JSON
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Unmarshal у map
	var rawMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &rawMap)
	if err != nil {
		return nil, err
	}

	// Створюємо нову мапу з маленькими ключами та не нульовими значеннями
	loweredMap := make(map[string]interface{})

	// Перетворюємо всі ключі на малі літери та перевіряємо на нульові значення
	for key, value := range rawMap {
		// Пропускаємо певні поля: createddate, updateddate, deleteddate
		if isIgnoredField(key) {
			continue
		}

		// Перевіряємо чи значення є нульовим або порожнім
		if !isZero(value) {
			loweredMap[strings.ToLower(key)] = value
		}
	}

	return loweredMap, nil
}

// Функція для перевірки чи значення є нульовим або пустим
func isZero(value interface{}) bool {
	// Перевіряємо чи значення є nil
	if value == nil {
		return true
	}

	// Перевіряємо чи це порожній рядок
	if str, ok := value.(string); ok && str == "" {
		return true
	}

	// Використовуємо рефлексію для перевірки нульових значень для інших типів
	val := reflect.ValueOf(value)
	return val.IsZero()
}

// Функція для перевірки чи поле є одним з ігнорованих
func isIgnoredField(fieldName string) bool {
	ignoredFields := []string{"createddate", "updateddate", "deleteddate"}
	for _, ignored := range ignoredFields {
		if strings.ToLower(fieldName) == ignored {
			return true
		}
	}
	return false
}
