// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package buoy contains the models and services for
// working with buoy data.
package buoy

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ArdanStudios/gotraining/06-testing/example1/mongodb"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// Condition contains information for an individual station.
	Condition struct {
		WindSpeed     float64 `bson:"wind_speed_milehour" json:"wind_speed_milehour"`
		WindDirection int     `bson:"wind_direction_degnorth" json:"wind_direction_degnorth"`
		WindGust      float64 `bson:"gust_wind_speed_milehour" json:"gust_wind_speed_milehour"`
	}

	// Location contains the buoys location.
	Location struct {
		Type        string    `bson:"type" json:"type"`
		Coordinates []float64 `bson:"coordinates" json:"coordinates"`
	}

	// Station contains information for an individual station.
	Station struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		StationID string        `bson:"station_id" json:"station_id"`
		Name      string        `bson:"name" json:"name"`
		LocDesc   string        `bson:"location_desc" json:"location_desc"`
		Condition Condition     `bson:"condition" json:"condition"`
		Location  Location      `bson:"location" json:"location"`
	}
)

// DisplayWindSpeed pretty prints wind speed.
func (buoyCondition *Condition) DisplayWindSpeed() string {
	return fmt.Sprintf("%.2f", buoyCondition.WindSpeed)
}

// DisplayWindGust pretty prints wind gust.
func (buoyCondition *Condition) DisplayWindGust() string {
	return fmt.Sprintf("%.2f", buoyCondition.WindGust)
}

// FindStation retrieves the specified station
func FindStation(stationID string) (*Station, error) {
	log.Printf("FindStation : Started : stationID[%s]\n", stationID)

	var buoyStation *Station
	if err := mongodb.Execute("buoy_stations",
		func(collection *mgo.Collection) error {
			queryMap := bson.M{"station_id": stationID}

			log.Printf("FindStation : MGO : db.buoy_stations.find(%s).limit(1)\n", mongodb.Log(queryMap))
			return collection.Find(queryMap).One(&buoyStation)
		}); err != nil {
		if err != mgo.ErrNotFound {
			log.Println("FindStation :", err)
			return nil, err
		}
	}

	log.Println("FindStation : Completed")
	return buoyStation, nil
}

// FindRegion retrieves the stations for the specified region
func FindRegion(region string, limit int) ([]Station, error) {
	log.Printf("FindRegion : Started : region[%s]\n", region)

	var buoyStations []Station
	if err := mongodb.Execute("buoy_stations",
		func(collection *mgo.Collection) error {
			queryMap := bson.M{"region": region}

			log.Printf("FindRegion : MGO : db.buoy_stations.find(%s).limit(%d)\n", mongodb.Log(queryMap), limit)
			return collection.Find(queryMap).Limit(limit).All(&buoyStations)
		}); err != nil {
		if err != mgo.ErrNotFound {
			log.Println("FindRegion :", err)
			return nil, err
		}
	}

	log.Println("FindRegion : Completed")
	return buoyStations, nil
}

// Print will display the station pretty.
func Print(station *Station) {
	data, err := json.MarshalIndent(station, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
