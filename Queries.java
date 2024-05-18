import org.bson.Document;

// import static com.mongodb.client.model.Filters.regex;
import static com.mongodb.client.model.Sorts.ascending;

import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoClients;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoCursor;
import com.mongodb.client.MongoDatabase;

public class Queries {
    public static void main(String[] args) {
        MongoClient client = MongoClients.create("mongodb://localhost:27017");
        MongoDatabase database = client.getDatabase("mydb");
        database.createCollection("customers");
        MongoCollection<Document> collection = database.getCollection("customers");

        Document document = new Document();
        document.put("name", "Charlie");
        document.put("email_address", "charlie@gmail.com");
        collection.insertOne(document);

        document = new Document();
        document.put("name", "Bob");
        document.put("email_address", "bob@gmail.com");
        collection.insertOne(document);

        document = new Document();
        document.put("name", "Alice");
        document.put("email_address", "alice@outlook.com");
        collection.insertOne(document);

        Document searchQuery = new Document();
        searchQuery.put("email_address", "bob@gmail.com");
        // collection.find(searchQuery).first() retrieves the first result only
        // collection.distinct("name", String.class).cursor() retrieves all distinct names
        // collection.find(regex("email_address", "^bob@")).cursor() retrieves
        // all documents with email addresses starting with "bob@"
        MongoCursor<Document> cursor = collection.find(searchQuery).cursor();
        Document customer = cursor.next();
        System.out.println(customer);
        System.out.println(customer.toJson());
        // use customer.containsKey("field") to see if a field exists
        System.out.println(customer.getString("name"));
        System.out.println(customer.getString("email_address"));

        searchQuery = new Document();
        searchQuery.put("email_address", "alice@outlook.com");
        Document updatedField = new Document();
        updatedField.put("email_address", "alice@gmail.com");
        Document updater = new Document();
        updater.put("$set", updatedField);
        collection.updateMany(searchQuery, updater);

        searchQuery = new Document();
        searchQuery.put("email_address", "charlie@gmail.com");
        collection.deleteMany(searchQuery);

        // use descending("name") to sort in reverse order
        cursor = collection.find().sort(ascending("name")).cursor();
        while (cursor.hasNext()) {
            System.out.println(cursor.next());
        }

        System.out.println(collection.countDocuments());

        client.close();
    }
}
