package utils

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	database "github.com/siddhantprateek/ohmyback-end/rest_postgres/database"
	model "github.com/siddhantprateek/ohmyback-end/rest_postgres/models"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func ReadXMLs() {
	var DB *gorm.DB
	DB = database.DatabaseConnection()
	// Need changes in model
	DB.AutoMigrate(model.TrademarkApplication{})
	for i := 1; i < 2; i++ {
		ext := ".xml"
		XMLFileName := strconv.Itoa(i) + ext
		// Reads XML file
		XMLFile, err := os.Open("data/" + XMLFileName)
		if err != nil {
			fmt.Println(err)
		}
		// Close XML file
		defer XMLFile.Close()

		// @returns byte[] value of xml file
		byteValue, _ := ioutil.ReadAll(XMLFile)

		var XMLDataType model.TrademarkApplication
		xml.Unmarshal([]byte(byteValue), &XMLDataType)

		// Retrieve data from XML files
		// id := XMLDataType.TrademarkBag.Trademark.ApplicationNumber.ST13ApplicationNumber
		trademarkName := XMLDataType.TrademarkBag.Trademark.ApplicantBag.Applicant.LegalEntityName
		// country := XMLDataType.TrademarkBag.Trademark.ApplicantBag.Applicant.NationalLegalEntityCode

		// Replacing  literals
		// if third argument is -1 then there is no limit to replacement
		trademarkName = strings.Replace(trademarkName, "-", "", -1)
		trademarkName = strings.Replace(trademarkName, ",", "", -1)
		trademarkName = strings.Replace(trademarkName, ".", "", -1)

		var jsonConvertedData datatypes.JSON

		jsonConvertedData, err = json.Marshal(XMLDataType)
		if err != nil {
			fmt.Println("Error in Marshalling XML data to JSON data")
		}
		// will be stored in database
		// in different format
		fmt.Println(jsonConvertedData)
	}

}
