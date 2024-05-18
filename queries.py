# import re

from bson.json_util import dumps
from pymongo import MongoClient, ASCENDING

client = MongoClient("localhost", 27017)
database = client.get_database("mydb")
database.create_collection("customers")
collection = database.get_collection("customers")

document = {"name": "Charlie", "email_address": "charlie@gmail.com"}
collection.insert_one(document)

document = {"name": "Bob", "email_address": "bob@gmail.com"}
collection.insert_one(document)

document = {"name": "Alice", "email_address": "alice@outlook.com"}
collection.insert_one(document)

search_query = {"email_address": "bob@gmail.com"}
# collection.find_one(search_query) retrieves the first result only
# collection.distinct("name") retrieves all distinct names
# collection.find({"email_address": re.compile("^bob@")}) retrieves all documents
# with email addresses starting with "bob@"
cursor = collection.find(search_query)
customer = cursor.next()
print(customer)
print(dumps(customer))
# use "field" in customer to see if a field exists
print(customer["name"])
print(customer["email_address"])

search_query = {"email_address": "alice@outlook.com"}
updated_field = {"email_address": "alice@gmail.com"}
updater = {"$set": updated_field}
collection.update_many(search_query, updater)

search_query = {"email_address": "charlie@gmail.com"}
collection.delete_many(search_query)

# use DESCENDING to sort in reverse order
cursor = collection.find().sort("name", ASCENDING)
for document in cursor:
    print(document)

print(collection.count_documents({}))

client.close()
