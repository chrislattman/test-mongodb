const mongodb = require("mongodb");

const client = new mongodb.MongoClient("mongodb://localhost:27017");
const database = client.db("mydb");

async function run() {
    try {
        await database.createCollection("customers");
        const collection = database.collection("customers");

        let document = {name: "Charlie", email_address: "charlie@gmail.com"};
        await collection.insertOne(document);

        document = {name: "Bob", email_address: "bob@gmail.com"};
        await collection.insertOne(document);

        document = {name: "Alice", email_address: "alice@outlook.com"};
        await collection.insertOne(document);

        let searchQuery = {email_address: "bob@gmail.com"};
        // await collection.findOne(searchQuery) retrieves the first result only
        // await collection.distinct("name") retrieves all distinct names
        // collection.find({email_address: /^bob@/}) retrieves all documents
        // with email addresses starting with "bob@"
        let cursor = collection.find(searchQuery);
        const customer = await cursor.next();
        console.log(customer);
        console.log(JSON.stringify(customer));
        // use customer.hasOwnProperty("field") to see if a field exists
        console.log(customer.name);
        console.log(customer.email_address);

        searchQuery = {email_address: "alice@outlook.com"};
        const updatedField = {email_address: "alice@gmail.com"};
        const updater = {$set: updatedField};
        await collection.updateMany(searchQuery, updater);

        searchQuery = {email_address: "charlie@gmail.com"};
        await collection.deleteMany(searchQuery);

        // use "descending" to sort in reverse order
        cursor = collection.find().sort("name", "ascending");
        while (await cursor.hasNext()) {
            console.log(await cursor.next());
        }

        console.log(await collection.countDocuments());
    } finally {
        await client.close();
    }
}
run();
