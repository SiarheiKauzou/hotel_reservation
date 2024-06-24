package main

import (
	"context"
	"fmt"
	"log"

	"github.com/SiarheiKauzou/hotel_reservation/db"
	"github.com/SiarheiKauzou/hotel_reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, db.DBNAME)

	hotel := types.Hotel{
		Name:     "Bellucia",
		Location: "Narnia",
	}
	rooms := []types.Room{
		{
			Type:      types.SingleRoomType,
			BasePrice: 88.9,
		},
		{
			Type:      types.DeluxeRoomType,
			BasePrice: 198.2,
		},
		{
			Type:      types.SeaSideRoomType,
			BasePrice: 122.4,
		},
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		insertedRoom, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(insertedRoom)
	}
}
