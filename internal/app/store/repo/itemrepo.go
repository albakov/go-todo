package repo

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/albakov/go-todo/internal/app/model"
)

type ItemRepo struct {
}

var path = "/home/izmail/www/go-todo.ru/internal/app/model/items.json"

func (i *ItemRepo) Get() []model.Item {
	return parseItems()
}

func (i *ItemRepo) GetById(id int) (model.Item, error) {
	items := parseItems()
	item := model.Item{}

	for i := 0; i < len(items); i++ {
		if items[i].Id == id {
			return items[i], nil
		}
	}

	return item, errors.New("not found")
}

func (i *ItemRepo) Create(item model.Item) (model.Item, error) {
	if item.Title == "" {
		return item, errors.New("item Title cannot be empty")
	}

	data := parseItems()
	count := len(data)
	item.Id = count + 1
	data = append(data, item)

	err := saveJson(data)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (i *ItemRepo) Update(id int, item model.Item) error {
	if item.Title == "" {
		return errors.New("item Title cannot be empty")
	}

	_, err := i.GetById(id)
	if err != nil {
		return err
	}

	data := parseItems()

	for i := 0; i < len(data); i++ {
		if data[i].Id == id {
			data[i].Title = item.Title
		}
	}

	err = saveJson(data)
	if err != nil {
		return err
	}

	return nil
}

func (i *ItemRepo) Delete(id int) error {
	data := parseItems()

	for i := 0; i < len(data); i++ {
		if data[i].Id == id {
			data[i] = data[len(data)-1]
			data[len(data)-1] = model.Item{}
			data = data[:len(data)-1]
		}
	}

	err := saveJson(data)
	if err != nil {
		return err
	}

	return nil
}

func parseItems() []model.Item {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}

	type items struct {
		Items []model.Item `json:"items"`
	}

	data := items{}

	_ = json.Unmarshal([]byte(file), &data)

	return data.Items
}

func saveJson(data []model.Item) error {
	type items struct {
		Items []model.Item `json:"items"`
	}

	file, err := json.MarshalIndent(items{Items: data}, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, file, 0777)
	if err != nil {
		return err
	}

	return nil
}
