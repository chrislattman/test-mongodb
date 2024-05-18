# import re

from bson.json_util import dumps
from pymongo import ASCENDING, MongoClient
from pymongo.errors import PyMongoError

client = MongoClient("localhost", 27017)
database = client.get_database("mydb")
database.create_collection("customers")
collection = database.get_collection("customers")

document = {"name": "Charlie", "email_address": "charlie@gmail.com"}
res = collection.insert_one(document)
if res.inserted_id is None:
    print("insert Charlie failed")

document = {"name": "Bob", "email_address": "bob@gmail.com"}
res = collection.insert_one(document)
if res.inserted_id is None:
    print("insert Bob failed")

document = {"name": "Alice", "email_address": "alice@outlook.com"}
res = collection.insert_one(document)
if res.inserted_id is None:
    print("insert Alice failed")

data = [
    {"name": "Daniel", "email_address": "daniel@gmail.com"},
    {"name": "Frank", "email_address": "frank@gmail.com"},
]
try:
    collection.insert_many(data)
except PyMongoError:
    print("insert Daniel and Frank failed")

# index = collection.create_index("email_address")

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

# collection.drop_index(index)

search_query = {"email_address": "alice@outlook.com"}
updated_field = {"email_address": "alice@gmail.com"}
updater = {"$set": updated_field}
updated = collection.update_many(search_query, updater)
if updated.modified_count != 1:
    print("update Alice failed")

search_query = {"email_address": "charlie@gmail.com"}
deleted = collection.delete_many(search_query)
if deleted.deleted_count != 1:
    print("delete Charlie failed")

# use DESCENDING to sort in reverse order
cursor = collection.find().sort("name", ASCENDING)
for document in cursor:
    print(document)

print(collection.count_documents({}))

client.close()
