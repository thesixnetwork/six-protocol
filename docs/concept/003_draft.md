## Concept Design For NFT Schema Combination aka `Virtual Schema`
    
### Why:
    1. We want to make two or more schema to work together. 
    
    For example:
        Alice have two schemas (Schema A and Schema B) User of two schemas are the same. 
        Mean that owner of Token A and Token B are the same person. (TokenId:1, Schema:A, Owner:Bob) (TokenId:1, Schema:B owner:Bob)
        Alice want to combine these schema together. When user of Schema A do some action, token point of Schema A increase and also point of Schema B will increase as well
        And user are able to bridge some attribute value accross schema


### How:
    1. Someone propose to create virtual schema with lock some asset (to prevent spamming issue)
    2. Creator of Virtual Schema add action info of each schema
    3. Owner of Actual Schema will accept or reject proposal of Virtual schema with in 3 days. 
    If owner of schema did not participate locked asset will return only 70%. Which mean owner of schema has to reject or accept
    4. When expiration occur virtual schema are invalid and locked asset will be return
    5. Owner of schema will allow some of attribute to participate, which mean virtual schema cannot fully access to all attribute
    6. Virtual Action can be perform by admin or action executor of each schema.
