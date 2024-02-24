package controller

const connectionString = "mongodb+srv://vijay:<#Vijay2532000>@cluster0.a3sydpp.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"

const colName = "watchlist"

// Taking reference of mongodb collection(Collection in MongoDB is like tables in MYSQL)
// Mngo DB in NoSQL Database
//var collection *mongo.Collection

//connect with mongoDB

func init() {
	//Client options
	// clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	//We need to provide a contect which mean how long do we need to connected with the DB in Webserver
	//We have given the context type as Todo
	// client, err := mongo.Connect(context.TODO(), clientOption)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Connection made")

	// collection = client.Database(dbName).Collection(colName)

	// //Collection instance
	// fmt.Println("Collection instance/Refernec is ready")
}
