package controllers

import (
	"log"
	transportation "mongoApp/app/models/transportation"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
	model *transportation.Transportation
}

func (c App) Get() revel.Result {
	var (
		transportations []transportation.SelectTransportation
		err             error
	)
	// Инициализация модели
	c.model = new(transportation.Transportation)
	if err = c.model.Init(); err != nil {
		log.Fatal(err)
	}

	if transportations, err = c.model.Get(); err != nil {
		log.Fatal(err)
	}

	//формируем мапу из массива структур
	/*dt, _ := json.Marshal(dataModel)

	var dataToFront map[string]interface{}
	dataToFront = make(map[string]interface{})

	dataToFront["data"] = string(dt)*/

	return c.RenderJson(transportations)
}

func (c App) Post() revel.Result {
	var (
		params map[string]string // Мапа для параметров запроса
		err    error
	)
	// Инициализация модели
	c.model = new(transportation.Transportation)
	if err = c.model.Init(); err != nil {
		log.Fatal(err)
	}
	params = make(map[string]string)

	params["routeLength"] = c.Params.Get("routeLength")
	params["fromAddress"] = c.Params.Get("fromAddress")
	params["toAddress"] = c.Params.Get("toAddress")
	params["carModel"] = c.Params.Get("carModel")
	params["carNumber"] = c.Params.Get("carNumber")
	params["driverName"] = c.Params.Get("driverName")
	params["driverPhone"] = c.Params.Get("driverPhone")

	var dataToFront map[string]interface{}
	dataToFront = make(map[string]interface{})

	// Вызываем метод добавления новой грузоперевозки
	if err := c.model.Post(params); err != nil {
		dataToFront["errorStatus"] = 1
		dataToFront["errorText"] = err
	} else {
		dataToFront["errorStatus"] = 0
	}
	return c.RenderJson(dataToFront)
}

func (c App) Delete() revel.Result {
	var (
		id          string
		dataToFront map[string]interface{}
		err         error
	)
	// Инициализация модели
	c.model = new(transportation.Transportation)
	if err = c.model.Init(); err != nil {
		log.Fatal(err)
	}

	id = c.Params.Get("id")
	dataToFront = make(map[string]interface{})

	// Вызываем метод удаления грузоперевозки
	if err = c.model.Delete(id); err != nil {
		dataToFront["errorStatus"] = 1
		dataToFront["errorText"] = err.Error()
	} else {
		dataToFront["errorStatus"] = 0
	}
	return c.RenderJson(dataToFront)
}