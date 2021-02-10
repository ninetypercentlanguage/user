wraps postgres that holds things that words are associated with:
- which words a user knows
- which words are in given text
user knows table would be like:
id | word | partofspeech | user

psql "dbname=postgres host=localhost user=postgres password=pw port=5432 sslmode=disable"
