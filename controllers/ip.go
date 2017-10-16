package controllers

import (
	"github.com/lflxp/ips/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about object
type IpController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (o *IpController) Post() {
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is emptyx
// @router /:objectId [get]
func (o *IpController) Get() {
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId != "" {
		json := models.ParseIp(objectId)

		o.Data["json"] = map[string]interface{}{
			"time": json.Time,
			"ip":json.Ip,
			"GeoIP":map[string]interface{}{
				"Locations":map[string]interface{}{
					"Geoname_id":json.Locations.GeonameId,
					"LocaleCode":json.Locations.LocaleCode,
					"ContinentCode":json.Locations.ContinentCode,
					"ContinentName":json.Locations.ContinentName,
					"CountryIsoCode":json.Locations.CountryIsoCode,
					"CountryName":json.Locations.CountryName,
					"S1IsoCode":json.Locations.S1IsoCode,
					"S1Name":json.Locations.S1Name,
					"S2IsoCode":json.Locations.S2IsoCode,
					"S2Name":json.Locations.S2Name,
					"CityName":json.Locations.CityName,
					"MetroCode":json.Locations.MetroCode,
					"TimeZone":json.Locations.TimeZone,
				},
				"Blocks":map[string]interface{}{
					"Start":json.Blocks.Start,
					"End":json.Blocks.End,
					"FirstIp":json.Blocks.FirstIp,
					"EndIp":json.Blocks.EndIp,
					"Network":json.Blocks.Network,
					"Geoname_id":json.Blocks.Geoname_id,
					"Registered_country_geoname_id":json.Blocks.Registered_country_geoname_id,
					"Represented_country_geoname_id":json.Blocks.Represented_country_geoname_id,
					"Is_anonymous_proxy":json.Blocks.Is_anonymous_proxy,
					"Is_satellite_provider":json.Blocks.Is_satellite_provider,
					"Postal_code":json.Blocks.Postal_code,
					"Latitude":json.Blocks.Latitude,
					"Longitude":json.Blocks.Longitude,
					"Accuracy_radius":json.Blocks.Accuracy_radius,
				},
				"Asn":map[string]interface{}{
					"Network":json.Asn.Network,
					"Autonomous_system_number":json.Asn.Autonomous_system_number,
					"Autonomous_system_organization":json.Asn.Autonomous_system_organization,
				},
			},
			"status":json.Status,
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *IpController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *IpController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (o *IpController) Delete() {
	objectId := o.Ctx.Input.Param(":objectId")
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}

