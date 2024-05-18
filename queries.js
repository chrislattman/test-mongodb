const mongodb = require("mongodb");

const client = new mongodb.MongoClient("mongodb://localhost:27017");
const database = client.db("mydb");

async function run() {
    try {
        await database.createCollection("customers");
        const collection = database.collection("customers");

        let document = {name: "Charlie", email_address: "charlie@gmail.com"};
        let res = await collection.insertOne(document);
        if (res.insertedId === null) {
            console.log("insert Charlie failed");
        }

        document = {name: "Bob", email_address: "bob@gmail.com"};
        res = await collection.insertOne(document);
        if (res.insertedId === null) {
            console.log("insert Bob failed");
        }

        document = {name: "Alice", email_address: "alice@outlook.com"};
        res = await collection.insertOne(document);
        if (res.insertedId === null) {
            console.log("insert Alice failed");
        }

        const data = [
            {name: "Daniel", email_address: "daniel@gmail.com"},
            {name: "Frank", email_address: "frank@gmail.com"},
        ];
        try {
            await collection.insertMany(data);
        } catch {
            console.log("insert Daniel and Frank failed");
        }

        // To create an index on the email_address field (for faster queries,
        // useful for unstructured data):
        // const index = await collection.createIndex("email_address");

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

        // To delete the index:
        // await collection.dropIndex(index);

        searchQuery = {email_address: "alice@outlook.com"};
        const updatedField = {email_address: "alice@gmail.com"};
        const updater = {$set: updatedField};
        const updated = await collection.updateMany(searchQuery, updater);
        if (updated.modifiedCount != 1) {
            console.log("update Alice failed");
        }

        searchQuery = {email_address: "charlie@gmail.com"};
        const deleted = await collection.deleteMany(searchQuery);
        if (deleted.deletedCount != 1) {
            console.log("delete Charlie failed");
        }

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
