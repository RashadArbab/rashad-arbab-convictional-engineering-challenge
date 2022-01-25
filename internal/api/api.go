package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rashad-arbab-convictional-engineering-interview/internal/datasource"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	data, err := datasource.GetData()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	var reqProducts []RequestProduct
	var products []Product

	json.Unmarshal(data, &reqProducts)
	// I questioned this decision for some time, because its never pretty when you have
	// loops within loops but i concluded that no matter what you do you will have to touch
	// each individual image and so the loops will run for the number of images from the datasource
	fmt.Println(fmt.Sprint(reqProducts[0].Variants[0].Id))
	for i, product := range reqProducts {
		pr := Product{
			Code:     strconv.Itoa(reqProducts[i].Id),
			Title:    reqProducts[i].Title,
			Vendor:   reqProducts[i].Vendor,
			BodyHtml: reqProducts[i].BodyHtml,
			Variants: ReqVariantToVariant(reqProducts[i].Variants),
			Images:   []Image{},
		}

		products = append(products, pr)
		for j, variant := range product.Variants {
			//this is necessary because the json data coming in, is an int for the id
			//and the output has to be in string.
			products[i].Variants[j].Id = strconv.Itoa(reqProducts[i].Variants[j].Id)

			//ASSUMPTION *** Im making the assumption that the position of a variant, describes the
			//availablity of inventory, and the number in stock. I looked through barcode, I looked
			//through the timestamps for clues but this is the most probable format for stock that
			//I found. Without more info or a breakdown it is hard to be sure though.
			//I initially I thought that the number of variants would dictate the number in stock however since we are looking for
			//the number in stock per variant this no longer made sense.
			if variant.Position != 0 {
				products[i].Variants[j].Available = true
				products[i].Variants[j].InventoryQuantity = int64(variant.Position)
			}

			for _, image := range variant.Images {
				im := Image{
					Source:    image.Source,
					VariantID: products[i].Variants[j].Id,
				}
				products[i].Images = append(products[i].Images, im)
			}
		}
	}

	fmt.Println(len(products))

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: "no id found"})
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: "invalid id supplied"})
	}

	data, err := datasource.GetData()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	var reqProducts []RequestProduct
	json.Unmarshal(data, &reqProducts)

	var specificProduct RequestProduct

	//Step 1 find the specific product with the id
	found := false
	for _, product := range reqProducts {
		if product.Id == idInt {
			byteData, _ := json.Marshal(product)
			json.Unmarshal(byteData, &specificProduct)
			found = true
		}
	}

	if !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Message: "Product not found"})
		return
	}

	product := Product{
		Code:     strconv.Itoa(specificProduct.Id),
		Title:    specificProduct.Title,
		Vendor:   specificProduct.Vendor,
		BodyHtml: specificProduct.BodyHtml,
		Variants: ReqVariantToVariant(specificProduct.Variants),
		Images:   []Image{},
	}

	//Step 2 move the specific items in the incoming json to the shape of the output json
	for j, variant := range specificProduct.Variants {

		product.Variants[j].Id = strconv.Itoa(specificProduct.Variants[j].Id)
		if variant.Position != 0 {
			product.Variants[j].Available = true
			product.Variants[j].InventoryQuantity = int64(variant.Position)
		}

		for _, image := range variant.Images {
			im := Image{
				Source:    image.Source,
				VariantID: strconv.Itoa(specificProduct.Variants[j].Id),
			}
			product.Images = append(product.Images, im)
		}

	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func GetInventory(w http.ResponseWriter, r *http.Request) {

	data, err := datasource.GetData()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	var reqProducts []RequestProduct
	var inventory []Inventory

	json.Unmarshal(data, &reqProducts)
	// I questioned this decision for some time, because its never pretty when you have
	// loops within loops but i concluded that no matter what you do you will have to touch
	// each individual image and so the loops will run for the number of images from the datasource
	fmt.Println(fmt.Sprint(reqProducts[0].Variants[0].Id))
	for _, product := range reqProducts {
		for _, variant := range product.Variants {
			inventory = append(inventory, Inventory{
				ProductId: strconv.Itoa(product.Id),
				VariantId: strconv.Itoa(variant.Id),
				Stock:     int64(variant.Position),
			})
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inventory)
}

func ReqVariantToVariant(req []RequestVariant) []Variant {
	var targetStruct []Variant
	temporaryVariable, _ := json.Marshal(req)
	err := json.Unmarshal(temporaryVariable, &targetStruct)
	if err != nil {
		log.Println()
	}

	for i, variant := range req {
		targetStruct[i].Weight.Value = variant.Weight
		targetStruct[i].Weight.Unit = variant.WeightUnit
	}

	return targetStruct
}
